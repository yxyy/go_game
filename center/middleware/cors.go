package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		// 接受指定域的请求，可以使用*不加以限制，但不安全
		c.Header("Access-Control-Allow-Origin", "*")
		//请求的允许的头字段与权限控制
		//c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		////请求类型
		//c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		//c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		////是否允许后续请求携带认证信息,该值只能是true,否则不返回
		//c.Header("Access-Control-Allow-Credentials", "true")
		// 放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}
