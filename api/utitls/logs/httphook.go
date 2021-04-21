package logs

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"lhc.go.game.sdk/utitls/nets"
	"net/url"
	"time"
)

type HttpHook struct {
}

func NewHttpHook() *HttpHook {
	return &HttpHook{}
}

func (l *HttpHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (l *HttpHook) Fire(entry *logrus.Entry) error {
	urls ,err := url.Parse(viper.GetString("log.urls"))
	if err != nil {
		return err
	}
	loguUrl := fmt.Sprintf("%s://%s%s",urls.Scheme,urls.Host,urls.Path)
	var data = make(map[string]interface{})
	data["data"] = entry.Data

	logData := make(map[string]interface{})
	formData := make(map[string]interface{})

	for k,v := range entry.Data  {
		formData[k] = v
	}
	logData["index"] = "http"
	logData["data"] = formData
	if val,ok := formData["postfrom"];ok {
		var form  = make(map[string]interface{})
		err := json.Unmarshal([]byte(val.(string)), &form)
		if err !=nil {
			return err
		}
		formData["postfrom"] = form
	}
	fmt.Printf("data:%#v\n",logData)

	ctx, cancelFunc := context.WithTimeout(context.Background(), 800*time.Millisecond)
	defer cancelFunc()

	go nets.Post(ctx,loguUrl, "application/json", logData)
	//bytes, err := nets.Post(ctx,loguUrl, "application/json", logData)
	//if err!=nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(string(bytes))
	select {
	case <-ctx.Done():
		fmt.Println("success......")
	}
	return nil
}