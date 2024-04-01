package docker

import (
	"time"
)

type DockerVolume struct {
	Volumes []struct {
		Name       string                 `json:"Name"`
		Driver     string                 `json:"Driver"`
		MountPoint string                 `json:"MountPoint"`
		CreatedAt  time.Time              `json:"CreatedAt"`
		Status     map[string]interface{} `json:"Status"`
		Labels     map[string]interface{} `json:"Labels"`
		Scope      string                 `json:"Scope"`
	} `json:"Volumes"`
}
