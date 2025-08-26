package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/docker/docker/api"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/donknap/dpanel/app/common/logic"
	"github.com/donknap/dpanel/common/accessor"
	"github.com/donknap/dpanel/common/dao"
	"github.com/donknap/dpanel/common/entity"
	"github.com/donknap/dpanel/common/function"
	"github.com/donknap/dpanel/common/service/docker"
	"github.com/donknap/dpanel/common/service/notice"
	"github.com/donknap/dpanel/common/service/plugin"
	"github.com/donknap/dpanel/common/service/registry"
	"github.com/donknap/dpanel/common/service/ssh"
	"github.com/donknap/dpanel/common/service/storage"
	"github.com/donknap/dpanel/common/service/ws"
	"github.com/donknap/dpanel/common/types/define"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/mcuadros/go-version"
	"github.com/we7coreteam/w7-rangine-go/v2/pkg/support/facade"
	"github.com/we7coreteam/w7-rangine-go/v2/src/http/controller"
	ssh2 "golang.org/x/crypto/ssh"
	"io"
	"log/slog"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type command struct {
	Type string `json:"type"`
	Size struct {
		Width  int `json:"width"`
		Height int `json:"height"`
	} `json:"size"`
	Content struct {
		Command string `json:"command"`
	} `json:"content"`
}

type Home struct {
	controller.Abstract
}

func (self Home) Index(ctx *gin.Context) {
	self.JsonResponseWithoutError(ctx, "hello world!")
	return
}

func (self Home) WsNotice(http *gin.Context) {
	if !websocket.IsWebSocketUpgrade(http.Request) {
		self.JsonResponseWithError(http, errors.New("please connect using websocket"), 500)
		return
	}

	client, err := ws.NewClient(http)
	if err != nil {
		self.JsonResponseWithError(http, err, 500)
		return
	}
	go client.ReadMessage()

	// 将自己的fd推回给客户端
	err = client.SendMessage(&ws.RespMessage{
		Type: ws.MessageTypeEventFd,
		Data: client.Fd,
	})
	if err != nil {
		slog.Error("websocket", "connect", err.Error())
	}
}

func (self Home) WsContainerConsole(http *gin.Context) {
	if !websocket.IsWebSocketUpgrade(http.Request) {
		self.JsonResponseWithError(http, errors.New("please connect using websocket"), 500)
		return
	}
	type ParamsValidate struct {
		Id      string `uri:"id" binding:"required"`
		Width   uint   `form:"width"`
		Height  uint   `form:"height"`
		Cmd     string `form:"cmd,default=/bin/sh"`
		WorkDir string `form:"workDir"`
	}
	params := ParamsValidate{}
	if !self.Validate(http, &params) {
		return
	}

	containerName := params.Id
	if _, pluginName, exists := strings.Cut(params.Id, ":"); exists {
		containerName = pluginName
	}
	if params.WorkDir == "" {
		params.WorkDir = "/"
	}
	var err error
	var shell types.HijackedResponse
	var out container.ExecCreateResponse

	messageType := fmt.Sprintf(ws.MessageTypeConsole, params.Id)
	client, err := ws.NewClient(http,
		ws.WithMessageRecvHandler(messageType, func(recvMessage *ws.RecvMessage) {
			var cmd command
			err = json.Unmarshal(recvMessage.Message, &cmd)
			if err != nil {
				slog.Error("console", "json unmarshal", err.Error())
			}
			if shell.Conn == nil {
				slog.Debug("console", "shell is nil", err.Error())
				return
			}
			if cmd.Content.Command != "" {
				_, err = shell.Conn.Write([]byte(cmd.Content.Command))
				if err != nil {
					slog.Error("console", "shell read", err.Error())
				}
			}
			if cmd.Size.Height > 0 && cmd.Size.Width > 0 {
				err = docker.Sdk.Client.ContainerExecResize(docker.Sdk.Ctx, out.ID, container.ResizeOptions{
					Height: uint(cmd.Size.Height),
					Width:  uint(cmd.Size.Width),
				})
				if err != nil {
					slog.Warn("console", "container tty resize", cmd.Size, "err", err)
				}
			}
		}),
	)
	if err != nil {
		self.JsonResponseWithError(http, err, 500)
		return
	}
	go func() {
		select {
		case <-client.CtxContext.Done():
			if shell.Conn != nil {
				_ = shell.CloseWrite()
				shell.Close()
			}
			return
		}
	}()

	out, err = docker.Sdk.Client.ContainerExecCreate(client.CtxContext, containerName, container.ExecOptions{
		Privileged:   true,
		Tty:          true,
		AttachStdin:  true,
		AttachStdout: true,
		AttachStderr: true,
		Cmd: []string{
			params.Cmd,
		},
		ConsoleSize: &[2]uint{
			params.Height, params.Width,
		},
		WorkingDir: params.WorkDir,
	})
	if err != nil {
		_ = notice.Message{}.Error(".consoleError", err.Error())
		self.JsonResponseWithError(http, err, 500)
		return
	}
	shell, err = docker.Sdk.Client.ContainerExecAttach(client.CtxContext, out.ID, container.ExecStartOptions{
		Tty: true,
	})

	if err != nil {
		self.JsonResponseWithError(http, err, 500)
		return
	}

	go client.ReadMessage()
	go func() {
		out := make([]byte, 2028)
		for {
			n, err := shell.Conn.Read(out)
			if err != nil {
				return
			}
			err = client.SendMessage(&ws.RespMessage{
				Type: messageType,
				Data: string(out[:n]),
			})
			if err != nil {
				slog.Error("websocket", "shell write", err.Error())
				return
			}
		}
	}()
}

