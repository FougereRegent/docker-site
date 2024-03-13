package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {

}

func DeleteUser(c *gin.Context) {

}

func UpdateUser(c *gin.Context) {

}

func GetUsers(c *gin.Context) {

}

func Login(c *gin.Context) {

}

func ConnexionPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login_page.html", nil)
}
