package models

import (
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type User struct {
	ID          int       `gorm:"primary_key" json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Name        string    `gorm:"default:''" json:"user_name"`
	Biography   string    `gorm:"default:''" json:"biography"`
	Country     string    `gorm:"default:''" json:"country"`
	Prefecture  string    `gorm:"default:''" json:"prefecture"`
	IconAddress string    `gorm:"default:''" json:"icon_address"`
	Birth       time.Time `json:"birth"`
}
