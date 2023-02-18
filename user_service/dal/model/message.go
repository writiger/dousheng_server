package model

type Message struct {
	// UserId 关注了 FollowID
	Id         int64  `gorm:"primaryKey" json:"id"`
	ToUserId   int64  `gorm:"to_user_id" json:"to_user_id"`
	FromUserId int64  `gorm:"from_user_id" json:"from_user_id"`
	Messages   string `gorm:"message" json:"content"`
	CreatedAt  int64  `json:"create_time"`
}
