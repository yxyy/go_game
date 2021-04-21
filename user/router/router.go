package router

import (
	"github.com/gin-gonic/gin"
	"lhc.go.game.user/controller/login"
	"lhc.go.game.user/controller/token"
	"lhc.go.game.user/controller/user"
	"lhc.go.game.user/controller/wechat"
	"lhc.go.game.user/middleware"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	// 请求日志收集
	r.Use(middleware.Log)
	// 跨域处理
	r.Use(middleware.CrossDomain)

	r.POST("user/login",login.Login)

	//r.Use(middleware.Auth)


	//刷新 access_token
	r.POST("user/refreshToken",token.RefreshToken)
	r.GET("user/get",user.Get)
	r.GET("user",user.GetList)
	r.POST("user",user.GetList)
	r.POST("user/updateData",user.UpdateData)


	r.GET("/wechat",wechat.Token)

	return r
}

