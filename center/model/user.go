package model

import (
	"crypto/md5"
	"errors"
	"fmt"
	"lhc.go.game.center/controller/common"
	"lhc.go.game.center/libs/mysql"
	"strconv"
	"time"
)

type User struct {
	Id int 	`json:"id" form:"id"`
	Account string `json:"account" form:"account"`
	Nickname string `gorm:"nickname" json:"nickname" form:"nickname"`
	Img  string `goem:"img" json:"img" form:"img"`
	Password string `gorm:"password" json:"-"`
	RePassword string `gorm:"-" json:"-" form:"password"`
	Salt string `json:"-"`
	Status int `json:"status" form:"status"`
	GroupId int ` json:"group_id" form:"group_id"`
	Phone int	`json:"phone" form:"phone"`
	Wechat string	`gorm:"wechat" json:"wechat" form:"wechat"`
	ExpirationTime string `gorm:"-" json:"-" form:"expiration"`
	Expiration int64 `json:"expiration"`
	LoginNum int	`json:"login_num"`
	LastLoginIp string `json:"last_login_ip"`
	LastLoginTime int64 `json:"last_login_time"`
	CreateTime int64 `json:"create_time"`
	UpdateTime int64 `json:"update_time"`
	Search string `gorm:"-" form:"search"`
	//UserGroup
}

func NewUser() *User {
	return &User{}
}

func (u *User) GetOne() error {
	if u.Id==0 {
		return	errors.New("id 不能为空")
	}
	if err:=mysql.MysqlConnet.Model(&u).Where("id = ?",u.Id).First(&u).Error;err!=nil{
		return err
	}
	return nil
}

func (u *User) GetList(page *Page)(data []*User,total int,err error)  {
	mysql.MysqlConnet.LogMode(true)
	Db :=mysql.MysqlConnet.Model(&u)
	if u.Account!="" {
		Db = Db.Where("account like ?","%"+u.Account+"%")
	}
	if u.Nickname!="" {
		Db = Db.Where("nickname like ?","%"+u.Nickname+"%")
	}
	if u.Phone !=0 {
		Db = Db.Where("phone like ?",u.Phone)
	}
	if u.Wechat !="" {
		Db = Db.Where("wechat like ?","%"+u.Wechat+"%")
	}
	if u.Search !="" {
		Db = Db.Where("account like ?","%"+u.Search+"%").
				Or("nickname like ?","%"+u.Search+"%").
				Or("phone like ?","%"+u.Search+"%").
				Or("wechat like ?","%"+u.Search+"%")
	}
	//Db = Db.Joins("LEFT JOIN yxyy_user_group as g on u.group_id=g.id")
	if err:= Db.Offset((page.Page-1)*page.Length).Limit(page.Length).Find(&data).Error;err!=nil{
		return nil,0,err
	}
	Db.Count(&total)
	return
}

func (u *User) Update() error {
	if u.Status<0 {
		return errors.New("参数错误")
	}
	if u.ExpirationTime!="" {
		tick,err := time.Parse("2006-01-02",u.ExpirationTime)
		if err != nil {
			return err
		}
		u.Expiration=tick.Unix()
	}
	if u.RePassword!="" {
		var user = &User{Id:u.Id}

		if err := user.GetOne();err!=nil{
			return err
		}
		if user.Salt=="" {
			return errors.New("账号异常")
		}
		u.Password=fmt.Sprintf("%x",md5.Sum([]byte(u.RePassword+user.Salt)))
	}
	u.UpdateTime=time.Now().Unix()
	if err := mysql.MysqlConnet.Model(&u).Where("id = ?", u.Id).Update(&u).Error;err!=nil{
		return err
	}
	return nil
}

func (u *User) UpdateData() error {
	if u.Account=="" {
		return errors.New("账号不能为空")
	}
	if u.GroupId==0 {
		return errors.New("用户组不能为空")
	}
	if u.ExpirationTime!="" {
		tick,err := time.Parse("2006-01-02",u.ExpirationTime)
		if err!=nil{
			tick, err := strconv.Atoi(u.ExpirationTime)
			if err != nil {
				return errors.New("过期时间格式有误")
			}
			u.Expiration = int64(tick)
		}
		u.Expiration=tick.Unix()
	}
	if u.Nickname=="" {
		u.Nickname=u.Account
	}
	u.UpdateTime=time.Now().Unix()
	if u.Id!=0 {
		if err := mysql.MysqlConnet.Model(&u).Where("id = ?", u.Id).Update(&u).Error;err!=nil{
			return err
		}
	}else {
		if u.MatchAccount() {
			return errors.New("账号已存在")
		}
		u.CreateTime=time.Now().Unix()
		u.Salt = common.GetRandSalt()
		u.Password = fmt.Sprintf("%x",md5.Sum([]byte(u.Account+u.Salt))) //初始密码为账号
		if err := mysql.MysqlConnet.Model(&u).Create(&u).Error;err!=nil{
			return err
		}
	}
	return nil
}

func (u *User) MatchAccount() bool {
	if u.Account=="" {
		return false
	}
	mysql.MysqlConnet.LogMode(true)
	if err := mysql.MysqlConnet.Model(&u).Where("account = ?", u.Account).First(&u).Error;err!=nil{
		return false
	}
	return true
}

func (u *User) MatchPassWord() bool {
	if u.Id==0 {
		return false
	}
	if u.Password=="" {
		return false
	}
	if u.Salt=="" {
		return false
	}
	if u.RePassword=="" {
		return false
	}
	if u.Password != fmt.Sprintf("%x",md5.Sum([]byte(u.RePassword+u.Salt))) {
		return false
	}
	return true
}

func (u *User) IncreaseLoginNum() error {
	if u.Id==0 {
		return errors.New("参数错误")
	}
	if err := mysql.MysqlConnet.Model(&u).Where("id = ?", u.Id).First(&u).Error;err!=nil{
		return err
	}
	u.LoginNum++
	if err := mysql.MysqlConnet.Model(&u).Where("id = ?", u.Id).Update(&u).Error;err!=nil{
		return err
	}
	return nil
}
