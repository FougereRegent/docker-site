package controller

import (
	"docker-site/dto"
	"docker-site/service"
	"log/slog"
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

type ResumeController struct{}

func (o *ResumeController) GetResumeElement(c *gin.Context) {
	const resumeElement = "resume_element.html"
	const notFound = "not_found.html"

	typeElement := c.Param("element")

	switch typeElement {
	case CONTAINER:
		resume, err := service.GetContainerResume()
		if err != nil {
			slog.Error(err.Error())
			return
		}
		c.HTML(http.StatusOK, resumeElement, resume)
		break
	case IMAGE:
		resume, err := service.GetImageResume()
		if err != nil {
			slog.Error(err.Error())
			return
		}
		c.HTML(http.StatusOK, resumeElement, resume)
		break
	case NETWORK:
		resume, err := service.GetNetworkResume()
		if err != nil {
			slog.Error(err.Error())
			return
		}
		c.HTML(http.StatusOK, resumeElement, resume)
		break
	case VOLUME:
		resume, err := service.GetVolumeResume()
		if err != nil {
			slog.Error(err.Error())
			return
		}
		c.HTML(http.StatusOK, resumeElement, resume)
		break
	default:
		c.HTML(http.StatusNotFound, notFound, nil)
		return
	}
}

func (o *ResumeController) GoToPageDisplay(c *gin.Context) {
	typePage := c.Param("page")

	switch typePage {
	case CONTAINER:
		c.HTML(http.StatusOK, "containers.html", nil)
		break
	case IMAGE:
		c.HTML(http.StatusOK, "images.html", nil)
		break
	default:
		c.Redirect(http.StatusPermanentRedirect, "./assets/hmlt/NotFound.html")
		break
	}

}

func (o *ResumeController) GetContainers(c *gin.Context) {
	var headers []string
	var result []map[string]interface{}
	containers, err := service.GetContainersList()
	if err != nil {
		slog.Error(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}

	if len(containers) < 1 {
		goto result_func
	}

	result = make([]map[string]interface{}, len(containers))
	for index, value := range containers {
		result[index] = structs.Map(&value)
	}
	headers = structs.Names(containers[0])

result_func:
	c.HTML(http.StatusOK, "tab_component.html", gin.H{
		"tableau": dto.TabDTO{
			UrlToScan: "/docker/containers",
			Headers:   headers,
			Values:    result,
		},
	})
}

func (o *ResumeController) GetImages(c *gin.Context) {
	images, err := service.GetImagesList()
	if err != nil {
		slog.Error(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}

	result := make([]map[string]interface{}, len(images))
	for index, value := range images {
		result[index] = structs.Map(&value)
	}

	headers := structs.Names(images[0])
	c.HTML(http.StatusOK, "tab_component.html", gin.H{
		"tableau": dto.TabDTO{
			UrlToScan: "/docker/images",
			Headers:   headers,
			Values:    result,
		},
	})
}

// TODO : Implements few later, it's use to get all networks
func (o *ResumeController) GetNetworks(c *gin.Context) {

}

// TODO : Implements few later, it's use to get all Volumes
func (o *ResumeController) GetVolumes(c *gin.Context) {

}
