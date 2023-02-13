package service

import (
	"dousheng_server/uuidmaker"
	"dousheng_server/video_service/dal/model"
	"dousheng_server/video_service/dal/query"
	"dousheng_server/video_service/kitex_gen"
	"fmt"
	"time"
)

type VideoCenter struct {
}

// Publish 投稿视频
func (vc VideoCenter) Publish(video *model.Video) (int64, error) {
	// 1. 申请UUID
	uuid, err := uuidmaker.GetUUID()
	if err != nil {
		return 0, err
	}
	video.UUID = uuid
	// 2. 保存
	saverIp := "http://192.168.31.84:8080"
	video.PlayURL = fmt.Sprintf(saverIp+"/static/videos/%d.%s", uuid, video.PlayURL)
	video.CoverURL = fmt.Sprintf(saverIp+"/static/covers/%d.%s", uuid, "png")
	err = query.CreateVideo(video)
	return uuid, err
}

// Delete 删除视频
func (vc VideoCenter) Delete(uuid int64) error {
	return query.DeleteVideo(uuid)
}

// Feed 获取视频
func (vc VideoCenter) Feed(timeStamp int64) ([]*kitex_gen.Video, error) {
	lastTime := time.UnixMilli(timeStamp)
	videos, err := query.Feed(lastTime)
	if err != nil {
		return nil, err
	}
	return modelToKitexVideo(videos), nil
}

// List 获取发布视频列表
func (vc VideoCenter) List(userId int64) ([]*kitex_gen.Video, error) {
	videos, err := query.List(userId)
	if err != nil {
		return nil, err
	}
	return modelToKitexVideo(videos), nil
}

// 将model中的video转换为kitex中生成的video
func modelToKitexVideo(videos *[]model.Video) []*kitex_gen.Video {
	var videoList []*kitex_gen.Video
	for _, item := range *videos {
		createTime := item.CreatedAt.UnixMilli()
		videoList = append(videoList, &kitex_gen.Video{
			Uuid:          item.UUID,
			UserId:        item.UserID,
			PlayUrl:       item.PlayURL,
			CoverUrl:      item.CoverURL,
			FavoriteCount: item.FavoriteCount,
			CommentCount:  item.CommentCount,
			Title:         item.Title,
			CreateTime:    createTime,
		})
	}
	return videoList
}
