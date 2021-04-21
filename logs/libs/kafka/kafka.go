package kafka

import (
	"errors"
	"github.com/Shopify/sarama"
	"github.com/spf13/viper"
)

var (
	Producer sarama.AsyncProducer
	Consumer sarama.Consumer
)

func InitKafka() (err error) {

	addrs := viper.GetStringSlice("kafka.addr")
	if addrs == nil {
		return errors.New("地址配置错误")
	}

	//定义一个kafka的生产者配置变量
	config := sarama.NewConfig()
	// 等待服务器所有副本都保存成功后的响应
	config.Producer.RequiredAcks =sarama.WaitForAll
	// 随机的分区类型：返回一个分区器，该分区器每次选择一个随机分区
	config.Producer.Partitioner=sarama.NewRandomPartitioner
	// 是否等待成功和失败后的响应
	config.Producer.Return.Successes=true

	Producer, err = sarama.NewAsyncProducer(addrs, config)
	if err!=nil {
		return err
	}

	Consumer, err = sarama.NewConsumer(addrs, nil)
	if err!=nil {
		return err
	}

	return nil
}
