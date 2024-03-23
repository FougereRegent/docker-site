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

type DockerContainer struct {
	Id      string       `json:"Id"`
	Names   []string     `json:"Names"`
	Image   string       `json:"Image"`
	ImageId string       `json:"ImageId"`
	State   string       `json:"State"`
	Status  string       `json:"Status"`
	Command string       `json:"Command"`
	Ports   []DockerPort `json:"Ports"`
}

type DockerPort struct {
	PrivatePort int    `json:"PrivatePort"`
	PublicPort  int    `json:"PublicPort"`
	Type        string `json:"Type"`
}

type DockerImage struct {
	Id         string            `json:"Id"`
	ParentId   string            `json:"ParentId"`
	RepoTags   []string          `json:"RepoTags"`
	RepoDigest []string          `json:"RepoDigest"`
	Created    int               `json:"Created"`
	Size       int64             `json:"Size"`
	Labels     map[string]string `json:"Labels"`
	Containers int               `json:"Containers"`
}

type DockerNetwork struct {
	Name       string    `json:"Name"`
	Id         string    `json:"Id"`
	Created    time.Time `json:"Created"`
	Scope      string    `json:"Scope"`
	Driver     string    `json:"Driver"`
	EnableIPv6 bool      `json:"EnableIPv6"`
	Internal   bool      `json:"Internal"`
	Attachable bool      `json:"Attachable"`
	Ingress    bool      `json:"Ingress"`
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

type ContainerDTO struct {
	Hash    string
	Name    string
	Image   string
	Command string
	Statut  string
}

func (c *DockerImage) ConvertSize() float64 {
	return float64(c.Size) / 10e09
}

func (c *DockerNetwork) CountElement(network *NetworkResume) {
	switch c.Driver {
	case "host":
		network.NbHost += 1
		break
	case "bridge":
		network.NbBridge += 1
		break
	case "none":
		network.NbNone += 1
		break
	case "overlay":
		network.NbOverlay += 1
		break
	case "macvlan":
		network.NbMacVlan += 1
		break
	case "ipvlan":
		network.NbIpVlan += 1
		break
	default:
		return
	}
}

func (c *DockerContainer) TransformToContainerDTO() ContainerDTO {
	result := ContainerDTO{
		Hash:    c.Id[0:13],
		Name:    c.Names[0],
		Image:   c.Image,
		Statut:  c.State,
		Command: c.Command,
	}
	return result
}
