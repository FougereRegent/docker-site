package docker

type MountType string
type ContainerStatus string

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
	UNKNOW     ContainerStatus = "unknow"
)
