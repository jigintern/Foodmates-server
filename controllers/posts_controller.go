package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jigintern/Foodmates-server/models"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
)

// ReadPosts   GET "/api/v1/posts/readall"
func ReadAllPosts(ctx *gin.Context) {
	var post []models.Post
	var db gorm.DB = *(models.GetDB())
	db.Table("Posts").Find(&post)
	fmt.Println(post)
	ctx.JSON(http.StatusOK, post)
}

// ReadPost   GET "/api/v1/posts/read/:user_id"
func ReadSpecificUsersPost(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("user_id"))
	if err != nil {
		return
	}
	if id == 0 {
		ctx.JSON(http.StatusBadRequest, nil)
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
