package service

import (
	"docker-site/helper"
	"net/http"
	"strings"
)

const (
	CONTAINER_INSPECT string = "/containers/:id/json"
	CONTAINER_CREATE  string = "/containers/create"
	CONTAINER_STOP    string = "/containers/:id/stop"
	CONTAINER_START   string = "/containers/:id/start"
	CONTAINER_RESTART string = "/containers/:id/restart"
	CONTAINER_KILL    string = "/containers/:id/kill"
	CONTAINER_PAUSE   string = "/containers/:id/pause"
	CONTAINER_DELETE  string = "/containers/:id"
)

const (
	START DockerCommand = iota
	RESTART
	STOP
	KILL
	PAUSE
)

var DICT_COMMAND = map[DockerCommand]string{
	START:   CONTAINER_START,
	STOP:    CONTAINER_STOP,
	PAUSE:   CONTAINER_PAUSE,
	RESTART: CONTAINER_RESTART,
	KILL:    CONTAINER_KILL,
}

type DockerCommand int

// TODO: Gestions des containers : Start, Kill, Pause, Stop, Restart
func DockerHandle(idContainer string, command DockerCommand) error {
	client := helper.MakeRequest(helper.POST)
	url := strings.Replace(DICT_COMMAND[command], ":id", idContainer, 1)
	_, err := client.Send(url, http.StatusNoContent)

	if err != nil {
		return err
	}

	return nil
}
