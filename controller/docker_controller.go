package controller

import (
	"docker-site/dto"
	"docker-site/service"
	"fmt"
	"net/http"

	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
)

const (
	CONTAINER string = "containers"
	IMAGE     string = "images"
	VOLUME    string = "volumes"
	NETWORK   string = "networks"
)

func GetResumeElement(c *gin.Context) {
	typeElement := c.Param("element")

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

func GoToPageDisplay(c *gin.Context) {
	typePage := c.Param("page")

	switch typePage {
	case CONTAINER:
		c.HTML(http.StatusOK, "containers.html", nil)
		break
	default:
		c.Redirect(http.StatusPermanentRedirect, "./assets/hmlt/NotFound.html")
		break
	}

}

func GetContainers(c *gin.Context) {
	containers, err := service.GetContainersList()
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	result := make([][]interface{}, len(containers))
	for index, value := range containers {
		result[index] = structs.Values(&value)
	}

	headers := structs.Names(containers[0])
	fmt.Println(result[0])

	c.HTML(http.StatusOK, "tab_component.html", gin.H{
		"tableau": dto.TabDTO{
			UrlToScan: "/docker/containers",
			Headers:   headers,
			Values:    result,
		},
	})
}
