package main

import (
	"fmt"
	"log"
	"github.com/spf13/viper"
	"lhc.go.game.sdk/cmd"
	"lhc.go.game.sdk/conf"
	"lhc.go.game.sdk/router"
	"lhc.go.game.sdk/libs/redis"
	"lhc.go.game.sdk/libs/mysql"
	"lhc.go.game.sdk/utitls/logs"
	"net/http"
	_ "net/http/pprof"

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
	if err:= mysql.InitMysql();err !=nil {
		log.Fatal("mysql..初始化失败:%s....",err)
	}
	defer mysql.MysqlConnet.Close()

	//if err:=mongo.InitMongo();err!=nil{
	//	log.Fatal("mysql启动失败:%s.....",err)
	//}
	//defer mongo.Mongo.Disconnect(context.TODO())

	go func() {
		// 开启pprof，监听请求
		ip := "0.0.0.0:6060"
		if err := http.ListenAndServe(ip, nil); err != nil {
			fmt.Printf("start pprof failed on %s\n", ip)
		}
	}()
	r := router.InitRouter()
	if *cmd.Port=="" {
		*cmd.Port = viper.GetString("port")
	}
	if err := r.Run(":"+*cmd.Port);err!=nil{
		log.Fatal("服务器启动失败失败:%s.....",err)
	}
}
