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

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	helper.InitClient("/var/run/docker.sock", "http://localhost")
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
		Templ: htmlTemplate,
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
		router.GET("/docker/resume/:element", resumeController.GetResumeElement)
		router.GET("/docker/containers", resumeController.GetContainers)
		router.GET("/docker/networks", resumeController.GetNetworks)
		router.GET("/docker/images", resumeController.GetImages)
		router.GET("/docker/volumes", resumeController.GetVolumes)
		router.GET("/:page", resumeController.GoToPageDisplay)
		router.GET("/docker/container/:id", containerController.ContainerInfo)
		router.GET("/docker/container/:id/inspect", containerController.InspectContainer)
		router.GET("/docker/container/:id/buttons", containerController.ButtonContainer)
		router.POST("/docker/container/:id/:operation", containerController.HandleContainer)
		router.GET("/docker/container/:id/logs", containerController.GetLogsContainer)
		router.GET("/settings/users", userController.GetUsers)
		router.GET("/settings/user/add", userController.AddUserPage)
		router.POST("/settings/user/add", userController.AddUserPage)
		router.DELETE("/settings/user/:id", userController.DeleteUser)
		router.GET("/settings/user/:id", userController.GetUserDetails)
		router.GET("/settings/user/:id/update", userController.UpdateUserDetails)
		router.PUT("/settings/user/:id", userController.UpdateUserDetails)
		router.GET("/settings/user/:id/password", userController.UpdateUserPassword)
		router.PUT("/settings/user/:id/password", userController.UpdateUserPassword)
	}

	router.Run("0.0.0.0:8080")
}

func initDb() error {
	db, err := helper.CreateDatabase("./docker-site.db")
	if db == nil {
		return err
	}

	var user entity.UserModel

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

		db.Create(&admin)
		db.Save(&admin)
	}

	return nil
}
