package routers

import (
	"github.com/jigintern/Foodmates-server/controllers"
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
			createPost := posts.Group("/create")
			{
				createPost.POST("/", controllers.CreatePost)
			}
			readAllPosts := posts.Group("/readall")
			{
				readAllPosts.GET("/", controllers.ReadAllPosts)
			}
			readSpecificUsersPost := posts.Group("/read", controllers.ReadSpecificUsersPost)
			{
				readSpecificUsersPost.GET("/:id")
			}
		}
		dishes := api.Group("/dishes")
		{
			readAllDishes := dishes.Group("/readall")
			{
				readAllDishes.GET("/", controllers.ReadAllDishes)
			}
		}
		users := api.Group("/users")
		{
			users.GET("/:id", controllers.ReadUsers)
		}
		friendships := api.Group("/friendships")
		{
			createFriendship := friendships.Group("/create")
			{
				createFriendship.POST("/", controllers.CreateFriendships)
			}
			destroyFriendship := friendships.Group("/destroy")
			{
				destroyFriendship.POST("/", controllers.DestroyFriendships)
			}
		}
	}
	uploadPicture := router.Group("/upload")
	{
		uploadPicture.POST("/", controllers.UploadPicture)
	}

	return router
}
