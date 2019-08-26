package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"log"
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


func CreateFriendships(ctx *gin.Context){
	EnvLoad()
	var jsonData FollowsData
	err := ctx.BindJSON(&jsonData)
	if err != nil{
		log.Fatalln(err.Error())
	}
	db := gormConnect()
	db.LogMode(true)
	var followsDatabase FollowsDBModel
	db.Table("Follows").Where("user_id=? and follow_id=?", jsonData.UserID, jsonData.FollowID).Create(&followsDatabase)
	ctx.JSON(200, gin.H{"data":jsonData.FollowID})
	fmt.Println("======== success!! ========")
}

func DestroyFriendships(ctx *gin.Context){
	EnvLoad()
	var jsonData FollowsData
	err := ctx.BindJSON(&jsonData)
	if err != nil{
		log.Fatalln(err.Error())
	}
	db := gormConnect()
	db.LogMode(true)
	var followsDatabase FollowsDBModel
	db.Table("Follows").Where("user_id=? and follow_id=?", jsonData.UserID, jsonData.FollowID).Delete(&followsDatabase)
	ctx.JSON(200, gin.H{"data":jsonData.FollowID})
	fmt.Println("======== success!! ========")
}