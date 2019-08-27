package routers

import (
	controllers "../controllers"
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
			users.GET("/", controllers.ReadUsers)
		}
		friendships := api.Group("/friendships")
		{
			create := friendships.Group("/create")
			{
				create.POST("/", controllers.CreateFriendships)
			}
			destroy := friendships.Group("/destroy")
			{
				destroy.POST("/", controllers.DestroyFriendships)
			}
		}
	}

	return router
}
