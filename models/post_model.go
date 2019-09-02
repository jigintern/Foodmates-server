package models

import (
	_ "github.com/go-sql-driver/mysql"
	"time"
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
	Name        string    `gorm:"default:''" json:"user_name"`
	Biography   string    `gorm:"default:''" json:"biography"`
	Birth       time.Time `json:"birth" sql:"type:date"`
	Country     string    `gorm:"default:''" json:"country"`
	Prefecture  string    `gorm:"default:''" json:"prefecture"`
	IconAddress string    `gorm:"default:''" json:"icon_address"`
	DishName    string    `gorm:"default:''" json:"dish_name"`
	StoreName   string    `gorm:"default:''" json:"store_name"`
}
