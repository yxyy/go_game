package test

import (
	"github.com/gin-gonic/gin"
	"lhc.go.game.sdk/libs/redis"
)

func ResetServer( c *gin.Context)  {
	server := `[
                {
                        "default" : 0,
                        "is_show" : 1,
                        "group_id" : 66,
                        "is_create" : 1,
                        "open_time" : "2020-03-26 10:00:00",
                        "combine_time" : "0",
                        "server_id" : "ly001",
                        "name" : "星玩#技术测试1",
                        "flag" : "ly001",
                        "sort" : 1,
                        "status" : 3,
                        "zone" : 29,
                        "md5_key" : "ljtd_phone_s001a_123!@qwe",
                        "domain" : "ljtd10004ase.shxingwan.com",
                        "api_url" : "http://ljtd10004ase.shxingwan.com:4001/",
                        "max_online" : 1000,
                        "server_addr" : "ljtd10004ase.shxingwan.com",
                        "server_port" : 6010,
                        "chat_addr" : "127.0.0.1",
                        "chat_port" : 6501,
                        "channel_flag" : "xingwan",
                        "package_flags" : "[\"100001\",\"100002\"]",
                        "created_time" : "0000-00-00 00:00:00"
                },
                {
                        "default" : 1,
                        "is_show" : 1,
                        "group_id" : 67,
                        "is_create" : 1,
                        "open_time" : "2019-04-08 18:09:57",
                        "combine_time" : "0",
                        "server_id" : "lj002",
                        "name" : "正式测试服",
                        "flag" : "lj002",
                        "sort" : 2,
                        "status" : 3,
                        "zone" : 29,
                        "md5_key" : "ljtd_test_s002a_123!@qwe",
                        "domain" : "47.102.147.91",
                        "api_url" : "http://47.102.147.91:4002/",
                        "max_online" : 1000,
                        "server_addr" : "47.102.147.91",
                        "server_port" : 6020,
                        "chat_addr" : "127.0.0.1",
                        "chat_port" : 6501,
                        "channel_flag" : "xingwan",
                        "package_flags" : "[\"100002\"]",
                        "created_time" : "0000-00-00 00:00:00"
                }
			]`

	val := redis.Redis.Set("xingwan", server, -1)
	c.JSON(200,gin.H{"data":val})
}
