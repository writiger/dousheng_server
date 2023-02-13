package service

import (
	"dousheng_server/uuidmaker"
	"dousheng_server/video_service/dal/model"
	"dousheng_server/video_service/dal/query"
	"dousheng_server/video_service/kitex_gen"
	"errors"
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

// VideoList 获取发布视频列表
func (vc VideoCenter) VideoList(userId int64) ([]*kitex_gen.Video, error) {
	videos, err := query.VideoList(userId)
	if err != nil {
		return nil, err
	}
	return modelToKitexVideo(videos), nil
}

// Like 点赞或取消
func (vc VideoCenter) Like(userId, videoId int64, actionType int32) error {
	switch actionType {
	case 1: // 点赞
		err := query.Favorite(userId, videoId)
		if err != nil {
			return err
		}
	case 2: // 取消
		err := query.UndoFavorite(userId, videoId)
		if err != nil {
			return err
		}
	default:
		return errors.New("wrong action type")
	}
	return nil
}

// GetVideo 获取视频
func (vc VideoCenter) GetVideo(uuid int64) (*kitex_gen.Video, error) {
	video, err := query.GetVideo(uuid)
	createTime := video.CreatedAt.UnixMilli()
	res := &kitex_gen.Video{
		Uuid:          video.UUID,
		UserId:        video.UserID,
		PlayUrl:       video.PlayURL,
		CoverUrl:      video.CoverURL,
		FavoriteCount: video.FavoriteCount,
		CommentCount:  video.CommentCount,
		Title:         video.Title,
		CreateTime:    createTime,
	}
	return res, err
}

// FavoriteList 获取点赞过的是视频列表
func (vc VideoCenter) FavoriteList(userId int64) ([]*kitex_gen.Video, error) {
	var res []*kitex_gen.Video
	videoIds, err := query.FavoriteList(userId)
	if err != nil {
		return nil, err
	}
	for _, item := range videoIds {
		temp, err := vc.GetVideo(item)
		if err != nil {
			return nil, err
		}
		res = append(res, temp)
	}
	return res, nil
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
