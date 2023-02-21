package model

import "time"

type Comment struct {
	CommentId  int64  `gorm:"primaryKey" json:"comment_id"`
	UserId     int64  `gorm:"index:idx_user" json:"user_id"`
	VideoId    int64  `gorm:"index:idx_video" json:"video_id"`
	Content    string `json:"content"`
	CreateDate string `gorm:"index:idx_create,sort:desc" json:"create_date"`
	CreatedAt  time.Time
}
