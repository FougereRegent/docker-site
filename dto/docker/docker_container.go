package docker

type DockerContainer struct {
	Id      string   `json:"Id"`
	Names   []string `json:"Names"`
	Image   string   `json:"Image"`
	ImageId string   `json:"ImageId"`
	State   string   `json:"State"`
	Status  string   `json:"Status"`
	Command string   `json:"Command"`
	Ports   []struct {
		PrivatePort int    `json:"PrivatePort"`
		PublicPort  int    `json:"PublicPort"`
		Type        string `json:"Type"`
	} `json:"Ports"`
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
