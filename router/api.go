package router

import (
	"github.com/gin-gonic/gin"
	"web/controllers"
)

func SetRouter (group *gin.RouterGroup) *gin.RouterGroup{

	// group.Use(controllers.Auth)
	group.POST("/new_account",controllers.NewUser)
	group.POST("/update",controllers.Update)
	group.POST("/get_balance",controllers.GetBalance)
	group.POST("/check",controllers.Check)
	group.POST("/import",controllers.ImportKey)
	group.POST("/send",controllers.Send)
	group.GET("/Price",controllers.GetPrice)
	group.POST("/query",controllers.Query)
	group.POST("/get_transaction",controllers.GetTransaction)
	return group
}
