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
			create := posts.Group("/create")
			{
				create.POST("/", controllers.CreatePost)
			}
			readAll := posts.Group("/readall")
			{
				readAll.GET("/", controllers.ReadAllPosts)
			}
			read := posts.Group("/read", controllers.ReadSpecificUsersPost)
			{
				read.GET("/:id")
			}
		}
		users := api.Group("/users")
		{
			users.GET("/:id", controllers.ReadUsers)
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
	upload := router.Group("/upload")
	{
		upload.POST("/", controllers.UploadPicture)
	}

	return router
}
