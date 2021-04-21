package redis

import (
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

var Redis = &redis.Client{}

func InitRedis() error {

	options := &redis.Options{Network:viper.GetString("network"),Addr:viper.GetString("addr")}
	Redis = redis.NewClient(options)

	return Redis.Ping().Err()
}