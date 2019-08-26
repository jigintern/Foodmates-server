package routers

import (
	"../controllers"
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
			posts.GET("/", controllers.ReadPosts)
			posts.POST("/", controllers.CreatePost)
		}
		users := api.Group("/users")
		{
			users.GET("/", controllers.GetUsers)
		}
	}

	return router
}
