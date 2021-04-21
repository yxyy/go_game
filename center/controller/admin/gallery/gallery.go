package gallery

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"lhc.go.game.center/model"
	"net/http"
)

func Index(c *gin.Context)  {
	c.HTML(http.StatusOK,"gallery/index.html",gin.H{"crawler_status":viper.Get("crawler_status")})
}

func GetList(c *gin.Context)  {
	var param  model.Params
	if err:=c.ShouldBind(&param);err!=nil {
		c.JSON(http.StatusOK,gin.H{"code":400,"msg":err.Error()})
		return
	}
	var params = model.NewNetbianImg()
	if err:=c.ShouldBind(&params);err!=nil {
		c.JSON(http.StatusOK,gin.H{"code":400,"msg":err.Error()})
		return
	}
	fmt.Println(params)
	data, total, err := params.GetList(param)
	if err!=nil {
		c.JSON(http.StatusOK,gin.H{"code":400,"msg":err.Error()})
		return
	}
	c.JSON(http.StatusOK,gin.H{"code":200,"data":data,"recordsTotal":total,"recordsFiltered":total})
}

func Update(c *gin.Context){
	var params = model.NewNetbianImg()
	if err:=c.ShouldBind(&params);err!=nil {
		c.JSON(http.StatusOK,gin.H{"code":400,"msg":err.Error()})
		return
	}
	if err:=params.GetOneById();err!=nil {
		c.JSON(http.StatusOK,gin.H{"code":400,"msg":err.Error()})
		return
	}
	c.JSON(http.StatusOK,gin.H{"code":200,"msg":"操作成功","url":"index"})
}
