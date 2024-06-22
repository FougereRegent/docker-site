package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthMiddleware(c *gin.Context) {
	session := sessions.Default(c)
	sessionToken := session.Get("token")
	userName := session.Get("username")

	if sessionToken == nil || userName == nil {
		c.HTML(http.StatusUnauthorized, "", nil)
		c.Abort()
	}

	c.Next()
}
