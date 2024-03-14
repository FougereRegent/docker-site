package service

import (
	"docker-site/dto"
	"docker-site/helper"
)

const (
	CONTAINER_LIST string = "http://localhost/containers/json"
	IMAGES_LIST    string = "http://localhost/images/json"
	VOLUMES_LIST   string = "http://localhost/volumes/json"
	NETWORK_LIST   string = "http://localhost/network/json"
)

func GetContainerResume() (*dto.ContainerResume, error) {
	client, err := helper.GetClient()
	if client == nil {
		return nil, err
	}

	return nil, nil
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
