package nets

import "github.com/gin-gonic/gin"

func GetIp(c *gin.Context) string {
	ip := c.Request.Header.Get("X-Forwarded-For")
	if ip !="" {
		return ip
	}

	ip = c.Request.Header.Get("X-Real-Ip")
	if ip !="" {
		return ip
	}

	return c.Request.RemoteAddr
}
