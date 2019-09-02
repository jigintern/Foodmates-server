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

type FollowsData struct {
	UserID   int "json:user_id"
	FollowID int "json:follow_id"
}

type FollowsDBModel struct {
	UserID   int
	FollowID int
}

func Follow(db *gorm.DB, userId int, followId int, ctx *gin.Context) error {
	var followsDatabase FollowsDBModel
	followsDatabase.FollowID = followId
	followsDatabase.UserID = userId
	result := db.Table("Follows").Where(&followsDatabase).Create(&followsDatabase)
	ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
	fmt.Println("======== success!! ========")
	fmt.Println(result.Error)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func Unfollow(db *gorm.DB, userId int, followId int, ctx *gin.Context) error {
	var followsDatabase FollowsDBModel
	followsDatabase.FollowID = followId
	followsDatabase.UserID = userId
	err := db.Table("Follows").Where("user_id=? and follow_id=?", userId, followId).Delete(&followsDatabase).Error
	ctx.JSON(http.StatusOK, gin.H{"data": followId})
	fmt.Println("======== success!! ========")
	if err != nil {
		return err
	}
	return nil
}

func CreateFriendships(ctx *gin.Context) {
	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	var jsonData FollowsData
	err := ctx.BindJSON(&jsonData)
	if err != nil {
		log.Fatalln(err.Error())
	}
	db, err := models.GetDB()
	err = Follow(db, jsonData.UserID, jsonData.FollowID, ctx)
	if err != nil {
		log.Fatalln(err)
	}
}

func CountFriendhsips(ctx *gin.Context) {
	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		log.Fatalln(err.Error())
	}
	db, err := models.GetDB()
	count := 0
	db.Table("Follows").Where("user_id=?", id).Count(&count)
	ctx.JSON(http.StatusOK, gin.H{"following": count})
}

func DestroyFriendships(ctx *gin.Context) {
	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	var jsonData FollowsData
	err := ctx.BindJSON(&jsonData)
	if err != nil {
		log.Fatalln(err.Error())
	}
	db, err := models.GetDB()
	err = Unfollow(db, jsonData.UserID, jsonData.FollowID, ctx)
	if err != nil {
		log.Fatalln(err)
	}
}
