package model

import "dousheng_server/user_service/dal/model"

type Favorite struct {
	UserId  int64 `gorm:"index:idx_user" json:"user_id"`
	VideoId int64 `json:"video_id"`
	//外键
	User model.User `gorm:"foreignKey:user_id"`
}
