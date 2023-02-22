package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	UUID          int64      `gorm:"primaryKey" json:"id"`
	Username      string     `json:"name"`
	Password      string     `json:"password"`
	FollowCount   int64      `json:"follow_count"`
	FollowerCount int64      `json:"follower_count"`
	Followers     []Follower `gorm:"foreignKey:user_id;foreignKey:follow_id"`
	Messages      []Message  `gorm:"foreignKey:from_user_id;foreignKey:to_user_id"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}
