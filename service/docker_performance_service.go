package service

import (
	"docker-site/dto/docker"
	"docker-site/helper"
	"encoding/json"
	"net/http"
	"strings"
)

type DockerPerformanceService struct {
}

type IDockerPerformanceService interface {
	GetContainerStatsStreams(id string, dataQueue chan docker.PerformanceDTO, quit chan int)
}

const (
	CONTAINER_STATS string = "/containers/:id/stats?stream=true"
)

func (o *DockerPerformanceService) GetContainerStatsStreams(id string, dataQueue chan docker.PerformanceDTO, quit chan int) {
	var closeLoop bool = false
	stringQueue := make(chan string)
	leaveQueue := make(chan int)

	url := strings.Replace(CONTAINER_STATS, ":id", id, 1)
	request := helper.MakeRequest(helper.GET)

	defer close(leaveQueue)
	defer close(stringQueue)

	go request.ReceiveStream(url, http.StatusOK, stringQueue, leaveQueue)

	for !closeLoop {
		select {
		case <-quit:
			leaveQueue <- 1
			closeLoop = true
		case data := <-stringQueue:
			if err := writeIntoChan([]byte(data), dataQueue); err != nil {
				return
			}
		}
	}
	return
}

func writeIntoChan(data []byte, channel chan docker.PerformanceDTO) error {
	var result docker.ContainerStats
	if err := json.Unmarshal(data, &result); err != nil {
		return err
	}

	channel <- result.ToPerformanceDTO()
	return nil
}

func parseJson(jsonString *string) (*docker.ContainerStats, error) {
	var containerStats docker.ContainerStats
	if err := json.Unmarshal([]byte(*jsonString), &containerStats); err != nil {
		return nil, err
	}
	return &containerStats, nil
}
