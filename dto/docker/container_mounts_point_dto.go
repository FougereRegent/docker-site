package docker

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
