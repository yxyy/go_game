package test

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"lhc.go.game.sdk/libs/redis"
	"lhc.go.game.sdk/utitls/random"
	"time"
)

func Zone(c *gin.Context)  {

	zone := `{"25":"1-100区"}`
	redis.Redis.Set("zone",zone,-1)
}

func Rand( c *gin.Context)  {

	var arr []string
	var mmp = make(map[string]bool)
	for i:=0;i<1000 ;i++  {
		go func() {
			arr = append(arr, random.GetRequstIdByLock())
		}()
	}
	time.Sleep(3*time.Second)
	fmt.Println(len(arr))

	for _,v := range arr  {
		mmp[v] = true
	}
	fmt.Println(len(mmp))
}
