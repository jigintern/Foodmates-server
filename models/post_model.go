package models

import (
	"time"
	_ "github.com/go-sql-driver/mysql"
)

type Post struct {
	ID           int       `gorm:"primary_key" json:"id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	UserId       int       `gorm:"default:0" json:"user_id"`
	DishId       int       `gorm:"default:0" json:"dish_id"`
	Comment      string    `gorm:"default:''" json:"comment"`
	ImageAddress string    `gorm:"default:''" json:"image_address"`
}

type PostResponse struct {
	Post
	User
	Dish
}