package controllers

import "github.com/gin-gonic/gin"

func Send_message_controller(context *gin.Context) {
	context.JSON(200, gin.H{
		"code": 200,
		"msg":  "hello world",
	})
}
