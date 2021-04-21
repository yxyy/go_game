package admin

import (
	"github.com/gin-gonic/gin"
	"lhc.go.game.center/model"
	"net/http"
)

func Index(c *gin.Context)  {
	menuList := model.GetMeunTree(0)
	c.HTML(http.StatusOK,"admin/index.html",gin.H{"menuList":menuList})
}

func Welcome(c *gin.Context)  {
	menuList := model.GetMeunTree(0)
	c.HTML(http.StatusOK,"admin/welcome.html",gin.H{"menuList":menuList})
}
