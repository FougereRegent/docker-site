package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HomeController struct{}

func (o *HomeController) HomePage(c *gin.Context) {
	c.HTML(http.StatusOK, "dashboard.html", nil)
}
