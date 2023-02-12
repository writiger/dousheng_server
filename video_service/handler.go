package main

import (
	"context"
	"dousheng_server/video_service/dal/model"
	kitex_gen "dousheng_server/video_service/kitex_gen"
	"dousheng_server/video_service/service"
)

// VideoCenterImpl implements the last service interface defined in the IDL.
type VideoCenterImpl struct{}

// Publish implements the VideoCenterImpl interface.
func (s *VideoCenterImpl) Publish(ctx context.Context, req *kitex_gen.PublishRequest) (*kitex_gen.PublishResponse, error) {
	videoIn := model.Video{
		UserID:   req.UserId,
		PlayURL:  req.PlayUrl,
		CoverURL: req.CoverUrl,
		Title:    req.Title,
	}

	uuid, err := service.VideoCenter{}.Publish(&videoIn)
	if err != nil {
		return &kitex_gen.PublishResponse{
			StatusCode: -1,
			StatusMsg:  "publish action failed",
			Uuid:       0,
		}, err
	}

	return &kitex_gen.PublishResponse{
		StatusCode: 0,
		StatusMsg:  "success",
		Uuid:       uuid,
	}, nil
}

// Delete implements the VideoCenterImpl interface.
func (s *VideoCenterImpl) Delete(ctx context.Context, req *kitex_gen.DeleteRequest) (*kitex_gen.BasicResponse, error) {
	err := service.VideoCenter{}.Delete(req.Uuid)
	if err != nil {
		return &kitex_gen.BasicResponse{
			StatusCode: -1,
			StatusMsg:  err.Error(),
		}, err
	}
	return &kitex_gen.BasicResponse{
		StatusCode: 0,
		StatusMsg:  "success",
	}, nil
}

// Feed implements the VideoCenterImpl interface.
func (s *VideoCenterImpl) Feed(ctx context.Context, req *kitex_gen.FeedRequest) (*kitex_gen.FeedResponse, error) {
	videos, err := service.VideoCenter{}.Feed(req.LastTime)
	if err != nil {
		return &kitex_gen.FeedResponse{Videos: nil}, err
	}
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
	return &kitex_gen.FeedResponse{Videos: videoList}, nil
}
