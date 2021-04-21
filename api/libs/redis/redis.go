package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

var Redis = &redis.Client{}

func InitRedis() error {
	fmt.Println("加载redis....")
	options := &redis.Options{Network:viper.GetString("redis.network"),Addr:viper.GetString("redis.addr")}
	Redis = redis.NewClient(options)

	return Redis.Ping().Err()
}