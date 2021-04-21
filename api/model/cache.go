package model

import (
	"sync"
)


var serverPool = sync.Pool{
	New: func() interface{} {
		return new(Server)
	}}

type Server struct {
	Param *CommonLoginParam
	ReturnData ReturnDatas
	mu sync.Mutex
}

func NewServer() *Server {
	server,ok := serverPool.Get().(*Server)
	if ok {
		return server
	}
	return newServer()
}

func newServer() *Server {
	return &Server{}
}

// 把对象放回缓冲池，减少GC
func resetPool(server *Server)  {
	serverPool.Put(server)
}

func  (this *Server) DefulatGetServerList() (interface{} ,error) {

	var zone = NewZone()
	Zone, err := zone.getZone()
	if err !=nil {
		return nil,err
	}
	this.ReturnData.Zone = Zone

	servers, err := this.getList()
	if err !=nil {
		return nil,err
	}
	this.ReturnData.Server = servers
	this.recentlyServer()

	resetPool(this)

	return this.ReturnData,nil

}
