package model

type Favorite struct {
	UserId  int64 `json:"user_id"`
	VideoId int64 `json:"video_id"`
}
