package docker

import (
	"docker-site/helper"
	"strings"
)

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

func (c *DockerImage) TransformToImageDTO() ImageDTO {
	var tag string = "none"
	var repository string = "none"
	var repo []string
	if len(c.RepoTags) >= 1 {
		repo = strings.Split(c.RepoTags[0], ":")
	}

	if len(repo) == 2 {
		repository = repo[0]
		tag = repo[1]
	}

	result := ImageDTO{
		Repository: repository,
		Tag:        tag,
		ImageID:    c.Id[0:13],
		Size:       helper.OctalToStringFormat(c.Size),
		Created:    "",
	}
	return result
}

func (c *DockerImage) ConvertSize() float64 {
	return float64(c.Size) / 10e09
}
