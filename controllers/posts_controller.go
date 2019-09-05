package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jigintern/Foodmates-server/models"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"strconv"
)

func Ints(input []int) []int {
	u := make([]int, 0, len(input))
	m := make(map[int]bool)

	for _, val := range input {
		if _, ok := m[val]; !ok {
			m[val] = true
			u = append(u, val)
		}
	}

	return u
}

// ReadPosts   GET "/api/v1/posts/readall"
func ReadAllPosts(ctx *gin.Context) {
	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	var results []models.PostResponse
	db, err := models.GetDB()

	// DBがなければ500を返す
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError})
		return
	}
	db.Raw("SELECT Posts.id, Posts.`user_id`, Posts.`created_at`, Posts.`updated_at`, Posts.`dish_id`, Posts.`comment`, Posts.`image_address`, Users.`login_name`, Users.`name`, Users.`biography`, Users.`birth`, Users.`country`, Users.`prefecture`, Users.`icon_address`, Dishes.`dish_name`, Dishes.`store_name` FROM `Posts` LEFT OUTER JOIN `Users` ON `Posts`.`user_id` = `Users`.`id` LEFT OUTER JOIN `Dishes` ON `Posts`.`dish_id` = `Dishes`.`id` ORDER BY Posts.created_at DESC").Scan(&results)
	ctx.JSON(http.StatusOK, results)
}

// ReadPost   GET "/api/v1/posts/read/:user_id"
func ReadSpecificUsersPost(ctx *gin.Context) {
	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	loginName := ctx.Param("login_name")
	var results []models.PostResponse
	db, err := models.GetDB()
	// DBがなければ500を返す
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError})
		return
	}
	var userData models.User
	recordNotFound := db.Table("Users").Where("login_name=?", loginName).First(&userData).RecordNotFound()
	if recordNotFound {
		ctx.JSON(http.StatusInternalServerError, nil)
	} else {
		db.Raw("SELECT Posts.id, Posts.`user_id`, Posts.`created_at`, Posts.`updated_at`, Posts.`dish_id`, Posts.`comment`, Posts.`image_address`, Users.`login_name`, Users.`name`, Users.`biography`, Users.`birth`, Users.`country`, Users.`prefecture`, Users.`icon_address`, Dishes.`dish_name`, Dishes.`store_name` FROM `Posts` LEFT OUTER JOIN `Users` ON `Posts`.`user_id` = `Users`.`id` LEFT OUTER JOIN `Dishes` ON `Posts`.`dish_id` = `Dishes`.`id` WHERE Posts.user_id = ? ORDER BY Posts.created_at DESC", userData.ID).Scan(&results)
		ctx.JSON(http.StatusOK, results)
	}
}

// CreatePost   POST "/api/v1/posts"
func CreatePost(ctx *gin.Context) {
	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	var param models.Post
	err := ctx.BindJSON(&param)
	if err != nil {
		fmt.Println("======== request couldn't bind to json!! ========")
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "The request couldn't bind to json."})
		return
	}

	db, err := models.GetDB()
	tx := db.Table("Posts").Begin()
	tx.Create(&param)
	if tx.Error != nil {
		fmt.Println("\x1b[31mstarting transition failed.\x1b[0m")
		tx.Rollback()
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "starting transition failed."})
		return
	}
	//db.Table("Posts").Create(&param)
	if len(db.GetErrors()) != 0 {
		fmt.Println("\x1b[31mchanging database failed.\x1b[0m")
		tx.Rollback()
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "changing database failed."})
		return
	}

	tx.Commit()
	ctx.JSON(http.StatusOK, param)
	fmt.Println("\x1b[32msuccess!!\x1b[0m")
	fmt.Println(param)
}

func Suggest(id int, db *gorm.DB) []models.Suggest {
	var records []models.Post
	db.Table("Posts").Where("user_id=?", id).Order("created_at desc").Find(&records)
	count := 0
	db.Table("Users").Count(&count)
	var suggestUsers []models.Suggest
	var dishes []int
	var users []int
	for _, record := range records {
		if record.DishId != 0 {
			var r []models.Post
			dishes = append(dishes, record.DishId)
			recordNotFound := db.Table("Posts").Where("not(user_id=?) and dish_id = ?", id, record.DishId).Group("user_id").Find(&r).RecordNotFound()
			if !recordNotFound {
				for _, user := range r {
					users = append(users, user.UserId)
				}
			}
		}
	}
	users = Ints(users)
	for _, user := range users {
		count := 0
		var r []models.Post
		db.Table("Posts").Where("user_id=? and dish_id in (?) and not(user_id=?)", user, dishes, id).Order("created_at desc").Find(&r).Count(&count)
		if count != 0 {
			var s models.Suggest
			s.UserId = user
			s.Times = count
			suggestUsers = append(suggestUsers, s)
		}
	}
	return suggestUsers
}

// SuggestUser   GET "/api/v1/posts/suggest/:id"
func SuggestUser(ctx *gin.Context) {
	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		log.Fatalln(err.Error())
	}
	db, err := models.GetDB()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
		log.Fatalln(err.Error())
	}
	response := Suggest(id, db)
	ctx.JSON(http.StatusOK, response)
}
