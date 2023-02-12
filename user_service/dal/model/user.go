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
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}
