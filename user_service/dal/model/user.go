package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	UUID          int64  `json:"id"`
	Username      string `json:"username"`
	Password      string `json:"password"`
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `json:"follower_count"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}
