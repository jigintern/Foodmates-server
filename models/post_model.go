package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Post struct {
	gorm.Model
	UserId       int    `gorm:"default:0", json:"user_id"`
	DishId       int    `gorm:"default:0", json:"dish_id"`
	Comment      string `gorm:"default:''", json:"comment"`
	ImageAddress string `gorm:"default:''", json:"image_address"`
}
