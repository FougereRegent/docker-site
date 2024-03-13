package controller

import (
	"crypto/sha256"
	"docker-site/dto"
	"docker-site/entity"
	"docker-site/helper"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	"net/http"
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
	var user dto.UserFrontDTO
	var userModel entity.UserModel
	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusUnauthorized, "error")
		return
	}

	db, err := helper.GetDb()
	if db == nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	if res := db.First(&userModel, "username = ?", user.Username); res.RowsAffected == 0 {
		c.HTML(http.StatusOK, "login_error.html", nil)
		return
	}

	if !helper.CheckPassword(user.Password, &helper.HashedPassword{
		Digest: userModel.HashedPassword,
		Salt:   userModel.Salt}) {

		c.HTML(http.StatusOK, "login_error.html", nil)
		return
	}

	sha := sha256.New()
	sha.Write([]byte(user.Password))

	sessionToken := xid.New().String()
	session := sessions.Default(c)
	session.Set("username", user.Username)
	session.Set("token", sessionToken)
	session.Options(sessions.Options{
		MaxAge: 30 * 60,
	})

	if err := session.Save(); err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Writer.Header().Add("HX-Redirect", "/home")
	c.Status(http.StatusOK)
}

func Logout(c *gin.Context) {

}

func ConnexionPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login_page.html", nil)
}
