package model

import (
	"encoding/json"
	"lhc.go.game.sdk/libs/redis"
)

type Zone struct {
	Id string
	Name string
}

func NewZone() *Zone {
	return &Zone{}
}

func (this *Zone) getZone() (zone map[string]interface{},err error) {

	val := redis.Redis.Get("zone").Val()
	if val == "" {
		return  this.getZoneByMysql()
	}

	if err = json.Unmarshal([]byte(val), &zone);err!=nil{
		return nil,err
	}

	return zone,nil
}

func (this *Zone)getZoneByMysql() (zone map[string]interface{},err error)  {

	return nil,nil
}