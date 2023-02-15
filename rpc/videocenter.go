package rpc

import (
	"context"
	"dousheng_server/conf"
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
	r, err := etcd.NewEtcdResolver([]string{conf.Conf.EtcdConfig.Url})
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

// CommentWithUser 将用户信息中的user_id替换为user
type CommentWithUser struct {
	UUID       int64          `json:"comment_id"`
	UserInfo   usermodel.User `json:"user"`
	Content    string         `json:"content"`
	CreateDate string         `json:"createDate"`
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
	req := kitex_gen.DeleteRequest{VideoId: uuid}
	_, err := videoClient.Delete(context.Background(), &req)
	return err
}

// Feed  返回视频列表和最早时间戳
func Feed(timeStamp, userId int64) ([]VideoWithUser, int64, error) {
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
		// 判断是否点过赞
		isFavorite, err := IsFavorite(userId, item.Uuid)
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
			IsFavorite:    isFavorite,
		})
	}
	return res, timestamp, nil
}

// VideoList .
func VideoList(userId int64) ([]VideoWithUser, error) {
	var res []VideoWithUser
	req := kitex_gen.VideoListRequest{UserId: userId}
	resp, err := videoClient.VideoList(context.Background(), &req)
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

// LikeVideo .
func LikeVideo(userId, videoId int64, actionType int32) error {
	req := kitex_gen.LikeRequest{
		UserId:     userId,
		VideoId:    videoId,
		ActionType: actionType,
	}
	_, err := videoClient.Like(context.Background(), &req)
	return err
}

// GetVideo 使用uuid获取视频
func GetVideo(videoId int64) (*VideoWithUser, error) {
	req := kitex_gen.GetVideoRequest{Uuid: videoId}
	resp, err := videoClient.GetVideo(context.Background(), &req)
	if err != nil {
		return nil, err
	}
	// 替换用户信息
	userInfo, err := GetUserInfo(resp.Video.UserId)
	res := &VideoWithUser{
		UUID:          resp.Video.Uuid,
		UserInfo:      *userInfo,
		PlayURL:       resp.Video.PlayUrl,
		CoverURL:      resp.Video.CoverUrl,
		FavoriteCount: resp.Video.FavoriteCount,
		CommentCount:  resp.Video.CommentCount,
		Title:         resp.Video.Title,

		// TODO IsFavorite

		IsFavorite: false,
	}
	return res, nil
}

// FavoriteVideoList .
func FavoriteVideoList(userId int64) ([]VideoWithUser, error) {
	var res []VideoWithUser
	req := kitex_gen.GetVideoRequest{Uuid: userId}
	resp, err := videoClient.GetFavoriteVideo(context.Background(), &req)
	if err != nil {
		return nil, err
	}
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
			IsFavorite:    true,
		})
	}
	return res, nil
}

// IsFavorite .
func IsFavorite(userId, videoId int64) (bool, error) {
	req := kitex_gen.IsFavoriteRequest{
		UserId:  userId,
		VideoId: videoId,
	}
	resp, err := videoClient.IsFavorite(context.Background(), &req)
	// 没研究出来为什么
	// 使用返回值中有false
	// 返回的resp就是nil
	if resp.StatusMsg == "true" {
		return true, err
	}
	return false, err
}

// PostComment .
func PostComment(comment model.Comment) (*CommentWithUser, error) {
	req := kitex_gen.PostCommentRequest{
		UserId:  comment.UserId,
		VideoId: comment.VideoId,
		Content: comment.Content,
	}
	resp, err := videoClient.PostComment(context.Background(), &req)
	if err != nil {
		return nil, err
	}
	// 替换用户信息
	userInfo, err := GetUserInfo(resp.Comment.UserId)
	// 转换参数
	return &CommentWithUser{
		UUID:       resp.Comment.Uuid,
		UserInfo:   *userInfo,
		Content:    resp.Comment.Content,
		CreateDate: resp.Comment.CreateDate,
	}, nil
}

// DeleteComment .
func DeleteComment(uuid int64) error {
	req := kitex_gen.DeleteCommentRequest{Uuid: uuid}
	_, err := videoClient.DeleteComment(context.Background(), &req)
	return err
}

// GetComment .
func GetComment(uuid int64) ([]CommentWithUser, error) {
	var res []CommentWithUser
	req := kitex_gen.GetCommentRequest{VideoId: uuid}
	resp, err := videoClient.GetComment(context.Background(), &req)
	if err != nil {
		return nil, err
	}
	for _, item := range resp.Comments {
		userInfo, err := GetUserInfo(item.UserId)
		if err != nil {
			return nil, err
		}
		res = append(res, CommentWithUser{
			UUID:       item.Uuid,
			UserInfo:   *userInfo,
			Content:    item.Content,
			CreateDate: item.CreateDate,
		})
	}
	return res, nil
}
