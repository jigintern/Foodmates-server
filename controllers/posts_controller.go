package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jigintern/Foodmates-server/models"
	"net/http"
	"strconv"
)

// ReadPosts   GET "/api/v1/posts/readall"
func ReadAllPosts(ctx *gin.Context) {
	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	var post []models.Post
	db, err := models.GetDB()
	
	// DBがなければ500を返す
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError})
		return
	}
	db.Table("Posts").Find(&post)
	fmt.Println(post)
	ctx.JSON(http.StatusOK, post)
}

// ReadPost   GET "/api/v1/posts/read/:user_id"
func ReadSpecificUsersPost(ctx *gin.Context) {
	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	id, err := strconv.Atoi(ctx.Param("user_id"))
	if err != nil {
		return
	}
	var posts []models.Post
	db, err := models.GetDB()
	db.Table("Posts").Where("user_id = ?", id).Find(&posts)
	ctx.JSON(http.StatusOK, posts)
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
