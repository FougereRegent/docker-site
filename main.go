package main

import (
	"docker-site/controller"
	"docker-site/entity"
	"docker-site/helper"
	"docker-site/middleware"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	helper.InitClient("/var/run/docker.sock")
	initDb()
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Static("/assets", "./assets")

	store := cookie.NewStore([]byte("secret"))
	/*Middleware*/
	router.Use(sessions.Sessions("auth_cookie", store))

	/*Init Route*/
	router.GET("/", controller.ConnexionPage)
	router.POST("/login", controller.Login)

	router.Use(middleware.AuthMiddleware)
	{
		router.GET("/home", controller.HomePage)
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
