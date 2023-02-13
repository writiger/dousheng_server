package rpc

import (
	"context"
	usermodel "dousheng_server/user_service/dal/model"
	"dousheng_server/video_service/dal/model"
	"dousheng_server/video_service/kitex_gen"
	"dousheng_server/video_service/kitex_gen/videocenter"
	"github.com/cloudwego/kitex/client"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
	"time"
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

// VideoWithUser 将用户信息中的user_id替换为user
type VideoWithUser struct {
	UUID          int64          `json:"id"`
	UserInfo      usermodel.User `json:"author"`
	PlayURL       string         `json:"play_url"`
	CoverURL      string         `json:"cover_url"`
	FavoriteCount int64          `json:"favorite_count"`
	CommentCount  int64          `json:"comment_count"`
	Title         string         `json:"title"`
	IsFavorite    bool           `json:"is_favorite"`
	CreatedAt     time.Time
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

// Feed  返回视频列表和最早时间戳
func Feed(timeStamp int64) ([]VideoWithUser, int64, error) {
	// TODO 判断是否喜欢

	var res []VideoWithUser
	timestamp := time.Now().UnixMilli()

	req := kitex_gen.FeedRequest{
		LastTime: timeStamp,
	}
	resp, err := videoClient.Feed(context.Background(), &req)
	if err != nil || resp == nil {
		return nil, 0, err
	}
	// 修饰返回值
	for _, item := range resp.Videos {
		if item.CreateTime < timestamp {
			timestamp = item.CreateTime
		}
		// 替换用户信息
		userInfo, err := GetUserInfo(item.UserId)
		if err != nil {
			return nil, 0, err
		}
		res = append(res, VideoWithUser{
			UUID:          item.Uuid,
			UserInfo:      *userInfo,
			PlayURL:       item.PlayUrl,
			CoverURL:      item.CoverUrl,
			FavoriteCount: item.FavoriteCount,
			CommentCount:  item.CommentCount,
			Title:         item.Title,
		})
	}
	return res, timestamp, nil
}

// VideoList .
func VideoList(userId int64) ([]VideoWithUser, error) {
	var res []VideoWithUser
	req := kitex_gen.ListRequest{UserId: userId}
	resp, err := videoClient.List(context.Background(), &req)
	if err != nil || resp == nil {
		return nil, err
	}
	// 修饰返回值
	for _, item := range resp.Videos {
		// 替换用户信息
		userInfo, err := GetUserInfo(item.UserId)
		if err != nil {
			return nil, err
		}
		res = append(res, VideoWithUser{
			UUID:          item.Uuid,
			UserInfo:      *userInfo,
			PlayURL:       item.PlayUrl,
			CoverURL:      item.CoverUrl,
			FavoriteCount: item.FavoriteCount,
			CommentCount:  item.CommentCount,
			Title:         item.Title,
		})
	}
	return res, nil
}