func (self Home) WsHostConsole(http *gin.Context) {
	if !websocket.IsWebSocketUpgrade(http.Request) {
		self.JsonResponseWithError(http, function.ErrorMessage(define.ErrorMessageCommonUseWsConnect), 500)
		return
	}
	type ParamsValidate struct {
		Name   string `uri:"name" binding:"required"`
		Width  int    `form:"width"`
		Height int    `form:"height"`
		Cmd    string `form:"cmd,default=/bin/sh"`
	}
	params := ParamsValidate{}
	if !self.Validate(http, &params) {
		return
	}
	if params.Name == "" {
		params.Name = docker.DefaultClientName
	}
	var err error
	var sshClient *ssh.Client
	var read io.Reader
	var write io.WriteCloser
	var session *ssh2.Session

	messageType := fmt.Sprintf(ws.MessageTypeConsoleHost, params.Name)
	client, err := ws.NewClient(http,
		ws.WithMessageRecvHandler(messageType, func(recvMessage *ws.RecvMessage) {
			var cmd command
			err = json.Unmarshal(recvMessage.Message, &cmd)
			if err != nil {
				slog.Error("console", "json unmarshal", err.Error())
			}
			if cmd.Content.Command != "" {
				_, err = write.Write([]byte(cmd.Content.Command))
				if err != nil {
					slog.Error("console", "json unmarshal", err.Error())
				}
			}
			if cmd.Size.Width > 0 && cmd.Size.Height > 0 {
				err = session.WindowChange(cmd.Size.Height, cmd.Size.Width)
				if err != nil {
					slog.Warn("console", "change size", cmd.Size, "err", err)
				}
			}
		}),
	)
	if err != nil {
		self.JsonResponseWithError(http, err, 500)
		return
	}

	err = func() error {
		dockerEnv, err := logic.Setting{}.GetDockerClient(params.Name)
		if err != nil {
			return err
		}
		if !dockerEnv.EnableSSH {
			return function.ErrorMessage(define.ErrorMessageHomeWsHostConsoleSshNotSetting)
		}
		sshClient, err = ssh.NewClient(ssh.WithServerInfo(dockerEnv.SshServerInfo)...)
		if err != nil {
			return err
		}
		session, read, write, err = sshClient.NewPtySession(params.Height, params.Width)
		if err != nil {
			return err
		}
		return nil
	}()

	if err != nil {
		_ = client.SendMessage(&ws.RespMessage{
			Type: messageType,
			Data: err.Error(),
		})
		return
	}

	go func() {
		out := make([]byte, 2028)
		for {
			n, err := read.Read(out)
			if err != nil {
				return
			}
			err = client.SendMessage(&ws.RespMessage{
				Type: messageType,
				Data: string(out[:n]),
			})
			if err != nil {
				slog.Error("websocket", "shell write", err.Error())
				return
			}
		}
	}()

	go func() {
		select {
		case <-client.CtxContext.Done():
			if sshClient != nil {
				sshClient.Close()
			}
			return
		}
	}()

	go client.ReadMessage()
}

