package model

import (
	"errors"
	"lhc.go.game.center/libs/mysql"
	"time"
)

type CrawlerMenu struct {
	Id int `json:"id" form:"id"`
	Name string `json:"name" form:"name"`
	Mark string `json:"mark" form:"mark"`
	CrawlerUrl string `json:"crawler_url" form:"crawler_url"`
	CronUrl string `json:"cron_url" form:"cron_url"`
	Status int `json:"status" form:"status"`
	CronTime   string `json:"cron_time" form:"cron_time"`
	CronFomatTime   string `json:"cron_fomat_time"`
	CreateTime int64  `json:"create_time"  `
	UpdateTime int64  `json:"update_time"  `
}

func NewCrawlerMenu() *CrawlerMenu {
	return &CrawlerMenu{}
}

func (this *CrawlerMenu) GetList()(data []*CrawlerMenu,total int,err error)  {
	Db :=mysql.MysqlConnet.Model(&this)
	if this.Status!=0 {
		Db = Db.Where("status = ?",this.Status)
	}
	if this.Name !="" {
		Db = Db.Where("title like  ?","%"+this.Name+"%")
	}
	if err:= Db.Find(&data).Error;err!=nil{
		return nil,0,err
	}
	Db.Count(&total)
	return
}

func (this *CrawlerMenu) GetOneCrawlerMenu() error {
	if this.Id ==0 {
		return errors.New("id 不能为空")
	}
	if err:= mysql.MysqlConnet.Model(&this).First(&this).Error;err!=nil{
			return err
	}
	return nil
}


func (m *CrawlerMenu) UpdateData () error {
	if m.Name=="" {
		return errors.New("名称不能为空")
	}
	m.UpdateTime=time.Now().Unix()
	if m.Id !=0 {
		if err:=mysql.MysqlConnet.Model(&m).Where("id = ?",m.Id).Update(&m).Error;err!=nil{
			return err
		}
	}else {
		m.CreateTime=time.Now().Unix()
		if err:=mysql.MysqlConnet.Model(&m).Create(&m).Error;err!=nil{
			return err
		}
	}
	return nil
}

func (this *CrawlerMenu) Update () error {
	if this.Id ==0 {
		return errors.New("id 不能为空")
	}
	this.UpdateTime=time.Now().Unix()
	if err:=mysql.MysqlConnet.Model(&this).Where("id = ?",this.Id).Update(&this).Error;err!=nil{
		return err
	}
	return nil
}
