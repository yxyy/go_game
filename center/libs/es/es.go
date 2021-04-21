package es

import (
	"fmt"
	"github.com/olivere/elastic/v7"
	"github.com/spf13/viper"
)

var Client *elastic.Client

func InitEs() (err error)  {

	host := viper.GetString("es.host")
	port := viper.GetString("es.port")
	urls := fmt.Sprintf("%s:%s",host,port)

	Client, err = elastic.NewClient(elastic.SetURL(urls))
	if err!=nil {
		return err
	}
	return nil
}