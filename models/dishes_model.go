package models

import (
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type Dish struct {
	ID           int       `gorm:"primary_key" json:"id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	DeletedAt    time.Time `json:"deleted_at"`
	DishName     string `gorm:"default:''" json:"dish_name"`
	StoreName    string `gorm:"default:''" json:"store_name"`
}
