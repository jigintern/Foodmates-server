package models

import (
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type User struct {
	UserID      int       `gorm:"default:0", json:"user_id"`
	Name        string    `gorm:"default:''", json:"name"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Biography   string    `gorm:"default:''", json:"biography"`
	Country     string    `gorm:"default:''", json:"country"`
	Prefecture  string    `gorm:"default:''", json:"prefecture"`
	IconAddress string    `gorm:"default:''", json:"icon_address"`
}
