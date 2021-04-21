package menu

import (
	//"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"lhc.go.game.center/model"
	"net/http"
)

func Index(c *gin.Context)  {
	menu := model.NewMenu()
	if err := c.ShouldBind(&menu);err!=nil{
		c.JSON(http.StatusOK,gin.H{"code":http.StatusInternalServerError,"msg":"服务器开小差了。。。"})
		return
	}
	menuList := menu.GetMeunTree(0)
	c.JSON(http.StatusOK,gin.H{"code":http.StatusOK,"msg":"数据查找成功","data":menuList})
}

//func Index(c *gin.Context)  {
//	menuList := model.GetMeunTree(0)
//	c.JSON(http.StatusOK,gin.H{"menuList":menuList})
//}

func GetOneMenu(c *gin.Context)  {
	menu := model.NewMenu()
	if err := c.ShouldBind(&menu);err!=nil{
		c.JSON(http.StatusOK,gin.H{"code":http.StatusInternalServerError,"msg":"服务器开小差了。。。"})
		return
	}
	if err := menu.Get();err!=nil{
		c.JSON(http.StatusOK,gin.H{"code":http.StatusBadRequest,"msg":"没找到数据"})
		return
	}
	c.JSON(http.StatusOK,gin.H{"code":http.StatusOK,"msg":"数据查找成功","data":menu})
}

func GetMenuParent(c *gin.Context)  {
	menuList := model.GetParentMenu(0)
	c.JSON(http.StatusOK,gin.H{"code":http.StatusOK,"menuList":menuList})
}

func UpdateData(c *gin.Context)  {
	var params model.Menu
	if err:=c.ShouldBind(&params);err!=nil {
		c.JSON(http.StatusOK,gin.H{"code":400,"msg":err.Error()})
		return
	}
	if params.Name=="" {
		c.JSON(http.StatusOK,gin.H{"code":400,"msg":"名称不为空"})
		return
	}
	if params.Url=="" {
		c.JSON(http.StatusOK,gin.H{"code":400,"msg":"url不能为空"})
		return
	}
	//session := sessions.Default(c)
	//params.Author= session.Get("account").(string)
	params.Author= "yxyy"
	if err:=params.UpdateData();err!=nil{
		c.JSON(http.StatusOK,gin.H{"code":400,"msg":err.Error()})
		return
	}
	c.JSON(http.StatusOK,gin.H{"code":200,"msg":"操作成功","url":"index"})
}