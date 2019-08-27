package models

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func InitDB() {
	USER := os.Getenv("MYSQL_USER")
	PASS := os.Getenv("MYSQL_PASSWORD")
	PROTOCOL := "tcp(mysql_host:3306)"
	DBNAME := os.Getenv("MYSQL_DATABASE")

	var err error
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=True&loc=Local"
	db, err = gorm.Open("mysql", CONNECT)

	if err != nil {
		panic(err.Error())
	}
	
	fmt.Printf("db_addr: %v\n", db)
	db.LogMode(true)
}

func GetDB() *gorm.DB {
	fmt.Printf("db_addr: %v\n", db)
	return db
}