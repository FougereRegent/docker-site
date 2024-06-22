package controller

import (
	"crypto/sha256"
	"docker-site/dto"
	"docker-site/dto/user"
	"docker-site/service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
)

type UserController struct {
	UserService service.IUserService
}

func (o *UserController) DeleteUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	if err = o.UserService.DeleteUser(uint(id)); err != nil {
		c.Status(http.StatusNotFound)
		return
	}
	c.Header("HX-Refresh", "true")
	c.Status(http.StatusOK)
}

func (o *UserController) UpdateUserPassword(c *gin.Context) {
	service := o.UserService
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return
	}

	if c.Request.Method == "GET" {
		c.HTML(http.StatusOK, "update_password.html", gin.H{
			"Id": id,
		})
	} else if c.Request.Method == "PUT" {
		var passwordUpdate user.UserUpdatePasswordDTO
		c.Bind(&passwordUpdate)
		err := service.UpdatePassword(uint(id), passwordUpdate)
		fmt.Println(err)
		c.Header("HX-Redirect", fmt.Sprintf("/settings/user/%d", id))
	}
}

func (o *UserController) GetUsers(c *gin.Context) {
	users, err := o.UserService.GetAllUsers()
	if err != nil {
		c.Status(http.StatusInternalServerError)
	} else {
		c.HTML(http.StatusOK, "all_users.html", gin.H{
			"Users": users,
		})
	}
}

func (o *UserController) Login(c *gin.Context) {
	var user dto.UserFrontDTO
	service := o.UserService

	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusUnauthorized, "error")
		return
	}

	if err := service.SetConnection(user.Username, user.Password); err != nil {
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

func (o *UserController) Logout(c *gin.Context) {

}

func (o *UserController) ConnexionPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login_page.html", nil)
}

func (o *UserController) AddUserPage(c *gin.Context) {
	service := o.UserService
	if c.Request.Method == "GET" {
		c.HTML(http.StatusOK, "add_user_page.html", nil)
	} else if c.Request.Method == "POST" {
		var userCreationDTO user.UserCreationDTO

		if err := c.Bind(&userCreationDTO); err != nil {
			return
		}

		if err := service.CreateUser(userCreationDTO); err != nil {
			c.HTML(http.StatusOK, "error_create_user.html", gin.H{
				"Error": err.Error(),
			})
			return
		}

		c.Header("HX-Redirect", "/settings/users")
		c.Status(http.StatusCreated)

	} else {
		c.HTML(http.StatusNotFound, "not_found.html", nil)
	}
}

func (o *UserController) GetUserDetails(c *gin.Context) {
	service := o.UserService
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.HTML(http.StatusNotFound, "not_found.html", nil)
		return
	}

	if userDetails, err := service.GetUserDetails(uint(id)); err != nil {
		c.HTML(http.StatusNotFound, "not_found.html", nil)
	} else {
		c.HTML(http.StatusFound, "user_details.html", userDetails)
	}
}

func (o *UserController) UpdateUserDetails(c *gin.Context) {
	service := o.UserService
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.HTML(http.StatusNotFound, "not_found.html", nil)
		return
	}

	if c.Request.Method == "GET" {
		if userDetails, err := service.GetUserDetails(uint(id)); err != nil {
			c.HTML(http.StatusNotFound, "not_found.html", nil)
		} else {
			c.HTML(http.StatusFound, "update_users.html", userDetails)
		}
	} else if c.Request.Method == "PUT" {
		var userUpdate user.UserUpdateDTO
		c.Bind(&userUpdate)
		service.UpdateUser(uint(id), userUpdate)
		c.Header("HX-Redirect", fmt.Sprintf("/settings/user/%d", id))
	}
}
