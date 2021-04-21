package internal

import (
	"errors"
	"github.com/gin-gonic/gin"
	"lhc.go.game.sdk/internal/pk5001"
	"lhc.go.game.sdk/model"
	"log"
	"net/http"
)

func Login(c *gin.Context)  {

	var login model.Loginer
	returnData := model.NewReturnData(c.GetString("request_id"))
	switch c.Param("package") {
	case "5001":
		var pck5001 = pk5001.NewXingwanLogin()
		if err := c.ShouldBind(&pck5001);err!=nil {
			returnData.Code = model.StatusServerError
			returnData.Msgs = errors.New("系统繁忙")
			c.JSON(http.StatusOK,returnData)
			log.Println(err)
			return
		}
		login = pck5001
	default:
		returnData.Code = model.StatusParamError
		returnData.Msgs = errors.New("无效的包")
		c.JSON(http.StatusOK,returnData)
		return
	}

	data, err := model.Login(login)
	if err!=nil {
		returnData.Code = model.StatusParamError
		returnData.Msg = err.Error()
		c.JSON(http.StatusOK,returnData)
		log.Println(err)
		return
	}
	returnData.Data = data
	c.JSON(http.StatusOK,returnData)
}
