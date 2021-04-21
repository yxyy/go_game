package test

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"lhc.go.game.sdk/libs/redis"
)

func Account(c *gin.Context)  {

	val := redis.Redis.HMGet("accountInfo:"+"123", "id","name","package").Val()

	fmt.Printf("%#v\n",val)

}
