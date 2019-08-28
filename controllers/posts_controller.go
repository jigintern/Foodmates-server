package controllers

import (
	"encoding/json"
	"github.com/jigintern/Foodmates-server/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
)

type PostData struct {
	UserID  int    `json:"user_id"`
	Content string `json:"comment"`
	DishID  int    `json:"dish_id"`
}

// ReadPosts   GET "/api/v1/posts/readall"
func ReadAllPosts(ctx *gin.Context) {
	var post []models.Post
	var db gorm.DB = *(models.GetDB())
	fmt.Printf("db_addr____controller: %v\n", db)
	db.Table("Posts").Find(&post)
	fmt.Println(post)
	poststr, err = json.Marshal(post)
	ctx.JSON(http.StatusOK, poststr)
}

// ReadPost		GET "/api/v1/posts/read"
func ReadSpecificUsersPost(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return
	}
	var posts []models.Post
	var db gorm.DB = *models.GetDB()
	db.Table("Posts").Where("user_id = ?", id).Find(&posts)
	ctx.JSON(http.StatusOK, posts)
}

// CreatePost   POST "/api/v1/posts"
func CreatePost(ctx *gin.Context) {
	var param models.Post
	err := ctx.BindJSON(&param)
	if err != nil {
		fmt.Println("======== request couldn't bind to json!! ========")
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "The request couldn't bind to json."})
		return
	}
	var db gorm.DB = *(models.GetDB())
	db.Table("Posts").Create(&param)
	ctx.JSON(http.StatusOK, gin.H{"data": param})
	fmt.Println("======== success!! ========")
	fmt.Println(param)
}
