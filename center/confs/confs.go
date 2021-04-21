package confs

import (
	"fmt"
	"github.com/spf13/viper"
)

func InitConfs() error {

	fmt.Println("加载配置文件。。。。")

	viper.SetConfigName("confs")

	viper.SetConfigType("yaml")

	viper.AddConfigPath("./confs")

	if err := viper.ReadInConfig();err!=nil{
		return err
	}
	return nil
}