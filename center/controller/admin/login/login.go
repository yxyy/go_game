package login

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"lhc.go.game.center/model"
	"log"
	"net/http"
	"time"
)

func Index(c *gin.Context)  {
	c.HTML(http.StatusOK,"admin/login/index.html",gin.H{})
}

func Login(c *gin.Context)  {
	var params model.User
	if err := c.ShouldBind(&params);err!=nil{
		c.JSON(http.StatusOK,gin.H{"code":500,"msg":"服务器繁忙，请稍候重试！"})
		return
	}
	if params.Account=="" {
		c.JSON(http.StatusOK,gin.H{"code":500,"msg":"账号错误"})
		return
	}
	if params.RePassword=="" {
		c.JSON(http.StatusOK,gin.H{"code":500,"msg":"密码错误"})
		return
	}
	if !params.MatchAccount() {
		c.JSON(http.StatusOK,gin.H{"code":500,"msg":"账号错误"})
		return
	}

	if !params.MatchPassWord() {
		c.JSON(http.StatusOK,gin.H{"code":500,"msg":"密码错误"})
		return
	}

	params.LoginNum++
	params.LastLoginTime=time.Now().Unix()
	params.LastLoginIp=c.ClientIP()
	if err := params.Update();err!=nil{
		log.Println(err)
	}

	session := sessions.Default(c)
	session.Set("user_id",params.Id)
	session.Set("account",params.Account)
	session.Set("nickname",params.Nickname)
	session.Set("group_id",params.GroupId)
	session.Set("expiration",params.Expiration)
	session.Save()

	c.JSON(http.StatusOK,gin.H{"code":200,"msg":"登陆成功","url":"index"})
}

func Logout(c *gin.Context)  {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.Redirect(http.StatusFound,"login")
}

type Name struct {
	
}
