package model

import (
	"dousheng_server/video_service/dal/model"
	"gorm.io/gorm"
	"time"
)

type User struct {
	UUID          int64         `gorm:"primaryKey" json:"id"`
	Username      string        `json:"name"`
	Password      string        `json:"password"`
	FollowCount   int64         `json:"follow_count"`
	FollowerCount int64         `json:"follower_count"`
	Videos        []model.Video `gorm:"foreignKey:user_id"`
	Followers     []Follower    `gorm:"foreignKey:user_id" gorm:"foreignKey:follow_id"`
	Messages      []Message     `gorm:"foreignKey:from_user_id" gorm:"foreignKey:to_user_id"`
	CreatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}
