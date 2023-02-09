package rpc

import (
	"context"
	"dousheng_server/video_service/dal/model"
	"dousheng_server/video_service/kitex_gen"
	"dousheng_server/video_service/kitex_gen/videocenter"
	"github.com/cloudwego/kitex/client"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var videoClient videocenter.Client

func init() {
	// 通过etcd发现服务
	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	if err != nil {
		panic(err)
	}

	videoClient, err = videocenter.NewClient(
		"videoservice",
		client.WithResolver(r),
		client.WithSuite(tracing.NewClientSuite()))
	if err != nil {
		panic(err)
	}
}

// PublishVideo .
func PublishVideo(video *model.Video) (int64, error) {
	req := kitex_gen.PublishRequest{
		UserId:   video.UserID,
		PlayUrl:  video.PlayURL,
		CoverUrl: video.CoverURL,
		Title:    video.Title,
	}
	resp, err := videoClient.Publish(context.Background(), &req)
	if err != nil {
		return 0, err
	}
	return resp.Uuid, nil
}

// DeleteVideo .
func DeleteVideo(uuid int64) error {
	req := kitex_gen.DeleteRequest{Uuid: uuid}
	_, err := videoClient.Delete(context.Background(), &req)
	return err
}
