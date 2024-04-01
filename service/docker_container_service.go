package service

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

type DockerCommand int

// TODO: Gestions des containers : Start, Kill, Pause, Stop, Restart
func DockerHandle(idContainer string, command DockerCommand) error {
	return nil
}
