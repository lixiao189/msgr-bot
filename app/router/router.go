package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lixiao189/msgr-bot/app/controllers"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	messageApi := router.Group("/message")
	{
		messageApi.GET("/send", controllers.Send_message_controller)
	}

	return router
}
