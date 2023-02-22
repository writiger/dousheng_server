package model

type Message struct {
	//FromUserId对ToUserId发消息
	Id         int64  `gorm:"primaryKey" json:"id"`
	ToUserId   int64  `gorm:"to_user_id,index:idx_to" json:"to_user_id"`
	FromUserId int64  `gorm:"from_user_id,index:idx_from" json:"from_user_id"`
	Messages   string `gorm:"message" json:"content"`
	CreatedAt  int64  `gorm:"index:idx_create" json:"create_time"`
}
