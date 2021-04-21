package model

import (
	"encoding/json"
	"errors"
	"github.com/Shopify/sarama"
	"lhc.go.game.log/libs/kafka"
)

type Message struct {
	RequestId string `json:"request_id" form:"request_id"`
	Index string `json:"index" form:"index"`
	Data interface{} `json:"data" form:"data"`
}

func NewMessage() *Message {
	return &Message{}
}

func (this *Message) Valid() error {
	if this.Data =="" {
		return errors.New("数据不能为空")
	}
	if this.Index == "" {
		this.Index = "info"
	}
	return nil
}

func (this *Message) Proucter() error {
	if this.Valid() !=nil {
		return errors.New("参数错误")
	}

	data,err:= json.Marshal(this.Data)
	if err !=nil {
		return err
	}

	message := sarama.ProducerMessage{}
	//var bytes = sarama.ByteEncoder{}
	message.Topic = this.Index
	//message.Value = sarama.ByteEncoder(data)
	message.Value = sarama.StringEncoder(data)

	kafka.Producer.Input() <- &message

	return nil
}


