package logs

import (
	"errors"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"time"
)
//致命错误
var Fatal = logrus.New()
//错误日志
var Error = logrus.New()
//http日志
var Http = logrus.New()
//http日志
var Corn = logrus.New()

func InitLogs() error {
	//设置输出json格式
	logrus.SetFormatter(&logrus.JSONFormatter{})
	// 设置日志级别为warn以上
	logrus.SetLevel(logrus.WarnLevel)

	Errorfile, err := LogFile("error.log")
	if err != nil {
		return err
	}
	Error.SetFormatter(&logrus.JSONFormatter{})
	Error.SetOutput(Errorfile)

	Httpfile, err := LogFile("http.log")
	if err != nil {
		return err
	}
	Http.SetOutput(Httpfile)

	Fatalfile, err := LogFile("fatal.log")
	if err != nil {
		return err
	}
	Fatal.SetOutput(Fatalfile)

	Cornfile, err := LogFile("corn.log")
	if err != nil {
		return err
	}
	Corn.SetOutput(Cornfile)

	return nil
}

func LogFile (name string) (*os.File,error) {

	if name=="" {
		return nil,errors.New("文件名不能为空")
	}

	//文件路径
	dir := "./logs/" + time.Now().Format("20060102")

	if !FileExist(dir) {
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
	}

	return os.OpenFile(dir + "/" + name, os.O_RDWR|os.O_APPEND|os.O_CREATE, os.ModePerm)

}


//判断文件是否存在
func FileExist(name string) bool {
	_, err := os.Stat(name)
	if err!=nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}


