package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-vgo/robotgo"
	"github.com/nmannaii/go-netflix-remote/dto"
	"log"
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

func GroupKeyPressRoutes(r *gin.Engine) {
	r.POST("/key-press", pressKey)
}
