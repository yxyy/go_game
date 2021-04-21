package main

import (
	"fmt"
	"lhc.go.game.center/confs"
	"lhc.go.game.center/libs/mysql"
	"lhc.go.game.center/logs"
	"lhc.go.game.center/router"
)

func main()  {

	fmt.Println("加载日志组件.....")
	if err:=logs.InitLogs();err!=nil {
		logs.Fatal.Fatal(err)
		fmt.Println(err)
	}
	logs.Error.Info("log init")

	//加载配置
	fmt.Println("加载配置组件.....")
	if err := confs.InitConfs();err!=nil{
		logs.Fatal.Fatal(err)
		fmt.Println(err)
	}
	//加载mysql连接
	fmt.Println("加载mysql连接组件.....")
	if err := mysql.InitMysql();err != nil {
		logs.Fatal.Fatal(err)
		fmt.Println(err)
	}
	defer mysql.MysqlConnet.Close()

	//加载mysql连接
	//fmt.Println("加载es连接组件.....")
	//if err := es.InitEs();err != nil {
	//	logs.Fatal.Fatal(err)
	//	fmt.Println(err)
	//}

	//启动定时任务
	//go corn.InitCornRouter()

	//加载路由
	fmt.Println("加载路由组件.....")
	r := router.InitRouter()
	if err := r.Run(":8090");err!=nil{
		logs.Fatal.Fatal(err)
		fmt.Println(err)
	}
}
