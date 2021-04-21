package middleware

import "github.com/gin-gonic/gin"

// 跨域
func CrossDomain(c *gin.Context)  {
	//实际生产环境 * 换成自己的域名
	c.Header("Access-Control-Allow-Origin", "*")
	c.Next()
}
