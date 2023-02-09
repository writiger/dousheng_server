package dal

import (
	"dousheng_server/video_service/dal/model"
	"dousheng_server/video_service/dal/query"
	"testing"
)

func TestCreateVideo(t *testing.T) {
	t.Run("成功投稿视频测试", func(t *testing.T) {
		err := query.CreateVideo(&model.Video{
			UUID:          1,
			UserID:        629377243281358848,
			PlayURL:       "http//localhost:8080/static/videos/bear.mp4",
			CoverURL:      "http//localhost:8080/static/covers/bear.jpg?",
			FavoriteCount: 0,
			CommentCount:  0,
			Title:         "测试1",
		})
		if err != nil {
			t.Fatalf(err.Error())
			return
		}
	})
}

func TestDeleteVideo(t *testing.T) {
	t.Run("删除视频测试", func(t *testing.T) {
		err := query.DeleteVideo(630179868037873664)
		if err != nil {
			t.Fatalf(err.Error())
			return
		}
	})
}
