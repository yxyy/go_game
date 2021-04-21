package model

import (
	"errors"
	"lhc.go.game.sdk/libs/mysql"
	"lhc.go.game.sdk/utitls/random"
)

type Order interface {
	Valid() error
	Add() error
}

func AddOrder(o Order) error  {
	if err := o.Valid();err!=nil{
		return err
	}
	if err := o.Add();err!=nil{
		return err
	}
	return nil
}

type Product struct {
	ProductId int `json:"product_id" form:"product_id" gorm:"product_id"`
	ProductName string `json:"product_name" form:"product_name" gorm:"product_name"`
	Money float64 `json:"money" form:"money" gorm:"money"`
	MoneyType int `json:"money_type" form:"money_type" gorm:"money_type"`
	Gold int `json:"gold" form:"gold" gorm:"gold"`
}

type Machine struct {
	Ip string `json:"ip"  form:"ip" gorm:"ip"`
	Imei string `json:"imei"  form:"imei" gorm:"imei"`
	Idfa string `json:"idfa"  form:"idfa" gorm:"idfa"`
	Mac	string `json:"mac"  form:"mac" gorm:"mac"`
}

type OrderTmp struct {
	OrderId string `json:"order_id" form:"order_id" gorm:"order_id"`
	Role
	Product
	Machine
	Time
	CommonParam
}

func NewOrderTmp(request_id string) *OrderTmp {
	return &OrderTmp{OrderId:request_id}
}

func (p *Product) Valid () error {
	if p.ProductId <= 0 {
		return errors.New("product_id is invalid")
	}
	if p.ProductName =="" {
		return errors.New("product_name is invalid")
	}
	if p.Money <= 0 {
		return errors.New("money is invalid")
	}
	if p.MoneyType <= 0 {
		return errors.New("money_type is invalid")
	}
	if p.Gold <= 0 {
		return errors.New("gold is invalid")
	}

	return nil
}

func (t *OrderTmp) Valid() error {
	if err:=t.CommonParam.Valid();err!=nil{
		return err
	}
	if err:=t.Role.Valid();err!=nil{
		return err
	}
	if err:=t.Product.Valid();err!=nil{
		return err
	}
	if t.OrderId == "" {
		t.OrderId = random.GetRequstId()
	}
	return nil
}

func (t *OrderTmp) Add() error  {
	if err := mysql.MysqlConnet.Model(&t).Create(&t).Error;err!=nil{
		return  err
	}
	return nil
}


