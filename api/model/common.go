package model

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"runtime"
)

const (
	StatusOK = 200
	StatusParamError = 40001
	StatusServerError = 50000
	StatusServerParseError = 50001
	MsgOK = "OK"
	MsgFail = "fail"
	MsgIllegal = "illegal"
)

type Valider interface {
	Valid() error
}

type ReturnData struct {
	Code int `json:"code,omitempty"`
	Msgs  error `json:"msgs,omitempty"`
	Msg  string `json:"msg,omitempty"`
	Data interface{} `json:"data,omitempty"`
	Total int `json:"total,omitempty"`
	RequestId string `json:"request_id,omitempty"`
}

type CommonParam struct {
	Channel string `json:"channel"  form:"channel" gorm:"channel"`
	Server string `json:"server"  form:"server" gorm:"server"`
	Package int `json:"package"  form:"package" gorm:"package"`
}

type Role struct {
	Account string `json:"account" form:"account" gorm:"account"`
	RoleId	int `json:"role_id" form:"role_id" gorm:"role_id"`
	RoleName string `json:"role_name" form:"role_name" gorm:"role_name"`
	RoleLevel int `json:"role_level" form:"role_level" gorm:"role_level"`
}

type Time struct {
	CreateTime int64 `json:"create_time" form:"create_time" gorm:"create_time"`
	UpdateTime int64 `json:"update_time"  form:"update_time" gorm:"update_time"`
}

func NewReturnData(request_id string) *ReturnData {
	return &ReturnData{RequestId:request_id,Code:StatusOK,Msg:MsgOK}
}

func (p *CommonParam) Valid () error {
	if p.Channel =="" {
		return errors.New("channel is invalid")
	}
	if p.Server =="" {
		return errors.New("server is invalid")
	}
	if p.Package ==0 {
		return errors.New("package is invalid")
	}
	return nil
}

func (r *Role) Valid () error {
	if r.Account =="" {
		return errors.New("account is invalid")
	}
	if r.RoleName =="" {
		return errors.New("role_name is invalid")
	}
	if r.RoleId ==0 {
		return errors.New("role_id is invalid")
	}
	if r.RoleLevel ==0 {
		return errors.New("role_level is invalid")
	}
	return nil
}

//错误处理
func CatchErr(Return ReturnData)  {
	pc, _, _, _ := runtime.Caller(1)
	if Return.Msgs !=nil && Return.Code >200 {
		logrus.WithFields(map[string]interface{}{
			"request_id" : Return.RequestId,
			"code" : Return.Code,
			"msg":Return.Msgs.Error(),
			"file" :fmt.Sprint(runtime.FuncForPC(pc).FileLine(pc)),
		}).Error("执行错误")
	}
}

