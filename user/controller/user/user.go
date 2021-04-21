package user

import (
	"github.com/gin-gonic/gin"
	"lhc.go.game.user/model"
	"net/http"
)

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

func GetList(c *gin.Context)  {

	var ReturnData = model.NewReturnDataByRequestId(c.GetString("request_id"))
	defer func() {
		model.CatchErr(*ReturnData)
	}()
	page := model.NewPage()
	if ReturnData.Msg=c.ShouldBind(&page);ReturnData.Msg!=nil {
		ReturnData.Code = 400
		c.JSON(http.StatusOK,ReturnData)
		return
	}
	if page.Length == 0 {
		page.Length = 10
	}
	user := model.NewUser()
	if ReturnData.Msg=c.ShouldBind(&user);ReturnData.Msg!=nil {
		ReturnData.Code = 401
		c.JSON(http.StatusOK,ReturnData)
		return
	}
	data, total, err := user.GetList(page)
	if err!=nil {
		ReturnData.Code = 402
		ReturnData.Msg = err
		c.JSON(http.StatusOK,ReturnData)
		return
	}
	ReturnData.Code = 200
	//ReturnData.Msg = errors.New("查找成功")
	ReturnData.Data = data
	ReturnData.Total = total
	c.JSON(http.StatusOK,ReturnData)
}

func UpdateData(c *gin.Context){
	params := model.NewUser()
	if err:=c.ShouldBind(&params);err!=nil {
		c.JSON(http.StatusOK,gin.H{"code":400,"msg":err.Error()})
		return
	}
	if err:=params.UpdateData();err!=nil {
		c.JSON(http.StatusOK,gin.H{"code":401,"msg":err.Error()})
		return
	}
	c.JSON(http.StatusOK,gin.H{"code":200,"msg":"操作成功","url":"index"})
}


