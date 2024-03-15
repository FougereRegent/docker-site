package service

import (
	"docker-site/dto"
	"docker-site/helper"
	"encoding/json"
	"net/http"
)

const (
	CONTAINER_LIST string = "/containers/json"
	IMAGES_LIST    string = "/images/json"
	VOLUMES_LIST   string = "/volumes/json"
	NETWORK_LIST   string = "/network/json"
)

func GetContainerResume() (*dto.ContainerResume, error) {
	var dtoContainer []dto.DockerContainer
	client := helper.MakeRequest(helper.GET)
	result, err := client.Send(CONTAINER_LIST, http.StatusOK)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(result, &dtoContainer)

	if err != nil {
		return nil, err
	}

	return &dto.ContainerResume{}, nil
}

func GetImageResume() (*dto.ImageResume, error) {
	return nil, nil
}

func GetVolumeResume() (*dto.VolumeResume, error) {
	return nil, nil
}

func GetNetworkResume() (*dto.NetworkResume, error) {
	return nil, nil
}
