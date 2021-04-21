package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"lhc.go.game.center/model"
	"net/http"
)

func GetList(c *gin.Context)  {
	page := model.NewPage()
	if err:=c.ShouldBind(&page);err!=nil {
		c.JSON(http.StatusOK,gin.H{"code":400,"msg":err.Error()})
		return
	}
	fmt.Println(1)
	fmt.Printf("%#v\n",page)
	fmt.Println(2)
	if page.Length == 0 {
		page.Length = 10
	}
	user := model.NewUser()
	if err:=c.ShouldBind(&user);err!=nil {
		c.JSON(http.StatusOK,gin.H{"code":400,"msg":err.Error()})
		return
	}
	data, total, err := user.GetList(page)
	if err!=nil {
		c.JSON(http.StatusOK,gin.H{"code":400,"msg":err.Error()})
		return
	}

	fmt.Printf("%#v\n",data)

	c.JSON(http.StatusOK,gin.H{"code":200,"data":data,"total":total})
}

func UpdateData(c *gin.Context){
	params := model.NewUser()
	if err:=c.ShouldBind(&params);err!=nil {
		c.JSON(http.StatusOK,gin.H{"code":400,"msg":err.Error()})
		return
	}
	if err:=params.UpdateData();err!=nil {
		c.JSON(http.StatusOK,gin.H{"code":400,"msg":err.Error()})
		return
	}
	c.JSON(http.StatusOK,gin.H{"code":200,"msg":"操作成功","url":"index"})
}

func Get(c *gin.Context){
	params := model.NewUser()
	if err:=c.ShouldBind(&params);err!=nil {
		c.JSON(http.StatusOK,gin.H{"code":400,"msg":err.Error()})
		return
	}
	if err:=params.GetOne();err!=nil {
		c.JSON(http.StatusOK,gin.H{"code":400,"msg":err.Error(),"data":nil})
		return
	}
	c.JSON(http.StatusOK,gin.H{"code":200,"msg":"查询成功","data":params})
}
