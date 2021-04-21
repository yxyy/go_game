package mysql

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

type Mysql struct {
	Host string
	Port int
	UserName string
	PassWord string
	DbName string
}

var MysqlConnet *gorm.DB


func InitMysql() (err error) {

	fmt.Println("加载mysql。。。。")

	MysqlConnet, err = gorm.Open("mysql", GetDSN())
	if err!=nil {
		return err
	}

	MysqlConnet.SingularTable(true)//gorm默认把表名都加上 "s"，关掉默认

	gorm.DefaultTableNameHandler= func(db *gorm.DB, defaultTableName string) string {
		return "yxyy_" + defaultTableName
	}

	return nil
}

func GetDSN() string {

	var mysql Mysql

	mysql.Host=viper.GetString("mysql.host")
	mysql.Port=viper.GetInt("mysql.port")
	mysql.UserName=viper.GetString("mysql.username")
	mysql.PassWord=viper.GetString("mysql.password")
	mysql.DbName=viper.GetString("mysql.dbname")

	return  fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",mysql.UserName,mysql.PassWord,mysql.Host,mysql.Port,mysql.DbName)
}