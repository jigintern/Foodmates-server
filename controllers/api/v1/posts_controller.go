package v1

import (
	"github.com/gin-gonic/gin"
)

var dummyData = []gin.H{
	{
		"user_name": "watano",
		"content":   "Hello, World!!",
	},
	{
		"user_name": "でみ",
		"content":   "Hello, golang!!",
	},
	{
		"user_name": "箒コウモリ",
		"content":   "Hello, Adobe!!",
	},
	{
		"user_name": "はたはた",
		"content":   "Hello, Vue!!",
	},
}

// GetPosts   GET "/api/v1/posts"
func GetPosts(ctx *gin.Context) {
	ctx.JSON(200, dummyData)
}
