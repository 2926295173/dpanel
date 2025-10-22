package accessor

import "github.com/donknap/dpanel/common/service/docker"

type ImageSettingOption struct {
	ImageId                string                `json:"imageId"`
	Tag                    string                `json:"tag,omitempty"`
	Registry               string                `json:"registry,omitempty"`
	BuildType              string                `json:"buildType,omitempty"`
	BuildDockerfileContent string                `json:"buildDockerfileContent" binding:"omitempty"`
	BuildDockerfileName    string                `json:"buildDockerfileName"`
	BuildDockerfileRoot    string                `json:"buildDockerfileRoot"`
	BuildGit               string                `json:"buildGit"`
	BuildZip               string                `json:"buildZip"`
	BuildArgs              []docker.EnvItem      `json:"buildArgs,omitempty"`
	Platform               *docker.ImagePlatform `json:"platform,omitempty"`
	BuildDockerfile        string                `json:"buildDockerfile,omitempty,deprecated"` // Deprecated: instead BuildDockerfileContent
	BuildRoot              string                `json:"buildRoot,omitempty,deprecated"`       // Deprecated: instead BuildDockerfileRoot
}
