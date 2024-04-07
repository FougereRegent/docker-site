package docker

import "time"

const (
	BIND    MountType = "bind"
	VOLUME  MountType = "volume"
	TMPFS   MountType = "tmpfs"
	NPIPE   MountType = "npipe"
	CLUSTER MountType = "cluster"
)

const (
	CREATED    ContainerStatus = "created"
	RUNNING    ContainerStatus = "running"
	PAUSED     ContainerStatus = "paused"
	RESTARTING ContainerStatus = "restarting"
	REMOVING   ContainerStatus = "removing"
	EXITED     ContainerStatus = "exited"
	DEAD       ContainerStatus = "dead"
)

type MountType string
type ContainerStatus string

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

type ContainerDTO struct {
	Hash    string
	Name    string
	Image   string
	Command string
	Statut  string
}

type ContainerInspectDTO struct {
	Id      string
	Image   string
	Name    string
	Path    string
	Args    []string
	Created time.Time
	State   ContainerStatusDTO
	Running bool
	Mount   []ContainerMountsPointDTO
}

type ContainerMountsPointDTO struct {
	Type         MountType
	Name         string
	Source       string
	Destincation string
	Driver       string
	Mode         string
	RW           bool
	Propagation  string
}

type ContainerStatusDTO struct {
	Status     ContainerStatus
	PID        int
	StartedAt  time.Time
	FinishedAT time.Time
}

func (c *DockerContainer) TransformToContainerDTO() ContainerDTO {
	result := ContainerDTO{
		Hash:    c.Id[0:13],
		Name:    c.Names[0][1:len(c.Names[0])],
		Image:   c.Image,
		Statut:  c.State,
		Command: c.Command,
	}
	return result
}
