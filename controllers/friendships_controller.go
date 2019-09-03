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

func Follow(db *gorm.DB, userId int, followId int, ctx *gin.Context) error {
	var followsDatabase models.FollowsDBModel
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
	var followsDatabase models.FollowsDBModel
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
	var jsonData models.FollowsData
	err := ctx.BindJSON(&jsonData)
	if err != nil {
		log.Fatalln(err.Error())
	}
	db, err := models.GetDB()
	if jsonData.UserID != jsonData.FollowID {
		err = Follow(db, jsonData.UserID, jsonData.FollowID, ctx)
	}
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
	var jsonData models.FollowsData
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

func IsFollowing(ctx *gin.Context) {
	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	var request models.FollowsData
	var response models.FollowsDBModel
	err := ctx.BindJSON(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		log.Fatalln(err.Error())
	}
	db, err := models.GetDB()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
		log.Fatalln(err.Error())
	}
	count := 0
	db.Table("Follows").Where("user_id=? and follow_id=?", request.UserID, request.FollowID).First(&response).Count(&count)
	if count == 1 {
		ctx.JSON(http.StatusOK, gin.H{"following": true})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"following": false})
	}
}
