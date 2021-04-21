package logs

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"lhc.go.game.sdk/utitls/nets"
	"net/url"
	"time"
	"context"
)

var RequestId string

type LogHook struct {
	RequestId string
}

type MyFormatter struct {

}

func NewLogHook() *LogHook {
	return &LogHook{}
}

func NewLoggerHook(request_id string) *LogHook {
	return &LogHook{RequestId:request_id}
}

func (l *LogHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (l *LogHook) Fire(entry *logrus.Entry) error {
	urls ,err := url.Parse(viper.GetString("log.urls"))
	if err != nil {
		return err
	}
	loguUrl := fmt.Sprintf("%s://%s%s",urls.Scheme,urls.Host,urls.Path)
	var data = make(map[string]interface{})
	switch entry.Level {
	case logrus.DebugLevel:
		data["index"] = "debug"
	case logrus.InfoLevel:
		data["index"] = "info"
	case logrus.ErrorLevel:
		data["index"] = "error"
	default:
		data["index"] = "other"
	}
	data["data"]=entry.Data

	ctx, cancelFunc := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelFunc()
	bytes, err := nets.Post(ctx,loguUrl, "application/json", data)
	if err!=nil {
		logrus.Println(err)
	}
	fmt.Println(string(bytes))
	return nil
}
