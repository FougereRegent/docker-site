package service

import (
	"docker-site/dto"
	"docker-site/helper"
	"encoding/json"
	"net/http"
)

const (
	CONTAINER_LIST         string = "/containers/json?all=true"
	CONTAINER_LIST_RUNNING string = "/containers/json?status=runnning"
	IMAGES_LIST            string = "/images/json"
	VOLUMES_LIST           string = "/volumes"
	NETWORK_LIST           string = "/networks"
)

func GetContainerResume() (*dto.ContainerResume, error) {
	var dtoContainer []dto.DockerContainer
	var dtoRunningContainer []dto.DockerContainer
	client := helper.MakeRequest(helper.GET)
	resultContainers, err := client.Send(CONTAINER_LIST, http.StatusOK)
	resultRunningContainers, err := client.Send(CONTAINER_LIST_RUNNING, http.StatusOK)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(resultContainers, &dtoContainer)
	err = json.Unmarshal(resultRunningContainers, &dtoRunningContainer)

	if err != nil {
		return nil, err
	}

	return &dto.ContainerResume{
		Resume: dto.Resume{
			Type:      dto.CONTAINER_TYPE,
			NbElement: len(dtoContainer),
		},
		NbActive: len(dtoRunningContainer),
	}, nil
}

func GetImageResume() (*dto.ImageResume, error) {
	var dtoImage []dto.DockerImage
	var totalSize float64
	client := helper.MakeRequest(helper.GET)
	result, err := client.Send(IMAGES_LIST, http.StatusOK)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(result, &dtoImage)
	if err != nil {
		return nil, err
	}

	for _, value := range dtoImage {
		if value.Size > 0 {
			totalSize += value.ConvertSize()
		}
	}

	return &dto.ImageResume{
		Resume: dto.Resume{
			Type:      dto.IMAGE_TYPE,
			NbElement: len(dtoImage),
		},
		TotalSize: totalSize,
	}, nil
}

func GetVolumeResume() (*dto.VolumeResume, error) {
	var dtoVolume dto.DockerVolume
	var resume dto.VolumeResume

	client := helper.MakeRequest(helper.GET)
	result, err := client.Send(VOLUMES_LIST, http.StatusOK)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(result, &dtoVolume)

	if err != nil {
		return nil, err
	}

	resume.Type = dto.VOLUME_TYPE
	resume.NbElement = len(dtoVolume.Volumes)

	return &resume, nil
}

func GetNetworkResume() (*dto.NetworkResume, error) {
	var dtoNetwork []dto.DockerNetwork
	var resume dto.NetworkResume

	client := helper.MakeRequest(helper.GET)
	result, err := client.Send(NETWORK_LIST, http.StatusOK)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(result, &dtoNetwork)
	if err != nil {
		return nil, err
	}
	resume.Type = dto.NETWORK_TYPE
	resume.NbElement = len(dtoNetwork)

	for _, value := range dtoNetwork {
		value.CountElement(&resume)
	}

	return &resume, nil
}

func GetContainersList() ([]dto.ContainerDTO, error) {
	var dtoDockerContainer []dto.DockerContainer

	client := helper.MakeRequest(helper.GET)
	result, err := client.Send(NETWORK_LIST, http.StatusOK)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(result, &dtoDockerContainer)
	if err != nil {
		return nil, err
	}

	dtoContainer := make([]dto.ContainerDTO, cap(dtoDockerContainer))

	for index, value := range dtoDockerContainer {
		dtoContainer[index] = value.TransformToContainerDTO()
	}

	return dtoContainer, nil
}
