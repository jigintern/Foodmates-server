package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/jigintern/Foodmates-server/controllers"
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
				readSpecificUsersPost.GET("/:user_id")
			}
			readSuggestUser := posts.Group("/suggest")
			{
				readSuggestUser.GET("/:id", controllers.SuggestUser)
			}
		}
		dishes := api.Group("/dishes")
		{
			readAllDishes := dishes.Group("/readall")
			{
				readAllDishes.GET("/", controllers.ReadAllDishes)
			}
			createDish := dishes.Group("/create")
			{
				createDish.POST("/", controllers.CreateDish)
			}
		}
		users := api.Group("/users")
		{
			users.GET("/:id", controllers.ReadSpecificUser)
			users.GET("/", controllers.ReadAllUsers)
			createUser := users.Group("/signup")
			{
				createUser.POST("/", controllers.SignUp)
			}
			signIn := users.Group("/signin")
			{
				signIn.POST("/", controllers.SignIn)
			}
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
			countFriendship := friendships.Group("/count")
			{
				countFriendship.GET("/:id", controllers.CountFriendhsips)
			}
			isFollowing := friendships.Group("/isfollowing")
			{
				isFollowing.POST("/", controllers.IsFollowing)
			}
		}
		upload := api.Group("/upload")
		{
			uploadPicture := upload.Group("/picture")
			{
				uploadPicture.POST("/", controllers.UploadPicture)
			}
			uploadIcon := upload.Group("/icon")
			{
				uploadIcon.POST("/", controllers.UploadIcon)
			}
		}
	}

	return router
}
