package internal

import (
	"github.com/gin-gonic/gin"
	"lhc.go.game.sdk/internal/pk5001"
	"lhc.go.game.sdk/model"
	"net/http"
)

func Callback(c *gin.Context)  {
	returnData := model.NewReturnData(c.GetString("request_id"))
	var callback  model.Order
	switch c.Param("package") {
	case "5001":
		pk5001 := pk5001.NewXingwanCallback()
		callback = pk5001
	default:
		returnData.Code = model.StatusParamError
		returnData.Msg = model.MsgIllegal
		c.JSON(http.StatusOK,returnData)
		return
	}
	if err:= model.AddOrder(callback);err!=nil {
		returnData.Code = model.StatusServerError
		returnData.Msg = err.Error()
		c.JSON(http.StatusOK,returnData)
		return
	}

	c.JSON(http.StatusOK,returnData)
}