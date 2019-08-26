package models

import (
//	"time"
	"github.com/jinzhu/gorm"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Post struct {
	gorm.Model
	UserId      int `gorm:"default:0"`
	DishId      int `gorm:"default:0"`
	Comment      string `gorm:"default:''"`
	ImageAddress   string `gorm:"default:''"`
}

func InitPostsModel() {
	EnvLoad()
	db := GormConnect()
	db.LogMode(true)
	var post []Post
	db.Table("Posts").Find(&post)
        fmt.Println(post)
}
