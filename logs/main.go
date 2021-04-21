package main

import (
	"lhc.go.game.log/conf"
	"lhc.go.game.log/libs/kafka"
	"lhc.go.game.log/router"
	"log"
)

func main()  {

	if err := conf.InitConfige();err!=nil{
		log.Fatal("加载配置文件失败:%s.....",err)
	}

	if err := kafka.InitKafka();err!=nil{
		log.Fatal("加载kafka失败:%s.....",err)
	}
	defer kafka.Producer.Close()
	defer kafka.Consumer.Close()

	r := router.InitRouter()
	if err := r.Run(":8060");err!=nil{
		log.Fatal("服务器启动失败失败:%s.....",err)
	}
}
