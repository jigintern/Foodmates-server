package controllers

import (
	"../models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type PostData struct {
	UserID  int    `json:"user_id"`
	Content string `json:"comment"`
	DishID  int    `json:"dish_id"`
}

// ReadPosts   GET "/api/v1/posts"
func ReadPosts(ctx *gin.Context) {
	var post []models.Post
	var db gorm.DB = *(models.GetDB())
	fmt.Printf("db_addr____controller: %v\n", db)
	db.Table("Posts").Find(&post)
	fmt.Println(post)
	ctx.JSON(200, post)
}

// CreatePost   POST "/api/v1/posts"
func CreatePost(ctx *gin.Context) {
	var param models.Post
	err := ctx.BindJSON(&param)
	if err != nil {
		fmt.Println("======== request couldn't bind to json!! ========")
		ctx.JSON(400, gin.H{"status": "The request couldn't bind to json."})
		return
	}
	var db gorm.DB = *(models.GetDB())
	db.Table("Posts").Create(&param)
	ctx.JSON(200, gin.H{"data": param})
	fmt.Println("======== success!! ========")
	fmt.Println(param)
}
