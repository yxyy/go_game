package middleware

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"lhc.go.game.user/utitls/logs"
	"lhc.go.game.user/utitls/random"
)

func Log(c *gin.Context)  {

	PostFrom := make(map[string]interface{})
	HttpRequest := make(map[string]interface{})

	HttpRequest["method"] = c.Request.Method

	query := ""
	if c.Request.URL.RawQuery !="" {
		query = "?" + c.Request.URL.RawQuery
	}
	url := c.Request.URL.Path + query

	if c.Request.Method == "POST" {
		switch c.Request.Header["Content-Type"][0] {
		case "application/x-www-form-urlencoded":
			if err:=c.Request.ParseForm();err!=nil {
				log.Println(err)
			}
			for k,v :=range c.Request.PostForm {
				PostFrom[k] = v[0]
			}
		case "application/json":
			if err := json.NewDecoder(c.Request.Body).Decode(&PostFrom);err!=nil{
				log.Error(err)
			}
		}

		bytes2, err := json.Marshal(PostFrom)
		if err !=nil {
			log.Error(err)
		}
		HttpRequest["postfrom"] = string(bytes2)

	}

	c.Set("request_id",random.GetRequstId())
	logs.RequestId = c.GetString("request_id")
	HttpRequest["request_id"] = c.GetString("request_id")
	HttpRequest["url_path"] = url
	logs.HttpLogger.WithFields(HttpRequest).Info("请求日志")
}