func (self Home) Info(http *gin.Context) {
	startTime := time.Now()
	dpanelContainerInfo := container.InspectResponse{}
	new(logic.Setting).GetByKey(logic.SettingGroupSetting, logic.SettingGroupSettingDPanelInfo, &dpanelContainerInfo)
	slog.Debug("dpanel info time", "use", time.Now().Sub(startTime).String())

	startTime = time.Now()
	info, err := docker.Sdk.Client.Info(docker.Sdk.Ctx)
	if err == nil && info.ID != "" {
		info.Name = fmt.Sprintf("%s - %s", docker.Sdk.Name, docker.Sdk.Client.DaemonHost())
	}
	slog.Debug("docker info time", "use", time.Now().Sub(startTime).String())

	public, _, _ := storage.GetCertRsaContent()
	self.JsonResponseWithoutError(http, gin.H{
		"info":          info,
		"clientVersion": docker.Sdk.Client.ClientVersion(),
		"sdkVersion":    api.DefaultVersion,
		"dpanel": map[string]interface{}{
			"version":       facade.GetConfig().GetString("app.version"),
			"family":        facade.GetConfig().GetString("app.family"),
			"env":           facade.GetConfig().GetString("app.env"),
			"containerInfo": dpanelContainerInfo,
		},
		"plugin": plugin.Wrapper{}.GetPluginList(),
		"rsa": gin.H{
			"public": string(public),
		},
	})
	return
}

func (self Home) CheckNewVersion(http *gin.Context) {
	var tags []string
	var err error
	currentVersion := facade.GetConfig().GetString("app.version")
	newVersion := ""

	for _, s := range []string{
		"registry.cn-hangzhou.aliyuncs.com",
	} {
		option := make([]registry.Option, 0)
		option = append(option, registry.WithRequestCacheTime(time.Hour*24))
		option = append(option, registry.WithRegistryHost(s))
		reg := registry.New(option...)
		tags, err = reg.Repository.GetImageTagList("dpanel/dpanel")
		if err == nil {
			break
		}
	}

	for _, ver := range tags {
		if !strings.Contains(ver, "-") && version.Compare(ver, currentVersion, ">") {
			newVersion = ver
			break
		}
	}

	self.JsonResponseWithoutError(http, gin.H{
		"version":    currentVersion,
		"newVersion": newVersion,
	})
	return
}

