package model

import (
	"errors"
	"lhc.go.game.center/libs/mysql"
	"time"
)

type UserGroup struct {
	Id int `json:"id"`
	Name string `json:"name" form:"name"`
	Mark string `json:"mark" form:"mark"`
	Description string `json:"description" form:"description"`
	CreateTime int64 `json:"create_time"`
	UpdateTime int64 `json:"update_time"`
}

func (ug *UserGroup) GetUserGourpList(sort Params) (data []*UserGroup,total int,err error) {

	Db:=mysql.MysqlConnet.Model(&ug)
	if ug.Mark!="" {
		Db = Db.Where("mark = ?",ug.Mark)
	}
	if ug.Name!="" {
		Db = Db.Where("name = ? ",ug.Name)
	}
	if err = Db.Offset(sort.Satrt).Limit(sort.Length).Find(&data).Error;err!=nil{
		return
	}
	Db.Count(&total)
	return
}

func (ug *UserGroup) UpdateData () error {
	if ug.Name=="" {
		return errors.New("名称不能为空")
	}
	if ug.Mark=="" {
		return errors.New("标识不能为空")
	}
	ug.UpdateTime=time.Now().Unix()
	if ug.Id !=0 {
		if err:=mysql.MysqlConnet.Model(&ug).Where("id = ?",ug.Id).Update(&ug).Error;err!=nil{
			return err
		}
	}else {
		ug.CreateTime=time.Now().Unix()
		if err:=mysql.MysqlConnet.Model(&ug).Create(&ug).Error;err!=nil{
			return err
		}
	}
	return nil
}


func GetUserGourpList() (data []*UserGroup) {
	mysql.MysqlConnet.Table("yxyy_user_group").Find(&data)
	return
}