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
			posts.GET("/", v1.ReadPosts)
			posts.POST("/", v1.CreatePost)
		}
		users := api.Group("/users")
		{
			users.GET("/", v1.GetUsers)
		}
		friendships := api.Group("/friendships")
		{
			create := friendships.Group("/create")
			{
				create.POST("/", v1.Create)
			}
			destroy := friendships.Group("/destroy")
			{
				destroy.POST("/", v1.Destroy)
			}
		}
	}

	return router
}
