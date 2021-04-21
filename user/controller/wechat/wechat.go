package wechat

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type WeChat struct {
	Signature string `from:"signature"`
	Timestamp string `from:"timestamp"`
	Nonce string `from:"nonce"`
	Echostr string  `from:"echostr"`
}

func NewWeChat() *WeChat  {
	return &WeChat{}
}

func Token(c *gin.Context)  {

	c.String(http.StatusOK,"%s",false)
	return
}
