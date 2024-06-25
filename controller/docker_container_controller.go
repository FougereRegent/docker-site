package controller

import (
	"bytes"
	"docker-site/dto/docker"
	"docker-site/service"
	"html/template"
	"log/slog"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var status = map[docker.ContainerStatus]int8{
	docker.RUNNING:    0b00010110,
	docker.CREATED:    0b00000000,
	docker.PAUSED:     0b00101010,
	docker.RESTARTING: 0b00000000,
	docker.REMOVING:   0b00000000,
	docker.EXITED:     0b00000001,
	docker.DEAD:       0b00000011,
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var polling = time.Second * 1

type ContainerController struct {
	Templ              *template.Template
	performanceService service.IDockerPerformanceService
}

func (o *ContainerController) HandleContainer(c *gin.Context) {
	var op service.DockerCommand
	operation := c.Param("operation")
	id := c.Param("id")

	if id == "" || operation == id {
		c.Status(http.StatusBadRequest)
		return
	}

	switch operation {
	case "start":
		op = service.START
		break
	case "restart":
		op = service.RESTART
		break
	case "stop":
		op = service.STOP
		break
	case "pause":
		op = service.PAUSE
		break
	case "unpause":
		op = service.UNPAUSE
		break
	case "kill":
		op = service.KILL
		break
	default:
		c.Status(http.StatusBadRequest)
		break
	}

	if err := service.DockerHandle(id, op); err != nil {
		slog.Error(err.Error())
		c.Status(http.StatusBadRequest)
	} else {
		c.Status(http.StatusCreated)
	}

}

func (o *ContainerController) ContainerInfo(c *gin.Context) {
	containerId := c.Param("id")
	_, err := service.DockerInspect(containerId)
	slog.Info("Get container info from ", "container id", containerId)
	if err != nil {
		slog.Error(err.Error())
		c.Redirect(http.StatusPermanentRedirect, "/home")
		return
	}

	c.HTML(http.StatusOK, "container_info.html", gin.H{
		"Name": containerId,
	})
}

func (o *ContainerController) InspectContainer(c *gin.Context) {
	var conn *websocket.Conn
	var err error

	containerId := c.Param("id")
	slog.Info("Get container inspect from ", "container id", containerId)
	containerExist := service.DockerExist(containerId)

	if !containerExist {
		slog.Error("Container not found", "container id", containerId)
		c.Status(http.StatusNotFound)
		return
	}

	if conn, err = upgrader.Upgrade(c.Writer, c.Request, nil); err != nil {
		slog.Error(err.Error())
		c.Redirect(http.StatusPermanentRedirect, "/home")
		return
	}

	defer conn.Close()

	for {
		var buffer bytes.Buffer
		var containerInspect *docker.ContainerInspectDTO

		if containerInspect, err = service.DockerInspect(containerId); err != nil {
			slog.Error(err.Error())
			c.Redirect(http.StatusPermanentRedirect, "/home")
			return
		}

		o.Templ.ExecuteTemplate(&buffer, "container_inspect.html", containerInspect)
		conn.WriteMessage(websocket.TextMessage, []byte(buffer.String()))
		time.Sleep(polling)
	}
}

func (o *ContainerController) ButtonContainer(c *gin.Context) {
	var conn *websocket.Conn
	var err error

	containerId := c.Param("id")
	containerExist := service.DockerExist(containerId)
	if !containerExist {
		c.Status(http.StatusNotFound)
		return
	}

	if conn, err = upgrader.Upgrade(c.Writer, c.Request, nil); err != nil {
		slog.Error(err.Error())
		c.Redirect(http.StatusPermanentRedirect, "/home")
		return
	}

	defer conn.Close()

	for {
		var buffer bytes.Buffer
		var containerStatus docker.ContainerStatus

		if containerStatus, _ = service.DockerContainerStatus(containerId); containerStatus == docker.UNKNOW {
			c.Redirect(http.StatusPermanentRedirect, "/home")
			break
		}

		o.Templ.ExecuteTemplate(&buffer, "container_button.html", gin.H{
			"Name":         containerId,
			"VectorButton": status[containerStatus],
		})
		conn.WriteMessage(websocket.TextMessage, []byte(buffer.String()))
		time.Sleep(polling)
	}
}

func (o *ContainerController) GetLogsContainer(c *gin.Context) {
	var conn *websocket.Conn
	var err error

	containerId := c.Param("id")

	if containerExist := service.DockerExist(containerId); !containerExist {
		c.Status(http.StatusNotFound)
		return
	}

	if conn, err = upgrader.Upgrade(c.Writer, c.Request, nil); err != nil {
		slog.Error(err.Error())
		c.Redirect(http.StatusPermanentRedirect, "/home")
		return
	}

	defer conn.Close()

	for {
		var buffer bytes.Buffer
		var logs string
		if logs, err = service.DockerContainerLogs(containerId); err != nil {
			c.Redirect(http.StatusPermanentRedirect, "/home")
			break
		}

		o.Templ.ExecuteTemplate(&buffer, "container_logs.html", gin.H{
			"Logs": logs,
		})

		conn.WriteMessage(websocket.TextMessage, buffer.Bytes())
		time.Sleep(polling)
	}
}

func (o *ContainerController) GetContainerPerformance(c *gin.Context) {
	var conn *websocket.Conn
	var err error

	containerId := c.Param("id")

	if containerExist := service.DockerExist(containerId); !containerExist {
		c.Status(http.StatusNotFound)
		return
	}

	if conn, err = upgrader.Upgrade(c.Writer, c.Request, nil); err != nil {
		slog.Error(err.Error())
		c.Redirect(http.StatusPermanentRedirect, "/home")
	}

	dataQueue := make(chan docker.PerformanceDTO)
	quitQueue := make(chan int)

	defer func() {
		quitQueue <- 1
	}()

	defer close(dataQueue)
	defer close(quitQueue)
	defer conn.Close()

	go o.performanceService.GetContainerStatsStreams(containerId, dataQueue, quitQueue)

	for {
	}
}
