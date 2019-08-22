package routers

import (
	v1 "../controllers/api/v1"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	api := router.Group("api/v1")
	{
		posts := api.Group("/posts")
		{
			posts.GET("/", v1.GetPosts)
		}
		users := api.Group("/users")
		{
			users.GET("/", v1.GetUsers)
		}
	}

	return router
}
