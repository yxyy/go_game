package router

import (
	"github.com/gin-gonic/gin"
	"lhc.go.game.log/controller/producter"
	"lhc.go.game.log/middleware"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	// 请求日志收集
	//r.Use(middleware.Log)
	// 跨域处理
	r.Use(middleware.CrossDomain)

	//r.Use(middleware.Auth)

	r.POST("logs/push",producter.Push)

	return r
}

