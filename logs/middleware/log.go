package middleware

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"time"
)

type Logger struct {
	Method string `json:"method"`
	URL string `json:"url"`
	PostFrom map[string]string `json:"post_from"`
	Tick string `json:"tick"`

}

func NewLogger() *Logger {
	return &Logger{
		PostFrom: make(map[string]string),
	}
}

func Log(c *gin.Context)  {

	logger := NewLogger()

	logger.Method = c.Request.Method
	query := ""
	if c.Request.URL.RawQuery !="" {
		query = "?" + c.Request.URL.RawQuery
	}
	logger.URL = c.Request.URL.Path + query

	if c.Request.Method == "POST" {
		switch c.Request.Header["Content-Type"][0] {
		case "application/x-www-form-urlencoded":
			if err:=c.Request.ParseForm();err!=nil {
				log.Println(err)
			}
			for k,v :=range c.Request.PostForm {
				logger.PostFrom[k] = v[0]
			}
		case "application/json":
			body, err := ioutil.ReadAll(c.Request.Body)
			if err != nil {
				log.Println(err)
			}
			if err := json.Unmarshal(body,&logger.PostFrom);err != nil {
				log.Println(err)
			}
			fmt.Println(string(body))
		}
	}

	logger.Tick = time.Now().Format("2006-01-02 15:04:05")

	fmt.Printf("%#v\n",logger)

}
