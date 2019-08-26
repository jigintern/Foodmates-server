package models

import (
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

func EnvLoad() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func GormConnect() *gorm.DB {
	USER := os.Getenv("MYSQL_USER")
	PASS := os.Getenv("MYSQL_PASSWORD")
	PROTOCOL := "tcp(mysql_host:3306)"
	DBNAME := os.Getenv("MYSQL_DATABASE")

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", CONNECT)

	if err != nil {
		panic(err.Error())
	}
	
	EnvLoad()
	db := GormConnect()
	db.LogMode(true)
	return db
}
