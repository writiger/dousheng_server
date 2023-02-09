package service

import (
	"dousheng_server/uuidmaker"
	"dousheng_server/video_service/dal/model"
	"dousheng_server/video_service/dal/query"
	"fmt"
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
	video.PlayURL = fmt.Sprintf("http://localhost:8080/static/videos/%d.%s", uuid, video.PlayURL)
	err = query.CreateVideo(video)
	return uuid, err
}

// Delete 删除视频
func (vc VideoCenter) Delete(uuid int64) error {
	return query.DeleteVideo(uuid)
}
