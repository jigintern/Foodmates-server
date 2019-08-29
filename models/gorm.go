package models

import (
	"fmt"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"os"
)

var db *gorm.DB

func InitDB() {
	USER := os.Getenv("MYSQL_USER")
	PASS := os.Getenv("MYSQL_PASSWORD")
	PROTOCOL := "tcp(mysql_host:3306)"
	DBNAME := os.Getenv("MYSQL_DATABASE")

	var err error
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=True&loc=Local"

	fmt.Println("* Opening Mysql database...")
	db, err = gorm.Open("mysql", CONNECT)

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("* Mysql database opened!!")
	db.LogMode(true)
}

func GetDB() (*gorm.DB, error) {
	if db == nil {
		return nil, errors.New("database reference is null")
	}
	return db, nil
}
