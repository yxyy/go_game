package token

import (
	"github.com/gin-gonic/gin"
	"lhc.go.game.user/model"
	"lhc.go.game.user/utitls/token"
	"net/http"
)

func RefreshToken (c *gin.Context)  {

	user := model.NewUser()
	user.Id = int(c.GetFloat64("id"))
	user.Nickname = c.GetString("account")
	user.Account = c.GetString("nickname")

	accessToken, err := token.CreateAccessToken(user)
	if err !=nil {
		c.JSON(http.StatusOK,gin.H{"code":503,"msg":"服务器繁忙"})
		return
	}
	c.JSON(http.StatusOK,gin.H{"code":200,"data":accessToken})
}
