package corn

import (
	"github.com/robfig/cron"
	"lhc.go.game.center/controller/corn/netbian"
	"lhc.go.game.center/logs"
)

func InitCornRouter()  {

	c := cron.New()

	if err := c.AddFunc("* * * 1 1 1", netbian.NewNetbian().Run);err!=nil{
		logs.Error.Info(err)
	}
}
