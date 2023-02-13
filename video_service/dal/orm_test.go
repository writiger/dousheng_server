package dal

import (
	"dousheng_server/video_service/dal/model"
	"dousheng_server/video_service/dal/query"
	"fmt"
	"testing"
	"time"
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

func TestFeed(t *testing.T) {
	t.Run("Feed测试", func(t *testing.T) {
		lastTime := time.Now()
		videos, err := query.Feed(lastTime)
		if err != nil {
			t.Fatalf(err.Error())
			return
		}
		for _, item := range *videos {
			fmt.Println(item)
		}
	})
}

func TestList(t *testing.T) {
	t.Run("测试投稿列表", func(t *testing.T) {
		list, err := query.List(629648834103869440)
		if err != nil {
			t.Fatalf(err.Error())
			return
		}
		for _, item := range *list {
			fmt.Println(item)
		}
	})
}
