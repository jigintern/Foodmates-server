package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_"github.com/go-sql-driver/mysql"
	"log"
	"os"
	"time"
)

type table struct {
	Id	int
	Name	string
	CreatedAt	time.Time
	UpdateAt	time.Time
}


func EnvLoad() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}


func gormConnect() *gorm.DB {
	USER := "root"
	PASS := os.Getenv("MYSQL_ROOT_PASSWORD")
	PROTOCOL := "tcp(mysql_host:3306)"
	DBNAME := os.Getenv("MYSQL_DATABASE")

	CONNECT := USER+":"+PASS+"@"+PROTOCOL+"/"+DBNAME+"?charset=utf8&parseTime=True&loc=Local"
	db,err := gorm.Open("mysql", CONNECT)

	if err != nil {
		panic(err.Error())
	}
	return db
}


func main(){
	EnvLoad()
	db := gormConnect()
	db.LogMode(true)
	var testdb [] table
	record := db.Find(&testdb,"id=?",1)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": record,
		})
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": record,
		})
	})
	r.Run()
}
