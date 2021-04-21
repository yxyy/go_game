package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"lhc.go.game.sdk/internal"
	"lhc.go.game.sdk/middleware"
	"lhc.go.game.sdk/test"
)

func InitRouter() *gin.Engine {
	fmt.Println("初始化路由")
	r := gin.Default()

	r.GET("/test/rand",test.Rand)
	r.GET("/test/zone",test.Zone)
	r.GET("/test/server",test.ResetServer)
	r.GET("/test/account",test.Account)

	// 请求日志收集
	r.Use(middleware.Log)
	// 跨域处理
	//r.Use(middleware.CrossDomain)

	r.POST("sdk/login/:package", internal.Login)
	r.POST("sdk/order", internal.Order)
	r.POST("sdk/callback/:package", internal.Callback)

	//r.Use(middleware.Auth)


	return r
}

