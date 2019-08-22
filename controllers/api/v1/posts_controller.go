package v1

import (
	"github.com/gin-gonic/gin"
)

func GetPosts(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"content": "Hello, World!!",
	})
}
