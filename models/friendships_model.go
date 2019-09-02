package models

import "time"

type FollowsData struct {
	UserID   int "json:user_id"
	FollowID int "json:follow_id"
}

type FollowsDBModel struct {
	ID        int       `gorm:"primary_key" json:"id"`
	UserID    int       `gorm:"default:0" json:"user_id"`
	FollowID  int       `gorm:"default:0" json:"follow_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
