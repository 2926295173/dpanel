package logic

import (
	"archive/zip"
	"github.com/donknap/dpanel/common/accessor"
	"github.com/donknap/dpanel/common/function"
	"github.com/donknap/dpanel/common/service/docker"
	"github.com/donknap/dpanel/common/service/exec/local"
	"github.com/donknap/dpanel/common/service/storage"
	"github.com/donknap/dpanel/common/types/define"
	"gopkg.in/yaml.v3"
	"io"
	"io/fs"
	"log/slog"
	"net/http"
	"os"
	exec2 "os/exec"
	"path/filepath"
	"strings"
	"time"
)

type StoreLogoFileSystem struct {
	fs.FS
}

func (self StoreLogoFileSystem) Open(name string) (fs.File, error) {
	return os.Open(filepath.Join(storage.Local{}.GetStorePath(), name))
}

type Store struct {
}

func (self Store) SyncByGit(path, gitUrl string) error {
	if _, err := exec2.LookPath("git"); err != nil {
		return function.ErrorMessage(define.ErrorMessageSystemStoreNotFoundGit)
	}
	// 先创建一个临时目录，下载完成后再同步数据，否则失败时原先的数据会被删除
	tempDownloadPath, _ := storage.Local{}.CreateTempDir("")
	defer func() {
		_ = os.RemoveAll(tempDownloadPath)
	}()
	slog.Debug("store git download", "path", tempDownloadPath)

	cmd, err := local.New(
		local.WithCommandName("git"),
		local.WithArgs("clone", "--depth", "1",
			gitUrl, tempDownloadPath),
	)
	if err != nil {
		return err
	}
	time.AfterFunc(time.Minute*5, func() {
		_ = cmd.Close()
	})
	_, err = cmd.RunWithResult()
	if err != nil {
		return err
	}
	err = os.RemoveAll(path)
	if err != nil {
		return err
	}
	err = os.CopyFS(path, os.DirFS(tempDownloadPath))
	if err != nil {
		return err
	}
	return nil
}

func (self Store) SyncByZip(path, zipUrl string, root string) error {
	zipTempFile, _ := storage.Local{}.CreateTempFile("")
	defer func() {
		_ = zipTempFile.Close()
		_ = os.RemoveAll(zipTempFile.Name())
	}()

	response, err := http.Get(zipUrl)
	if err != nil {
		return err
	}
	defer func() {
		_ = response.Body.Close()
	}()
	if response.StatusCode != http.StatusOK {
		return function.ErrorMessage(define.ErrorMessageSystemStoreDownloadFailed, "url", zipUrl, "error", response.Status)
	}
	_, err = io.Copy(zipTempFile, response.Body)
	if err != nil {
		return err
	}
	zipArchive, err := zip.OpenReader(zipTempFile.Name())
	if err != nil {
		return err
	}
	defer func() {
		_ = zipArchive.Close()
	}()
	for _, file := range zipArchive.File {
		if file.FileInfo().IsDir() {
			continue
		}
		if strings.HasPrefix(file.Name, "__MACOSX") {
			continue
		}
		targetFilePath := filepath.Join(path, file.Name)
		if root != "" {
			if before, after, exists := strings.Cut(file.Name, root); exists {
				if before != "" {
					targetFilePath = filepath.Join(path, root, after)
				}
			} else {
				continue
			}
		}
		err = os.MkdirAll(filepath.Dir(targetFilePath), os.ModePerm)
		if err != nil {
			return err
		}
		targetFile, err := os.OpenFile(targetFilePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			return err
		}
		sourceFile, _ := file.Open()
		_, _ = io.Copy(targetFile, sourceFile)
	}
	return nil
}

func (self Store) SyncByJson(path, jsonUrl string) error {
	file, err := os.OpenFile(filepath.Join(path, "template.json"), os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0666)
	defer func() {
		_ = file.Close()
	}()

	response, err := http.Get(jsonUrl)
	if err != nil {
		return err
	}
	defer func() {
		_ = response.Body.Close()
	}()

	if response.StatusCode != http.StatusOK {
		return function.ErrorMessage(define.ErrorMessageSystemStoreDownloadFailed, "url", jsonUrl, "error", response.Status)
	}
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}
	return nil
}

func (self Store) GetAppByCasaos(storePath string) ([]accessor.StoreAppItem, error) {
	if !filepath.IsAbs(storePath) {
		storePath = filepath.Join(storage.Local{}.GetStorePath(), storePath, "Apps")
	}
	result := make([]accessor.StoreAppItem, 0)

	err := filepath.WalkDir(storePath, func(path string, d fs.DirEntry, err error) error {
		if path == storePath {
			return nil
		}
		if !d.IsDir() {
			return nil
		}

		appName, _ := filepath.Rel(storePath, path)
		appPath := filepath.Join(storePath, appName)

		storeItem := accessor.StoreAppItem{
			Name:     appName,
			Version:  make(map[string]accessor.StoreAppVersionItem),
			Contents: make(map[string]string),
		}

		composeYaml, err := os.ReadFile(filepath.Join(appPath, "docker-compose.yml"))
		if err != nil {
			return err
		}
		yamlData := new(function.YamlGetter)
		err = yaml.Unmarshal(composeYaml, &yamlData)
		if err != nil {
			return err
		}
		storeItem.Description = yamlData.GetString("x-casaos.description.zh_cn")
		storeItem.Descriptions = map[string]string{
			define.LangZh: yamlData.GetString("x-casaos.description.zh_cn"),
			define.LangEn: yamlData.GetString("x-casaos.description.en_us"),
		}
		storeItem.Tag = []string{
			yamlData.GetString("x-casaos.category"),
		}
		storeItem.Logo = yamlData.GetString("x-casaos.icon")
		if v := yamlData.GetString("x-casaos.tips.before_install.zh_cn"); v != "" {
			storeItem.Content = "markdown-file://" + v
			storeItem.Contents[define.LangZh] = "markdown://" + v
		}
		if v := yamlData.GetString("x-casaos.tips.before_install.en_us"); v != "" {
			storeItem.Contents[define.LangEn] = "markdown://" + v
		}
		resourcePath, _ := filepath.Rel(filepath.Dir(filepath.Dir(storePath)), appPath)
		storeItem.Version["latest"] = accessor.StoreAppVersionItem{
			Name:        "latest",
			ComposeFile: filepath.Join(resourcePath, "docker-compose.yml"),
			Environment: make([]docker.EnvItem, 0),
		}
		if err == nil {
			result = append(result, storeItem)
		}
		return filepath.SkipDir
	})

	if err != nil {
		return nil, err
	}
	return result, nil
}

func (self Store) ParseSettingField(field map[string]string, call func(item *docker.EnvValueRule)) *docker.EnvValueRule {
	valueRule := &docker.EnvValueRule{}

	if field["required"] == "true" {
		valueRule.Kind |= docker.EnvValueRuleRequired
	}
	if field["disabled"] == "true" {
		valueRule.Kind |= docker.EnvValueRuleDisabled
	}

	switch field["type"] {
	case "text":
		valueRule.Kind |= docker.EnvValueTypeText
		break
	case "number":
		valueRule.Kind |= docker.EnvValueTypeNumber
		break
	case "select":
		if field["multiple"] == "true" {
			valueRule.Kind |= docker.EnvValueTypeSelectMultiple
		} else {
			valueRule.Kind |= docker.EnvValueTypeSelect
		}
	}

	if call != nil {
		call(valueRule)
	}
	return valueRule
}