func (self Home) Usage(http *gin.Context) {
	// 有些设备的docker获取磁盘占用比较耗时，跑一下后台协程去获取数据
	go func() {
		diskUsage, err := docker.Sdk.Client.DiskUsage(docker.Sdk.Ctx, types.DiskUsageOptions{
			Types: []types.DiskUsageObject{
				types.ContainerObject,
				types.ImageObject,
				types.VolumeObject,
				types.BuildCacheObject,
			},
		})
		if err == nil {
			// 去掉无用的信息
			for i := range diskUsage.Containers {
				diskUsage.Containers[i].Labels = make(map[string]string)
			}
			for i := range diskUsage.Images {
				diskUsage.Images[i].Labels = make(map[string]string)
			}
			for i := range diskUsage.Volumes {
				diskUsage.Volumes[i].Labels = make(map[string]string)
			}
			if !function.IsEmptyArray(diskUsage.Images) {
				sort.Slice(diskUsage.Images, func(i, j int) bool {
					return diskUsage.Images[i].Size > diskUsage.Images[j].Size
				})
			}
			if !function.IsEmptyArray(diskUsage.Containers) {
				sort.Slice(diskUsage.Containers, func(i, j int) bool {
					return diskUsage.Containers[i].SizeRw+diskUsage.Containers[i].SizeRootFs > diskUsage.Containers[j].SizeRw+diskUsage.Containers[j].SizeRootFs
				})
			}

			if !function.IsEmptyArray(diskUsage.Volumes) {
				sort.Slice(diskUsage.Volumes, func(i, j int) bool {
					if diskUsage.Volumes[i].UsageData != nil && diskUsage.Volumes[j].UsageData != nil {
						return diskUsage.Volumes[i].UsageData.Size > diskUsage.Volumes[j].UsageData.Size
					}
					return false
				})
			}

			_ = logic.Setting{}.Save(&entity.Setting{
				GroupName: logic.SettingGroupSetting,
				Name:      logic.SettingGroupSettingDiskUsage,
				Value: &accessor.SettingValueOption{
					DiskUsage: &accessor.DiskUsage{
						Usage:     &diskUsage,
						UpdatedAt: time.Now(),
					},
				},
			})
		}
		return
	}()

	diskUsage := accessor.DiskUsage{
		Usage: &types.DiskUsage{},
	}
	logic.Setting{}.GetByKey(logic.SettingGroupSetting, logic.SettingGroupSettingDiskUsage, &diskUsage)
	type portItem struct {
		Port docker.PortItem `json:"port"`
		Name string          `json:"name"`
	}
	ports := make([]*portItem, 0)
	containerList, err := docker.Sdk.Client.ContainerList(docker.Sdk.Ctx, container.ListOptions{
		All: true,
	})
	containerRunningTotal := struct {
		Stop      int `json:"stop"`
		Pause     int `json:"pause"`
		Unhealthy int `json:"unhealthy"`
	}{
		Stop:      0,
		Pause:     0,
		Unhealthy: 0,
	}
	if err == nil {
		for _, item := range containerList {
			if item.State == "exited" {
				containerRunningTotal.Stop += 1
			}
			if item.State == "paused" {
				containerRunningTotal.Pause += 1
			}
			if strings.Contains(item.Status, "unhealthy") {
				containerRunningTotal.Unhealthy += 1
			}
			usePort := make([]*portItem, 0)
			if function.IsEmptyArray(item.Ports) {
				if info, err := docker.Sdk.Client.ContainerInspect(docker.Sdk.Ctx, item.ID); err == nil && info.HostConfig != nil && !function.IsEmptyMap(info.HostConfig.PortBindings) {
					for port, bindings := range info.HostConfig.PortBindings {
						for _, binding := range bindings {
							hostPort, _ := strconv.Atoi(binding.HostPort)
							if binding.HostIP == "" {
								binding.HostIP = "0.0.0.0"
							}
							usePort = append(usePort, &portItem{
								Name: item.Names[0],
								Port: docker.PortItem{
									Host:     strconv.Itoa(hostPort),
									Dest:     strconv.Itoa(port.Int()),
									HostIp:   binding.HostIP,
									Protocol: port.Proto(),
								},
							})
						}
					}
				}
			} else {
				for _, port := range item.Ports {
					if port.PublicPort == 0 {
						continue
					}
					usePort = append(usePort, &portItem{
						Name: item.Names[0],
						Port: docker.PortItem{
							Host:   strconv.Itoa(int(port.PublicPort)),
							Dest:   strconv.Itoa(int(port.PrivatePort)),
							HostIp: port.IP,
						},
					})
				}
			}
			ports = append(ports, usePort...)
		}
		sort.Slice(ports, func(i, j int) bool {
			return ports[i].Port.Host < ports[j].Port.Host
		})
	}

	networkRow, _ := docker.Sdk.Client.NetworkList(docker.Sdk.Ctx, network.ListOptions{})
	recycleQuery := dao.Site.Where(dao.Site.DeletedAt.IsNotNull()).Unscoped()
	if containerList != nil {
		names := make([]string, 0)
		for _, summary := range containerList {
			for _, name := range summary.Names {
				names = append(names, strings.TrimPrefix(name, "/"))
			}
		}
		recycleQuery = recycleQuery.Where(dao.Site.SiteName.NotIn(names...))
	}
	containerTask, _ := recycleQuery.Count()
	imageTask, _ := dao.Image.Count()
	backupData, _ := dao.Backup.Count()

	self.JsonResponseWithoutError(http, gin.H{
		"diskUsage": diskUsage,
		"total": map[string]interface{}{
			"network":          len(networkRow),
			"containerTask":    int(containerTask),
			"containerRunning": containerRunningTotal,
			"imageTask":        int(imageTask),
			"backup":           int(backupData),
			"port":             len(ports),
		},
		"port": ports,
	})
}

