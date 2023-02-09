package service

import (
	"dousheng_server/video_service/dal/model"
	"testing"
)

func TestVideoCenter_Publish(t *testing.T) {
	err := VideoCenter{}.Publish(&model.Video{
		UUID:          0,
		UserID:        0,
		PlayURL:       "TestVideoCenter_Publish",
		CoverURL:      "TestVideoCenter_Publish",
		FavoriteCount: 0,
		CommentCount:  0,
		Title:         "TestVideoCenter_Publish",
	})
	if err != nil {
		return
	}
}
