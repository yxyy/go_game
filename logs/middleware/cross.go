package middleware

import "github.com/gin-gonic/gin"

// 跨域
func CrossDomain(c *gin.Context)  {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Next()
}
