package models

import (
	"github.com/jinzhu/gorm"

	_ "github.com/go-sql-driver/mysql"
)

type Post struct {
	gorm.Model
	// created_at   time.Date
	// updated_at   time.Date
	// name         string
	// biography    string
	// birth        time.Date
	// country      string
	// prefecture   string
	// icon_address string
}

func InitPostsModel() {
	EnvLoad()
	db := GormConnect()
	db.LogMode(true)
	var post Post
	db.Table("Posts").Create(&post)
}
