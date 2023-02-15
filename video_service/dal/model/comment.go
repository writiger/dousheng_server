package model

import "time"

type Comment struct {
	CommentId  int64  `gorm:"primaryKey" json:"comment_id"`
	UserId     int64  `json:"user_id"`
	VideoId    int64  `json:"video_id"`
	Content    string `json:"content"`
	CreateDate string `json:"create_date"`
	CreatedAt  time.Time
}
