package docker

type ContainerInspectDTO struct {
	Id      string
	Image   string
	Name    string
	Path    string
	Args    []string
	Created string
	State   ContainerStatusDTO
	Running bool
	Mount   []ContainerMountsPointDTO
}
