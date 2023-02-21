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
			UserID:        1,
			PlayURL:       "http://localhost:8080/static/videos/bear.mp4",
			CoverURL:      "http://localhost:8080/static/covers/bear.png",
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
		list, err := query.VideoList(629648834103869440)
		if err != nil {
			t.Fatalf(err.Error())
			return
		}
		for _, item := range *list {
			fmt.Println(item)
		}
	})
}

func TestCreateFavorite(t *testing.T) {
	t.Run("点赞", func(t *testing.T) {
		err := query.Favorite(629648834103869440, 631256877056917504)
		if err != nil {
			t.Fatalf(err.Error())
			return
		}
	})
}

func TestUndoFavorite(t *testing.T) {
	t.Run("取消点赞", func(t *testing.T) {
		err := query.UndoFavorite(629648834103869440, 631256877056917504)
		if err != nil {
			t.Fatalf(err.Error())
			return
		}
	})
}

func TestIsLiked(t *testing.T) {
	t.Run("是否点赞", func(t *testing.T) {
		res := query.IsLiked(629648834103869440, 631256877056917504)

		fmt.Println(res)
	})
}

func TestFavoriteList(t *testing.T) {
	t.Run("喜欢列表测试", func(t *testing.T) {
		list, err := query.FavoriteList(629648834103869440)
		if err != nil {
			t.Fatalf(err.Error())
			return
		}
		fmt.Println(list)
	})
}

func TestCreateComment(t *testing.T) {
	t.Run("发表评论测试", func(t *testing.T) {
		comment, err := query.CreateComment(&model.Comment{
			CommentId:  629648834103869443,
			UserId:     629648834103869440,
			VideoId:    631256877056917504,
			Content:    "发表评论测试2",
			CreateDate: time.Now().Format("01-02"),
		})
		if err != nil {
			t.Fatalf(err.Error())
			return
		}
		fmt.Println(comment)
	})
}

func TestDeleteComment(t *testing.T) {
	t.Run("删除评论测试", func(t *testing.T) {
		err := query.DeleteComment(629648834103869442)
		if err != nil {
			t.Fatalf(err.Error())
			return
		}
	})
}

func TestGetComment(t *testing.T) {
	t.Run("获取视频评论", func(t *testing.T) {
		comments, err := query.GetComment(631256877056917504)
		if err != nil {
			t.Fatalf(err.Error())
			return
		}
		fmt.Println(comments)
	})
}
