package model

import (
	"errors"
	"lhc.go.game.center/libs/mysql"
	"time"
)

type Menu struct {
	Id          int    `json:"id" form:"id"`
	Mark        string `json:"mark"  form:"mark"`
	Type        string `json:"type"  form:"type"`
	Name        string `json:"name"  form:"name"`
	Url         string `json:"url"  form:"url"`
	Parent      int    `json:"parent" gorm:"default:0"  form:"parent"`
	Icon        string `json:"icon"  form:"icon"`
	Sort      	int    `json:"sort"  form:"sort"`
	Status      int    `json:"status"  form:"status"`
	Author      string `json:"author"`
	CreateTime int64  `json:"create_time"  `
	UpdateTime int64  `json:"update_time"  `
}

type MenuTree struct {
	Id       int         `json:"id"`
	Mark     string      `json:"mark"`
	Type     string      `json:"type"`
	Name     string      `json:"name"`
	Url      string      `json:"url"`
	Parent   int         `json:"parent" gorm:"default:0"`
	Icon     string      `json:"icon"`
	Sort   int         `json:"sort"`
	Status   int         `json:"status"`
	Children []*MenuTree  `json:"children"`
}

func NewMenu() *Menu {
	return &Menu{}
}

func (this *Menu) Get() error {
	if err := mysql.MysqlConnet.Model(&this).Where("id = ?",this.Id).First(&this).Error;err!=nil{
		return err
	}

	return nil
}

func GetParentMenu(parent int) (menu []*Menu)  {
	mysql.MysqlConnet.Table("yxyy_menu").Where("parent = ?",parent).Order("sort").Order("create_time").Find(&menu)
	return
}

func GetMeunTree(parent int) (menuList []*MenuTree) {
	var menu []Menu
	mysql.MysqlConnet.Table("yxyy_menu").Where("parent = ?",parent).Order("sort").Order("create_time").Find(&menu)
	for _,v := range menu {
		Children := GetMeunTree(v.Id) // 拿到每个父菜单的子菜单
		node := &MenuTree{
			Id:     v.Id,
			Name:   v.Name,
			Url:    v.Url,
			Icon:   v.Icon,
			Mark:   v.Mark,
			Sort: v.Sort,
			Parent: v.Parent,
			Status: v.Status,
		}
		node.Children = Children
		menuList = append(menuList,node)
	}
	return
}

func (this *Menu) GetMeunTree(parent int) (menuList []*MenuTree) {
	var menu []Menu
	Db := mysql.MysqlConnet.Table("yxyy_menu").Where("parent = ?",parent)
	if this.Name != "" && this.Parent == 0 {
		Db = Db.Where("name like ","%"+this.Name+"%")
	}
	Db.Order("sort").Order("create_time").Find(&menu)
	for _,v := range menu {
		Children := this.GetMeunTree(v.Id) // 拿到每个父菜单的子菜单
		node := &MenuTree{
			Id:     v.Id,
			Name:   v.Name,
			Url:    v.Url,
			Icon:   v.Icon,
			Mark:   v.Mark,
			Sort: v.Sort,
			Parent: v.Parent,
			Status: v.Status,
			Children:Children,
		}
		menuList = append(menuList,node)
	}
	return
}

func (this *Menu) GetMen()  {
	Db := mysql.MysqlConnet.Table("yxyy_menu")
	if this.Name != "" {
		Db = Db.Where("name like ")
	}
}

func (m *Menu) UpdateData () error {
	if m.Name=="" {
		return errors.New("名称不能为空")
	}
	if m.Url=="" {
		return errors.New("url不能为空")
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