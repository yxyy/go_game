package model

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"runtime"
)

type Page struct {
	Page int `form:"page"`
	Length int `form:"length"`
}

func NewPage() *Page  {
	return &Page{}
}

type ReturnData struct {
	Code int `json:"code,omitempty"`
	Msg  error `json:"msg,omitempty"`
	Data interface{} `json:"data,omitempty"`
	Total int `json:"total,omitempty"`
	RequestId string `json:"request_id,omitempty"`
}

type HttpLogData struct {
	Index string `json:"index"`
	Data map[string]interface{} `json:"data"`
}

func NewHttpLogData() *HttpLogData  {
	return &HttpLogData{Data:make(map[string]interface{})}
}

func NewReturnData() *ReturnData {
	return &ReturnData{}
}

func NewReturnDataByRequestId(request_id string) *ReturnData {
	return &ReturnData{RequestId:request_id}
}

//错误处理
func CatchErr(Return ReturnData)  {
	pc, _, _, _ := runtime.Caller(1)
	if Return.Msg !=nil && Return.Code >200 {
		logrus.WithFields(map[string]interface{}{
			"request_id" : Return.RequestId,
			"code" : Return.Code,
			"msg":Return.Msg.Error(),
			"file" :fmt.Sprint(runtime.FuncForPC(pc).FileLine(pc)),
		}).Error("执行错误")
	}
}

