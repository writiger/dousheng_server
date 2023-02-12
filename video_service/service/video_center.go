package service

import (
	"dousheng_server/uuidmaker"
	"dousheng_server/video_service/dal/model"
	"dousheng_server/video_service/dal/query"
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
func (vc VideoCenter) Feed(timeStamp int64) (*[]model.Video, error) {
	lastTime := time.UnixMilli(timeStamp)
	return query.Feed(lastTime)
}
