package middleware

import (
	"docker-site/helper/htmx"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {
	session := sessions.Default(c)
	sessionToken := session.Get("token")
	userName := session.Get("username")

	if sessionToken == nil || userName == nil {
		c.Header(htmx.Redirect, "/")
		c.Status(http.StatusPermanentRedirect)
		c.Abort()
	}

	c.Next()
}
