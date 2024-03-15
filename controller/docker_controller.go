package controller

import (
	"docker-site/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	CONTAINER string = "container"
	IMAGE     string = "image"
	VOLUME    string = "volume"
	NETWORK   string = "network"
)

func GetResumeElement(c *gin.Context) {
	typeElement := c.Param("element")
	fmt.Println(typeElement)

	switch typeElement {
	case CONTAINER:
		resume, err := service.GetContainerResume()
		if err != nil {
			return
		}
		c.HTML(http.StatusOK, "resume_element.html", resume)
		break
	case IMAGE:
		resume, err := service.GetImageResume()
		if err != nil {
			return
		}
		c.HTML(http.StatusOK, "resume_element.html", resume)
		break
	case NETWORK:
		resume, err := service.GetNetworkResume()
		if err != nil {
			return
		}
		c.HTML(http.StatusOK, "resume_element.html", resume)
		break
	case VOLUME:
		resume, err := service.GetVolumeResume()
		if err != nil {
			return
		}
		c.HTML(http.StatusOK, "resume_element.html", resume)
		break
	default:
		c.HTML(http.StatusNotFound, "not_found.html", nil)
		return
	}
}