func (self Home) GetStatList(http *gin.Context) {
	type ParamsValidate struct {
		Follow bool `json:"follow"`
	}
	params := ParamsValidate{}
	if !self.Validate(http, &params) {
		return
	}
	var err error
	statCmd := []string{
		"stats", "-a",
		"--format", "json",
	}
	if !params.Follow {
		statCmd = append(statCmd, "--no-stream")
		cmd, err := docker.Sdk.Run(statCmd...)
		if err != nil {
			self.JsonResponseWithError(http, err, 500)
			return
		}
		defer func() {
			_ = cmd.Close()
		}()
		out, err := cmd.RunWithResult()
		if err != nil {
			self.JsonResponseWithError(http, err, 500)
			return
		}
		list, err := logic.Stat{}.GetStat(string(out))
		if err != nil {
			self.JsonResponseWithError(http, err, 500)
			return
		}
		self.JsonResponseWithoutError(http, gin.H{
			"list": list,
		})
		return
	}

	progress, err := ws.NewFdProgressPip(http, ws.MessageTypeContainerAllStat)
	if err != nil {
		self.JsonResponseWithError(http, err, 500)
		return
	}
	time.AfterFunc(time.Hour, func() {
		progress.Close()
	})

	if progress.IsShadow() {
		self.JsonResponseWithoutError(http, gin.H{
			"list": "",
		})
		return
	}
	defer progress.Close()
	lastSendTime := time.Now()
	progress.OnWrite = func(p string) error {
		if p == "" {
			return nil
		}
		list, err := logic.Stat{}.GetStat(p)
		if err != nil {
			return err
		}
		if function.IsEmptyArray(list) {
			return nil
		}
		if time.Now().Sub(lastSendTime) > 2*time.Second {
			progress.BroadcastMessage(list)
			lastSendTime = time.Now()
		}
		return nil
	}
	cmd, err := docker.Sdk.Run(statCmd...)
	if err != nil {
		self.JsonResponseWithError(http, err, 500)
		return
	}
	defer func() {
		_ = cmd.Close()
	}()
	out, err := cmd.RunInPip()
	if err != nil {
		self.JsonResponseWithError(http, err, 500)
		return
	}
	go func() {
		slog.Debug("home get stat list progress close")
		<-progress.Done()
		_ = cmd.Close()
	}()
	_, err = io.Copy(progress, out)
	if err != nil {
		self.JsonResponseWithError(http, err, 500)
		return
	}
	self.JsonResponseWithoutError(http, gin.H{
		"list": "",
	})
	return
}

func (self Home) UpgradeScript(http *gin.Context) {
	dpanelContainerInfo := container.InspectResponse{}
	new(logic.Setting).GetByKey(logic.SettingGroupSetting, logic.SettingGroupSettingDPanelInfo, &dpanelContainerInfo)
	execPath, _ := os.Executable()
	self.JsonResponseWithoutError(http, gin.H{
		"info":     dpanelContainerInfo,
		"execPath": execPath,
	})
	return
}
