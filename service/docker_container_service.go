package service

import (
	"docker-site/dto/docker"
	"docker-site/helper"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

const (
	CONTAINER_INSPECT string = "containers/:id/json"
	CONTAINER_CREATE  string = "containers/create"
	CONTAINER_STOP    string = "containers/:id/stop"
	CONTAINER_START   string = "containers/:id/start"
	CONTAINER_RESTART string = "containers/:id/restart"
	CONTAINER_KILL    string = "containers/:id/kill"
	CONTAINER_PAUSE   string = "containers/:id/pause"
	CONTAINER_UNPAUSE string = "containers/:id/unpause"
	CONTAINER_DELETE  string = "containers/:id"
	CONTAINER_LOGS    string = "containers/:id/logs?stdout=1&stderr=1"
)

const (
	START DockerCommand = iota
	RESTART
	STOP
	KILL
	PAUSE
	UNPAUSE
)

var DICT_COMMAND = map[DockerCommand]string{
	START:   CONTAINER_START,
	STOP:    CONTAINER_STOP,
	PAUSE:   CONTAINER_PAUSE,
	RESTART: CONTAINER_RESTART,
	KILL:    CONTAINER_KILL,
	UNPAUSE: CONTAINER_UNPAUSE,
}

type DockerCommand int

func DockerHandle(nameContainer string, command DockerCommand) error {
	client := helper.MakeRequest(helper.POST)
	url := strings.Replace(DICT_COMMAND[command], ":id", nameContainer, 1)
	_, err := client.Send(url, http.StatusNoContent)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func DockerExist(idContainer string) bool {
	client := helper.MakeRequest(helper.GET)
	url := strings.Replace(CONTAINER_INSPECT, ":id", idContainer, 1)

	if _, err := client.Send(url, http.StatusOK); err != nil {
		return false
	}

	return true
}

func DockerInspect(idContainer string) (*docker.ContainerInspectDTO, error) {
	var containerInspect docker.DockerContainerInspect
	client := helper.MakeRequest(helper.GET)
	url := strings.Replace(CONTAINER_INSPECT, ":id", idContainer, 1)
	result, err := client.Send(url, http.StatusOK)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(result, &containerInspect)

	if err != nil {
		return nil, err
	}
	resultDTO := containerInspect.ConvertIntoDockerInspectDTO()
	return &resultDTO, nil
}

func DockerContainerStatus(idContainer string) (docker.ContainerStatus, error) {
	var err error
	result := docker.UNKNOW

	if val, err := DockerInspect(idContainer); err == nil {
		result = val.State.Status
	}

	return result, err
}

func DockerContainerLogs(idContainer string) (string, error) {
	var result []byte
	var err error
	var logs strings.Builder
	client := helper.MakeRequest(helper.GET)
	url := strings.Replace(CONTAINER_LOGS, ":id", idContainer, 1)

	if result, err = client.Send(url, http.StatusOK); err != nil {
		return "", err
	}

	for index := 0; index < len(result); {
		res := binary.BigEndian.Uint32(result[index+4 : index+8])
		logs.WriteString(string(result[index+8 : index+int(res)]))
		logs.WriteByte('\n')
		index += int(res) + 8
	}
	return logs.String(), nil
}
