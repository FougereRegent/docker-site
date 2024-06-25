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
	GetContainerStatsStreams(id string, dataQueue chan docker.PerformanceDTO, quit chan int) error
}

const (
	CONTAINER_STATS string = "/containers/:id/stats"
)

func (o DockerPerformanceService) GetContainerStats(id string, dataQueue chan docker.PerformanceDTO, quit chan int) error {
	url := strings.ReplaceAll(CONTAINER_STATS, ":id", id)
	dataQueueRequest := make(chan string)
	quitRequest := make(chan int)

	defer close(dataQueueRequest)
	defer close(quitRequest)

	client := helper.MakeRequest(helper.GET)

	go client.ReceiveStream(url, http.StatusOK, dataQueueRequest, quitRequest)

_loop:
	select {
	case msg := <-dataQueueRequest:
		if value, err := parseJson(&msg); err != nil {
			quitRequest <- 1
		} else {
			dataQueue <- value.ToPerformanceDTO()
		}
		goto _loop
	case msg := <-quit:
		quitRequest <- msg
	}

	return nil
}

func parseJson(jsonString *string) (*docker.ContainerStats, error) {
	var containerStats docker.ContainerStats
	if err := json.Unmarshal([]byte(*jsonString), &containerStats); err != nil {
		return nil, err
	}
	return &containerStats, nil
}
