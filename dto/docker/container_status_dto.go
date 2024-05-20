package docker

type ContainerStatusDTO struct {
	Status     ContainerStatus
	PID        int
	StartedAt  string
	FinishedAt string
}
