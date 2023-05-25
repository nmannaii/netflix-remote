package controllers

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-vgo/robotgo"
	"github.com/nmannaii/fyne-netflix-remote/backend/dto"
)

func pressKey(context *gin.Context) {
	var keyPressReq dto.KeyPressRequest
	err := context.BindJSON(&keyPressReq)
	if err != nil {
		log.Print(err)
	}
	context.JSON(200, gin.H{
		"Success": robotgo.KeyTap(keyPressReq.Key),
	})
}

func mouseClick(context *gin.Context) {
	robotgo.Click()
	context.JSON(200, gin.H{
		"Success": true,
	})
}
func GroupKeyPressRoutes(r *gin.Engine) {
	r.POST("/key-press", pressKey)
	r.POST("/mouse-click", mouseClick)
}
