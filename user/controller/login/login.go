package login

import (
	"github.com/gin-gonic/gin"
	"lhc.go.game.user/model"
	"lhc.go.game.user/utitls/token"
	"log"
	"net/http"
	"time"
)

func Login(c *gin.Context)  {
	var user = model.NewUser()
	if err := c.ShouldBind(&user);err!=nil{
		c.JSON(http.StatusOK,gin.H{"code":500,"msg":"服务器繁忙，请稍候重试！"})
		return
	}
	if user.Account=="" {
		c.JSON(http.StatusOK,gin.H{"code":400,"msg":"账号不能为空"})
		return
	}
	if user.RePassword=="" {
		c.JSON(http.StatusOK,gin.H{"code":401,"msg":"密码不能为空"})
		return
	}
	if !user.MatchAccount() {
		c.JSON(http.StatusOK,gin.H{"code":402,"msg":"账号错误"})
		return
	}

	if !user.MatchPassWord() {
		c.JSON(http.StatusOK,gin.H{"code":403,"msg":"密码错误"})
		return
	}

	user.LoginNum++
	user.LastLoginTime=time.Now().Unix()
	user.LastLoginIp=c.ClientIP()
	//user.ExpirationTime = time.Now().Add(time.Hour * 24 * 30 *3).Unix()
	if err := user.Update();err!=nil{
		log.Println(err)
	}

	//获取 access_token
	access_token , err := token.CreateAccessToken(user)
	if err !=nil {
		log.Println(err)
		c.JSON(http.StatusOK,gin.H{"code":500,"msg":"服务器繁忙","data":nil})
		return
	}
	//获取
	refresh_token, err := token.CreateRefreshToken(user)
	if err !=nil {
		log.Println(err)
		c.JSON(http.StatusOK,gin.H{"code":500,"msg":"服务器繁忙","data":nil})
		return
	}

	var data = make(map[string]string)
	data["access_token"] = access_token
	data["refresh_token"] = refresh_token

	c.JSON(http.StatusOK,gin.H{"code":200,"msg":"登陆成功","data":data , "userInfo" : user})
}

//func Logout(c *gin.Context)  {
//	session := sessions.Default(c)
//	session.Clear()
//	session.Save()
//	c.Redirect(http.StatusFound,"login")
//}