package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/registry/consul"
)

func main() {
	// 添加consul地址
	cr := consul.NewRegistry(
		registry.Addrs("120.78.167.190:8500"))

	// 使用gin作为router
	router := gin.Default()
	router.GET("/user", func(c *gin.Context) {
		c.String(http.StatusOK, "user api")
	})

	// 初始化go micro
	server := web.NewService(
		web.Name("productService"),                          // 当前微服务服务名
		web.Registry(cr),                                    // 注册到consul
		web.Address(":8081"),                                // 端口
		web.Metadata(map[string]string{"protocol": "http"}), // 元信息
		web.Handler(router)) // 路由

	_ = server.Run()
}
