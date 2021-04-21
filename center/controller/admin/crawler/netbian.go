package crawler

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"lhc.go.game.center/controller/corn/netbian"
	"lhc.go.game.center/model"
	"net/http"
)

func Index(c *gin.Context)  {
	c.HTML(http.StatusOK,"netbian/index.html",gin.H{"crawler_status":viper.Get("crawler_status")})
}

func GetList(c *gin.Context)  {
	var params = model.NewCrawlerMenu()
	if err:=c.ShouldBind(&params);err!=nil {
		c.JSON(http.StatusOK,gin.H{"code":400,"msg":err.Error()})
		return
	}
	data, total, err := params.GetList()
	if err!=nil {
		c.JSON(http.StatusOK,gin.H{"code":400,"msg":err.Error()})
		return
	}
	c.JSON(http.StatusOK,gin.H{"code":200,"data":data,"recordsTotal":total,"recordsFiltered":total})
}

func GetOneCategory(c *gin.Context)  {
	var params = model.NewCrawlerMenu()
	if err:=c.ShouldBind(&params);err!=nil {
		c.JSON(http.StatusOK,gin.H{"code":400,"msg":err.Error()})
		return
	}
	if err := params.GetOneCrawlerMenu();err!=nil {
		c.JSON(http.StatusOK,gin.H{"code":400,"msg":err.Error()})
		return
	}
	c.JSON(http.StatusOK,gin.H{"code":200,"data":params})
}

func UpdateData(c *gin.Context){
	var params = model.NewCrawlerMenu()
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

func Update(c *gin.Context){
	var params = model.NewCrawlerMenu()
	if err:=c.ShouldBind(&params);err!=nil {
		c.JSON(http.StatusOK,gin.H{"code":400,"msg":err.Error()})
		return
	}
	if err:=params.Update();err!=nil {
		c.JSON(http.StatusOK,gin.H{"code":400,"msg":err.Error()})
		return
	}
	c.JSON(http.StatusOK,gin.H{"code":200,"msg":"操作成功","url":"index"})
}

func Crawler(c *gin.Context)  {
	var params = model.NewCrawlerMenu()
	if err:=c.ShouldBind(&params);err!=nil {
		c.JSON(http.StatusOK,gin.H{"code":400,"msg":err.Error()})
		return
	}
	if err := params.GetOneCrawlerMenu();err!=nil {
		c.JSON(http.StatusOK,gin.H{"code":400,"msg":err.Error()})
		return
	}
	go func() {
		netbian.NewNetbian().Run()
	}()
	c.JSON(http.StatusOK,gin.H{"code":200,"msg":"爬虫已启动"})
}

