package main

import (
	"docker-site/controller"
	"docker-site/entity"
	"docker-site/helper"
	"docker-site/helper/template"
	"docker-site/middleware"
	"docker-site/service"
	"fmt"
	tpl "html/template"
	"log/slog"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	helper.InitClient("/var/run/docker.sock", "http://localhost")
	initLoger()
	initDb()

	router := gin.Default()

	htmlTemplate := tpl.New("html_template")

	funcs := tpl.FuncMap{
		"band":       template.Band,
		"bor":        template.Bor,
		"bxor":       template.Bxor,
		"formatDate": template.FomatDate,
	}
	htmlTemplate.Funcs(funcs)

	if _, err := htmlTemplate.ParseGlob("./templates/*.html"); err != nil {
		fmt.Println(err)
		return
	}
	if _, err := htmlTemplate.ParseGlob("./templates/users/*.html"); err != nil {
		fmt.Println(err)
		return
	}
	if _, err := htmlTemplate.ParseGlob("./templates/container/*.html"); err != nil {
		fmt.Println(err)
		return
	}
	if _, err := htmlTemplate.ParseGlob("./templates/home/*.html"); err != nil {
		fmt.Println(err)
		return
	}
	if _, err := htmlTemplate.ParseGlob("./templates/components/*.html"); err != nil {
		fmt.Println(err)
		return
	}

	router.SetHTMLTemplate(htmlTemplate)
	router.Static("/assets", "./assets")

	store := cookie.NewStore([]byte("secret"))
	/*Middleware*/
	router.Use(sessions.Sessions("auth_cookie", store))

	/*Initialisation des controllers*/
	containerController := controller.ContainerController{
		Templ:              htmlTemplate,
		PerformanceService: &service.DockerPerformanceService{},
	}
	db, err := helper.GetDb()
	if err != nil {
		fmt.Println(err)
		return
	}
	userController := controller.UserController{
		UserService: &service.UserService{
			Db: db,
		},
	}
	homeController := controller.HomeController{}
	resumeController := controller.ResumeController{}

	/*Init Route*/
	router.GET("/", userController.ConnexionPage)
	router.POST("/login", userController.Login)

	router.Use(middleware.AuthMiddleware)
	{
		router.GET("/home", homeController.HomePage)
		router.GET("/:page", resumeController.GoToPageDisplay)

		dockerGroup := router.Group("/docker")
		{
			dockerGroup.GET("/resume/:element", resumeController.GetResumeElement)
			dockerGroup.GET("/containers", resumeController.GetContainers)
			dockerGroup.GET("/networks", resumeController.GetNetworks)
			dockerGroup.GET("/images", resumeController.GetImages)
			dockerGroup.GET("/volumes", resumeController.GetVolumes)
		}
		dockerContainer := router.Group("/docker/container")
		{
			dockerContainer.GET("/:id", containerController.ContainerInfo)
			dockerContainer.GET("/:id/inspect", containerController.InspectContainer)
			dockerContainer.GET("/:id/buttons", containerController.ButtonContainer)
			dockerContainer.POST("/:id/:operation", containerController.HandleContainer)
			dockerContainer.GET("/:id/logs", containerController.GetLogsContainer)
			dockerContainer.GET("/:id/stats", containerController.GetContainerPerformance)
		}
		settings := router.Group("/settings")
		{
			settings.GET("/users", userController.GetUsers)
			userManagement := settings.Group("/user")
			{
				userManagement.GET("/add", userController.AddUserPage)
				userManagement.POST("/add", userController.AddUserPage)
				userManagement.DELETE("/:id", userController.DeleteUser)
				userManagement.GET("/:id", userController.GetUserDetails)
				userManagement.GET("/:id/update", userController.UpdateUserDetails)
				userManagement.PUT("/:id", userController.UpdateUserDetails)
				userManagement.GET("/:id/password", userController.UpdateUserPassword)
				userManagement.PUT("/:id/password", userController.UpdateUserPassword)
			}
		}
	}

	router.Run("0.0.0.0:8080")
}

func initDb() error {
	var user entity.UserModel
	db, err := helper.CreateDatabase("./docker-site.db")
	if db == nil {
		slog.Error(err.Error())
		return err
	}

	hashPassword := helper.HashPassword("admin")
	db.AutoMigrate(&entity.UserModel{}, &entity.UserDetailsModel{}, &entity.UserConnectionModel{})

	if res := db.Find(&user, "username = ?", "admin"); res.RowsAffected == 0 {
		admin := entity.UserModel{
			Username:       "admin",
			HashedPassword: hashPassword.Digest,
			Salt:           hashPassword.Salt,
			UserDetails: entity.UserDetailsModel{
				FirstName: "admin",
				LastName:  "admin",
			},
		}

		if ctx := db.Create(&admin); ctx.Error != nil {
			slog.Error(ctx.Error.Error())
		}
		if ctx := db.Save(&admin); ctx.Error != nil {
			slog.Error(ctx.Error.Error())
		}
	}

	return nil
}

func initLoger() {
	if file, err := os.OpenFile("docker-site.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	} else {
		logger := slog.New(slog.NewTextHandler(file, nil))
		slog.SetDefault(logger)
	}
}
