package producter

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"lhc.go.game.log/model"
	"log"
	"net/http"
)

func Push(c *gin.Context)  {
	message := model.NewMessage()
	if err:= c.ShouldBind(&message);err!=nil {
		log.Println(err)
		c.JSON(http.StatusOK,gin.H{"code":-1,"msg":err.Error()})
		return
	}

	fmt.Printf("%#v\n",message)
	if err := message.Proucter();err!=nil {
		log.Println(err)
		c.JSON(http.StatusOK,gin.H{"code":-1,"msg":err.Error()})
		return
	}
}
