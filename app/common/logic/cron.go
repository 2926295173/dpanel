package logic

import (
	"bytes"
	"context"
	"fmt"
	"github.com/donknap/dpanel/common/accessor"
	"github.com/donknap/dpanel/common/dao"
	"github.com/donknap/dpanel/common/entity"
	"github.com/donknap/dpanel/common/function"
	"github.com/donknap/dpanel/common/service/crontab"
	"github.com/donknap/dpanel/common/service/docker"
	"github.com/donknap/dpanel/common/service/exec"
	"github.com/robfig/cron/v3"
	"io"
	"log/slog"
	"os"
	"sync"
	"syscall"
	"time"
)

const (
	CronRunDockerExec = "dockerExec"
)

var (
	lock = sync.RWMutex{}
)

type Cron struct {
}

func (self Cron) AddJob(task *entity.Cron) ([]cron.EntryID, error) {
	option := make([]crontab.Option, 0)
	option = append(option, crontab.WithName(task.Title))
	expression := make([]string, 0)
	for _, item := range task.Setting.Expression {
		expression = append(expression, item.ToString())
	}

	option = append(option, crontab.WithRunFunc(func() error {
		if task.Setting.EnableRunBlock {
			if lock.TryLock() {
				defer lock.Unlock()
			} else {
				slog.Debug("task skipped")
				return nil
			}
		}
		var ctx context.Context
		if task.Setting.ScriptRunTimeout > 0 {
			ctx, _ = context.WithTimeout(context.Background(), time.Second*time.Duration(task.Setting.ScriptRunTimeout))
		}
		startTime := time.Now()
		containerName := task.Setting.ContainerName

		defaultDockerEnv, err := Setting{}.GetDockerClient(task.Setting.DockerEnvName)
		if err != nil {
			return err
		}
		globalEnv := make([]string, 0)
		globalEnv = append(globalEnv, defaultDockerEnv.GetDockerEnv()...)
		globalEnv = append(globalEnv, function.PluckArrayWalk(task.Setting.Environment, func(i docker.EnvItem) (string, bool) {
			return fmt.Sprintf("%s=%s", i.Name, i.Value), true
		})...)

		var out string
		var script string

		if containerName != "" {
			dockerClient, err := docker.NewBuilderWithDockerEnv(defaultDockerEnv)
			if err != nil {
				return err
			}
			if task.Setting.EntryShell == "" {
				task.Setting.EntryShell = "/bin/sh"
			}
			defer func() {
				dockerClient.Close()
			}()
			script = fmt.Sprintf(`docker exec %s %s -c "%s"`, containerName, task.Setting.EntryShell, task.Setting.Script)
		} else {
			script = task.Setting.Script
		}

		// 如果没有指定容器，则直接在面板 shell 中执行
		// 在面板容器中执行还需要把 env 注入到命令中
		globalEnv = append(globalEnv, os.Environ()...)
		options := []exec.Option{
			exec.WithCommandName("/bin/sh"),
			exec.WithArgs("-c", script),
			exec.WithEnv(globalEnv),
		}

		cmd, _ := exec.New(options...)

		if ctx != nil {
			cmd.Cmd().SysProcAttr = &syscall.SysProcAttr{
				Setpgid: true,
				Pgid:    0,
			}
			go func() {
				select {
				case <-ctx.Done():
					slog.Debug("cron run timeout", "timeout", task.Setting.ScriptRunTimeout, "task", task)
					cmd.Close()
				}
			}()
		}

		response, err := cmd.Run()
		if err != nil {
			_ = dao.CronLog.Create(&entity.CronLog{
				CronID: task.ID,
				Value: &accessor.CronLogValueOption{
					Error:   err.Error(),
					RunTime: startTime,
					UseTime: time.Now().Sub(startTime).Seconds(),
				},
			})
			return err
		}
		buffer := new(bytes.Buffer)
		_, err = io.Copy(buffer, response)
		if err != nil {
			_ = dao.CronLog.Create(&entity.CronLog{
				CronID: task.ID,
				Value: &accessor.CronLogValueOption{
					Error:   err.Error(),
					RunTime: startTime,
					UseTime: time.Now().Sub(startTime).Seconds(),
				},
			})
			return err
		}
		out = buffer.String()

		_ = dao.CronLog.Create(&entity.CronLog{
			CronID: task.ID,
			Value: &accessor.CronLogValueOption{
				Message: out,
				RunTime: startTime,
				UseTime: time.Now().Sub(startTime).Seconds(),
			},
		})
		keepLogTotal := task.Setting.KeepLogTotal
		if keepLogTotal <= 0 {
			keepLogTotal = 10
		}
		logIds := make([]int32, 0)
		_ = dao.CronLog.Where(dao.CronLog.CronID.Eq(task.ID)).Order(dao.CronLog.ID.Desc()).Limit(keepLogTotal).Pluck(dao.CronLog.ID, &logIds)
		if len(logIds) >= keepLogTotal {
			_, _ = dao.CronLog.Where(dao.CronLog.CronID.Eq(task.ID)).Where(dao.CronLog.ID.NotIn(logIds...)).Delete()
		}
		return nil
	}))

	option = append(option, crontab.WithRunFunc(func() error {
		task.Setting.NextRunTime = crontab.Wrapper.GetNextRunTime(task.Setting.JobIds...)
		err := dao.Cron.Save(task)
		return err
	}))

	cronJob := crontab.New(option...)
	return crontab.Wrapper.AddJob(cronJob, expression...)
}
