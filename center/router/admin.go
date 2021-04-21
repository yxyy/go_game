package router

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"lhc.go.game.center/controller/admin"
	"lhc.go.game.center/controller/admin/category"
	"lhc.go.game.center/controller/admin/crawler"
	"lhc.go.game.center/controller/admin/gallery"
	"lhc.go.game.center/controller/admin/login"
	"lhc.go.game.center/controller/admin/menu"
	"lhc.go.game.center/controller/admin/user_group"
	"lhc.go.game.center/middleware"
)

func InitAdminRouter(r *gin.Engine)  {

	//设置存储引擎，
	store := cookie.NewStore([]byte(viper.GetString("session_secret")))
	// 路由使用session中间件，参数mysession，指的是session的名字，也是cookie的名字
	// store是前面创建的存储引擎，我们可以替换成其他存储引擎
	r.Use(sessions.Sessions("yxyy.session", store))

	r.GET("/admin/login", login.Index)
	r.POST("admin/login",login.Login)

	//AdminGroup := r.Group("/admin").Use(middleware.AdminSessionAuth())
	AdminGroup := r.Group("/admin").Use(middleware.Cors(),middleware.Auth)
	{
		AdminGroup.GET("logout",login.Logout)
		AdminGroup.GET("index",admin.Index)
		AdminGroup.GET("welcome",admin.Welcome)

		AdminGroup.GET("menu/index",menu.Index)
		AdminGroup.GET("menu/id/:id",menu.GetOneMenu)
		AdminGroup.GET("menu",menu.GetOneMenu)
		AdminGroup.POST("menu",menu.GetOneMenu)
		AdminGroup.GET("getParentMenu/:parent",menu.GetMenuParent)
		AdminGroup.POST("menu/UpdataMenu",menu.UpdateData)


		//AdminGroup.POST("user/password",user.Update)

		AdminGroup.GET("user_group/index",user_group.Index)
		AdminGroup.POST("user_group/index",user_group.GetUserGourpList)
		AdminGroup.GET("user_group/add",user_group.Add)
		AdminGroup.POST("user_group/add",user_group.UpdateData)

		AdminGroup.GET("category/index",category.Index)
		AdminGroup.POST("category/index",category.GetList)
		AdminGroup.POST("category/update",category.GetOneCategory)
		AdminGroup.POST("category/updateData",category.UpdateData)

		AdminGroup.GET("netbian/index", crawler.Index)
		AdminGroup.POST("netbian/index", crawler.GetList)
		AdminGroup.POST("netbian/update", crawler.GetOneCategory)
		AdminGroup.POST("netbian/updateData", crawler.UpdateData)
		AdminGroup.POST("netbian/status", crawler.Update)
		AdminGroup.POST("netbian/crawler", crawler.Crawler)


		AdminGroup.GET("gallery/index", gallery.Index)
		AdminGroup.POST("gallery/index", gallery.GetList)
		AdminGroup.POST("gallery/update", gallery.Update)

	}


}