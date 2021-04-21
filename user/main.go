package main

import (
	log "github.com/sirupsen/logrus"
	"lhc.go.game.user/libs/mysql"
	"lhc.go.game.user/conf"
	"lhc.go.game.user/router"
	"lhc.go.game.user/libs/redis"
	"lhc.go.game.user/utitls/logs"
)

func main()  {

	if err := conf.InitConfige();err!=nil{
		log.Fatal("加载配置文件失败:%s.....",err)
	}

	if err := redis.InitRedis();err!=nil{
		log.Fatal("redis启动失败:%s.....",err)
	}
	defer redis.Redis.Close()

	if err := logs.InitLogger();err!=nil{
		log.Fatal("logs..初始化失败:%s.....",err)
	}

	if err:=mysql.InitMysql();err!=nil{
		log.Fatal("mysql启动失败:%s.....",err)
	}
	defer mysql.MysqlConnet.Close()

	r := router.InitRouter()
	if err := r.Run(":8070");err!=nil{
		log.Fatal("服务器启动失败失败:%s.....",err)
	}
}
