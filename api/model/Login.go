package model

import "sync"

type Loginer interface {
	Vaild() error
	GetServer() (interface{} ,error)
}

type CommonLoginParam struct {
	Package int `json:"package" form:"package"`
	Channel string `json:"channel" form:"channel"`
	UserId string `json:"userId" form:"userId"`
	Account string `json:"account"`
}

var ServerCommonLoginParamPool = sync.Pool{
	New: func() interface{} {
		return new(CommonLoginParam)
	}}

// 把对象放回缓冲池，减少GC
func reseCommonLoginParamtPool(CommonLoginParam *CommonLoginParam)  {
	ServerCommonLoginParamPool.Put(CommonLoginParam)
}

func Login(loginer Loginer) (interface{} ,error) {
	if err:=loginer.Vaild();err !=nil {
		return nil,err
	}
	return loginer.GetServer()
}



