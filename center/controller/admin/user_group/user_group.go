package user_group

import (
	"github.com/gin-gonic/gin"
	"lhc.go.game.center/model"
	"net/http"
)

func Index(c *gin.Context)  {
	c.HTML(http.StatusOK,"user_group/index.html",gin.H{})
}

func Add(c *gin.Context)  {
	c.HTML(http.StatusOK,"user_group/add.html",gin.H{})
}

func GetUserGourpList(c *gin.Context)  {

	var sort model.Params
	if err:=c.ShouldBind(&sort);err!=nil {
		c.JSON(http.StatusOK,gin.H{"code":400,"msg":err.Error()})
		return
	}
	var params model.UserGroup
	if err:=c.ShouldBind(&params);err!=nil {
		c.JSON(http.StatusOK,gin.H{"code":400,"msg":err.Error()})
		return
	}
	data, total, err := params.GetUserGourpList(sort)
	if err!=nil {
		c.JSON(http.StatusOK,gin.H{"code":400,"msg":err.Error(),"data":nil})
		return
	}
	c.JSON(http.StatusOK,gin.H{"code":200,"data":data,"recordsTotal":total,"recordsFiltered":total})
}

func UpdateData(c *gin.Context)  {
	var params model.UserGroup
	if err:=c.ShouldBind(&params);err!=nil {
		c.JSON(http.StatusOK,gin.H{"code":400,"msg":err.Error()})
		return
	}
	if params.Name=="" {
		c.JSON(http.StatusOK,gin.H{"code":400,"msg":"名称不为空"})
		return
	}
	if params.Mark=="" {
		c.JSON(http.StatusOK,gin.H{"code":400,"msg":"标识不能为空"})
		return
	}
	if err:=params.UpdateData();err!=nil{
		c.JSON(http.StatusOK,gin.H{"code":400,"msg":err.Error()})
		return
	}
	c.JSON(http.StatusOK,gin.H{"code":200,"msg":"操作成功","url":"index"})
}