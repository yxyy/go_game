package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"lhc.go.game.user/utitls/token"
	"log"
	"net/http"
)

func Auth(c *gin.Context)  {

	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusForbidden,gin.H{"code":404,"msg":"token 无效 , 拒绝访问"})
		c.Abort()
		return
	}

	parseToken, err := token.ParseToken(tokenString)
	if err!=nil {
		log.Println(err)
		c.JSON(http.StatusForbidden,gin.H{"code":403,"msg":err.Error()})
		c.Abort()
		return
	}

	//暂存数据
	c.Set("id",parseToken.Claims.(jwt.MapClaims)["id"])
	c.Set("account",parseToken.Claims.(jwt.MapClaims)["account"])
	c.Set("nickname",parseToken.Claims.(jwt.MapClaims)["nickname"])

	c.Next()

}
