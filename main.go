package main

import (
	"docker-site/controller"
	"docker-site/entity"
	"docker-site/helper"
	"docker-site/helper/template"
	"docker-site/middleware"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	tpl "html/template"
)

func main() {
	helper.InitClient("/var/run/docker.sock", "http://localhost")
	initDb()

	router := gin.Default()

	htmlTemplate := tpl.New("html_template")

	funcs := tpl.FuncMap{
		"band": template.Band,
		"bor":  template.Bor,
		"bxor": template.Bxor,
	}
	htmlTemplate.Funcs(funcs)
	htmlTemplate.ParseGlob("./templates/*.html")

	router.SetHTMLTemplate(htmlTemplate)
	router.Static("/assets", "./assets")

	store := cookie.NewStore([]byte("secret"))
	/*Middleware*/
	router.Use(sessions.Sessions("auth_cookie", store))

	/*Init Route*/
	router.GET("/", controller.ConnexionPage)
	router.POST("/login", controller.Login)

	/*Initialisation des controllers*/
	containerController := controller.ContainerController{
		Templ: htmlTemplate,
	}

	router.Use(middleware.AuthMiddleware)
	{
		router.GET("/home", controller.HomePage)
		router.GET("/docker/resume/:element", controller.GetResumeElement)
		router.GET("/docker/containers", controller.GetContainers)
		router.GET("/docker/networks", controller.GetNetworks)
		router.GET("/docker/images", controller.GetImages)
		router.GET("/docker/volumes", controller.GetVolumes)
		router.GET("/docker/container/:id", containerController.ContainerInfo)
		router.GET("/docker/container/:id/inspect", containerController.InspectContainer)
		router.GET("/docker/container/:id/buttons", containerController.ButtonContainer)
		router.POST("/docker/container/:id/:operation", containerController.HandleContainer)
		router.GET("/docker/container/:id/logs", containerController.GetLogsContainer)
		router.GET("/:page", controller.GoToPageDisplay)
	}

	router.Run("0.0.0.0:8080")
}

func initDb() error {
	db, err := helper.CreateDatabase("docker-site.sql")
	if db == nil {
		return err
	}

	var user entity.UserModel

	hashPassword := helper.HashPassword("admin")
	db.AutoMigrate(&entity.UserModel{})
	if res := db.First(&user, "username = ?", "admin"); res.RowsAffected == 0 {
		db.Create(&entity.UserModel{
			Username:       "admin",
			HashedPassword: hashPassword.Digest,
			Salt:           hashPassword.Salt,
		})
	}

	return nil
}
