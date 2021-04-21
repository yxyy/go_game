package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	"net/http"
)

func AdminSessionAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		if session.Get("user_id")==nil {
			c.Abort()
			c.Redirect(http.StatusFound,"/admin/login")
			return
		}
	}
}
