package pk5001

import "lhc.go.game.sdk/model"

type XingwanCallback struct {
	OrderId string `json:"orderId" form:"orderId" gorm:"orderId"`
	UserId string `json:"user_id"`
	AppId string
	ServerId string
	Amount float64
	Time int64
	Extension string
	GoodsName string
	model.CommonParam
	model.Role
	model.Machine
}

func NewXingwanCallback() *XingwanCallback {
	return &XingwanCallback{}
}

func (x *XingwanCallback) Valid () error  {
	return nil
}

func (x *XingwanCallback) Add () error {

	return nil
}
