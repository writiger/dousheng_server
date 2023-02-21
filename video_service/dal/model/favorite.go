package model

type Favorite struct {
	UserId  int64 `gorm:"index:idx_user" json:"user_id"`
	VideoId int64 `json:"video_id"`
}
