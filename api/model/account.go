package model

import "lhc.go.game.sdk/libs/redis"

type Account struct {
	RoleId string
	RoleName string
	Level int
	Career int
	CommonLoginParam
}

func (this *Account) getInfo()  {
	if this.Account=="" {
		return
	}
	val := redis.Redis.HMGet("accountInfo:" + this.Account,"role_id","role_name","server","level","career","update_time").Val()
	if val==nil {
		return
	}

}
