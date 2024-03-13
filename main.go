package main

import (
	"docker-site/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Static("/assets", "./assets")

	/*Init Route*/
	router.GET("/", controller.ConnexionPage)

	router.Run("0.0.0.0:8080")
}
