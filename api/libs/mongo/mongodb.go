package mongo

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var Mongo *mongo.Client

func InitMongo() (err error) {
	// 设置客户端连接配置
	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s", viper.GetString("mongodb.host"), viper.GetString("mongodb.port")))

	// 连接到MongoDB
	Mongo, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
		return
	}
	// 检查连接
	err = Mongo.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("Connected to MongoDB!")
	return
}
