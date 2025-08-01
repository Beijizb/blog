package middleware

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		if session.Get("user_id") == nil {
			c.Redirect(http.StatusFound, "/admin/login")
			c.Abort()
			return
		}
		c.Next()
	}
}
