package model

import "dousheng_server/user_service/dal/model"

type Favorite struct {
	UserId  int64 `json:"user_id" gorm:"primaryKey" `
	VideoId int64 `json:"video_id" gorm:"primaryKey" `
	//外键
	User model.User `gorm:"foreignKey:user_id"`
}
