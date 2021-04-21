package logs

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"lhc.go.game.sdk/utitls/file"
	"os"
	"time"
)

var HttpLogger = logrus.New()

func InitLogger() error {
	hook := NewLogHook()
	//全局日志
	textformatter := logrus.TextFormatter{
		DisableColors:             true,
		TimestampFormat: "2006-01-02 15:04:05",
	}
	logrus.SetFormatter(&textformatter)
	logrus.AddHook(hook)

	//记录参数日志
	httpHook := NewHttpHook()
	HttpLogger.SetFormatter(&MyTextFormatter{})
	HttpLogger.AddHook(httpHook)

	return logsOut()
}

func logsOut() error {
	filepath := viper.GetString("log.log_dir")
	if !file.FileExist(filepath) {
		err := os.MkdirAll(filepath, os.ModePerm)
		if err !=nil {
			return err
		}
	}

	filename := filepath+"/"+time.Now().Format("200601")
	openFile, err := fileOp(filename)
	if err!=nil {
		return err
	}
	logrus.SetOutput(openFile)

	//httpfilename := filepath+"/http_"+time.Now().Format("200601")
	//httpFile, err := fileOp(httpfilename)
	//if err!=nil {
	//	return err
	//}
	HttpLogger.SetOutput(openFile)

	return nil
}

func fileOp(filename string) (*os.File,error)  {
	return os.OpenFile(filename, os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModePerm)
}

