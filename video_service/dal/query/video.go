package query

import "dousheng_server/video_service/dal/model"

// CreateVideo 新增视频
func CreateVideo(video *model.Video) error {
	return GormClient.Create(video).Error
}

// DeleteVideo 软删除视频
func DeleteVideo(uuid int64) error {
	return GormClient.Where("uuid = ?", uuid).Delete(&model.Video{}).Error
}
