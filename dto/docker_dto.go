package dto

import (
	"time"
)

const (
	CONTAINER_TYPE ElementType = "Container"
	IMAGE_TYPE     ElementType = "Image"
	VOLUME_TYPE    ElementType = "Volume"
	NETWORK_TYPE   ElementType = "Network"
)

type ElementType string

type Resume struct {
	Type      ElementType
	NbElement int
}

type ContainerResume struct {
	Resume
	NbActive int
}

type ImageResume struct {
	Resume
	TotalSize float64
}

type VolumeResume struct {
	Resume
}

type NetworkResume struct {
	Resume
	NbDriver  int
	NbHost    int
	NbBridge  int
	NbNone    int
	NbOverlay int
	NbIpVlan  int
	NbMacVlan int
}

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
