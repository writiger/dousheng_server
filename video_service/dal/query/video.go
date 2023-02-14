package query

import (
	"dousheng_server/video_service/dal/model"
	"time"
)

// CreateVideo 新增视频
func CreateVideo(video *model.Video) error {
	return GormClient.Create(video).Error
}

// DeleteVideo 软删除视频
func DeleteVideo(uuid int64) error {
	return GormClient.Where("uuid = ?", uuid).Delete(&model.Video{}).Error
}

// Feed 获取视频
func Feed(lastTime time.Time) (*[]model.Video, error) {
	var videos []model.Video
	err := GormClient.Limit(3).Order("created_at desc").Where("created_at < ?", lastTime).Find(&videos).Error
	return &videos, err
}

// VideoList 获取用户投稿过的视频列表
func VideoList(userId int64) (*[]model.Video, error) {
	var videos []model.Video
	err := GormClient.Where("user_id = ?", userId).Find(&videos).Error
	return &videos, err
}

// GetVideo 通过uuid获取video
func GetVideo(uuid int64) (*model.Video, error) {
	var res model.Video
	err := GormClient.Where("uuid = ?", uuid).Find(&res).Error
	return &res, err
}
