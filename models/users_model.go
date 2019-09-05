package models

import (
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type User struct {
	ID          int       `gorm:"primary_key" json:"id"`
	LoginName   string    `gorm:"default:''" json:"login_name"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Name        string    `gorm:"default:''" json:"name"`
	Biography   string    `gorm:"default:''" json:"biography"`
	Birth       time.Time `json:"birth" sql:"type:date"`
	Country     string    `gorm:"default:''" json:"country"`
	Prefecture  string    `gorm:"default:''" json:"prefecture"`
	IconAddress string    `gorm:"default:''" json:"icon_address"`
}

type SignUpParams struct {
	LoginName   string    `gorm:"default:''" json:"login_name"`
	Name        string    `gorm:"default:''" json:"name"`
	Biography   string    `gorm:"default:''" json:"biography"`
	Birth       time.Time `json:"birth" sql:"type:date" time_format:"2006-01-02"`
	Country     string    `gorm:"default:''" json:"country"`
	Prefecture  string    `gorm:"default:''" json:"prefecture"`
	IconAddress string    `gorm:"default:''" json:"icon_address"`
}
