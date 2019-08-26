package controllers

import (
	"github.com/gin-gonic/gin"
)

func GetUsers(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"user_name": "watano",
	})
}
