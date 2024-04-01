package docker

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
