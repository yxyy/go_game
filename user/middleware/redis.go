package middleware

import (
	"github.com/gin-gonic/gin"
	"lhc.go.game.user/libs/redis"
	"net/http"
	"time"
)

//find the redis , if exist retun
func Login(c *gin.Context)  {

	account := c.PostForm("account")

	if(redis.Redis.Get(account).Val()!=""){
		c.JSON(http.StatusOK,gin.H{"code":200,"msg":"登陆成功"})
		c.Abort()
	}else {
		redis.Redis.Set(account,"",10*time.Minute)
	}
	c.Next()
}