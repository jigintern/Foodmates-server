package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Dish struct {
	gorm.Model
	DishId    int    `gorm:"default:0", json:"dish_id"`
	DishName  string `gorm:"default:''", json:"dish_name"`
	StoreName string `gorm:"default:''", json:"store_name"`
}
