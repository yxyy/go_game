package conf

import (
	"fmt"
	"github.com/spf13/viper"
	"lhc.go.game.sdk/cmd"
)

var Conf Config

type Config struct {
	LogIsOpen bool
}

func InitConfige() error {

	if *cmd.Config !="" {

	}
	viper.SetConfigType("yaml")
	viper.SetConfigName("confs")
	viper.AddConfigPath("./conf")
	if err := viper.ReadInConfig();err!=nil{
		fmt.Printf("Fatal error config file: %s \n", err)
		return err
	}
	Conf.LogIsOpen = viper.GetBool("log.is_open")
	return nil
}
