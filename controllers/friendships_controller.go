package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

type FollowsData struct {
	UserID int "json:user_id"
	FollowID int "json:follow_id"
}


type FollowsDBModel struct {
	UserID int
	FollowID int
}

func EnvLoad() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func gormConnect() *gorm.DB {
	EnvLoad()
	USER := os.Getenv("MYSQL_USER")
	PASS := os.Getenv("MYSQL_PASSWORD")
	PROTOCOL := "tcp(t2.intern.jigd.info:3306)"
	DBNAME := os.Getenv("USERS_DATABASE")

	CONNECT := USER+":"+PASS+"@"+PROTOCOL+"/"+DBNAME+"?charset=utf8&parseTime=True&loc=Local"
	db,err := gorm.Open("mysql", CONNECT)

	if err != nil {
		panic(err.Error())
	}
	return db
}

func Follow(db *gorm.DB, userId int, followId int, ctx *gin.Context) error{
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

func Unfollow(db *gorm.DB,userId int,followId int, ctx *gin.Context) error{
	var followsDatabase FollowsDBModel
	followsDatabase.FollowID = followId
	followsDatabase.UserID = userId
	err := db.Table("Follows").Where("user_id=? and follow_id=?", userId, followId).Delete(&followsDatabase).Error
	ctx.JSON(http.StatusOK, gin.H{"data":followId})
	fmt.Println("======== success!! ========")
	if err != nil {
		return err
	}
	return nil
}


func CreateFriendships(ctx *gin.Context){
	var jsonData FollowsData
	err := ctx.BindJSON(&jsonData)
	if err != nil{
		log.Fatalln(err.Error())
	}
	db := gormConnect()
	db.LogMode(true)
	err = Follow(db,jsonData.UserID, jsonData.FollowID, ctx)
	if err != nil {
		log.Fatalln(err)
	}
}

func DestroyFriendships(ctx *gin.Context){
	var jsonData FollowsData
	err := ctx.BindJSON(&jsonData)
	if err != nil{
		log.Fatalln(err.Error())
	}
	db := gormConnect()
	db.LogMode(true)
	err = Unfollow(db,jsonData.UserID, jsonData.FollowID, ctx)
	if err != nil {
		log.Fatalln(err)
	}
}
