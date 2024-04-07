package controller

import (
	"docker-site/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

//This files groups every containers controller like this :
// - Delete
// - Creation
// - Start
// - Stop
// - Restart

// TODO : Implements deleting container in docker engine
func DeleteContainer(c *gin.Context) {

}

// TODO : Implements creating container in docker engine
func CreateContainer(c *gin.Context) {

}

// TODO : Implements starting container in docker engine
func StartContainer(c *gin.Context) {

}

// TODO : Implements stopping container in docker engine
func StopContainer(c *gin.Context) {

}

// TODO : Implements restarting container in docker engine
func RestartContainer(c *gin.Context) {

}

// TODO : Implements inspect container
func InspectContainer(c *gin.Context) {
	containerId := c.Param("id")
	containerInspect, err := service.DockerInspect(containerId)

	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}

	c.HTML(http.StatusOK, "container_inspect.html", containerInspect)
}
