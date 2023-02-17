package model

import (
	"gorm.io/gorm"
	"time"
)

type Message struct {
	// UserId 关注了 FollowID
	Id         int64  `gorm:"primaryKey" json:"id"`
	Messages   string `gorm:"message"`
	FromUserId int64  `gorm:"from_user_id"`
	ToUserId   int64  `gorm:"to_user_id"`
	CreatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}
