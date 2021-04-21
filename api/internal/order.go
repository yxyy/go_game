package internal

import (
	"github.com/gin-gonic/gin"
	"lhc.go.game.sdk/model"
	"net/http"
	"time"
)

func Order(c *gin.Context)  {
	returnData := model.NewReturnData(c.GetString("request_id"))
	tmp := model.NewOrderTmp(c.GetString("request_id"))

	if err := c.ShouldBind(&tmp);err!=nil{
		returnData.Code = model.StatusServerParseError
		returnData.Msg = err.Error()
		c.JSON(http.StatusOK,returnData)
		return
	}

	tmp.Ip = c.ClientIP()
	tmp.CreateTime = time.Now().Unix()

	if err := model.AddOrder(tmp);err !=nil {
		returnData.Code = model.StatusServerError
		returnData.Msg = err.Error()
		c.JSON(http.StatusOK,returnData)
		return
	}

	returnData.Data = map[string]interface{}{"orderId":tmp.OrderId}
	c.JSON(http.StatusOK,returnData)
}
