package docker

type ImageDTO struct {
	Repository string `structs:"REPOSITORY"`
	Tag        string `structs:"TAG"`
	ImageID    string `structs:"IMAGE ID"`
	Created    string `structs:"CREATED"`
	Size       string `structs:"SIZE"`
}
