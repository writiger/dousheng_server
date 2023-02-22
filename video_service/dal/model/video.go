package model

import (
	"dousheng_server/user_service/dal/model"
	"gorm.io/gorm"
	"time"
)

type Video struct {
	UUID          int64  `gorm:"primaryKey" json:"id"`
	UserID        int64  `json:"user_id"`
	PlayURL       string `json:"play_url"`
	CoverURL      string `json:"cover_url"`
	FavoriteCount int64  `json:"favorite_count"`
	CommentCount  int64  `json:"comment_count"`
	Title         string `json:"title"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
	//外键
	User model.User `gorm:"foreignKey:user_id"`
}
