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
		return nil, err
	}
	return &kitex_gen.FeedResponse{Videos: videos}, nil
}

// VideoList implements the VideoCenterImpl interface.
func (s *VideoCenterImpl) VideoList(ctx context.Context, req *kitex_gen.VideoListRequest) (resp *kitex_gen.VideoListResponse, err error) {
	videos, err := service.VideoCenter{}.VideoList(req.UserId)
	if err != nil {
		return nil, err
	}
	return &kitex_gen.VideoListResponse{Videos: videos}, nil
}

// Like implements the VideoCenterImpl interface.
func (s *VideoCenterImpl) Like(ctx context.Context, req *kitex_gen.LikeRequest) (*kitex_gen.BasicResponse, error) {
	err := service.VideoCenter{}.Like(req.UserId, req.VideoId, req.ActionType)
	if err != nil {
		return nil, err
	}
	return &kitex_gen.BasicResponse{
		StatusCode: 0,
		StatusMsg:  "success",
	}, nil
}

// GetVideo implements the VideoCenterImpl interface.
func (s *VideoCenterImpl) GetVideo(ctx context.Context, req *kitex_gen.GetVideoRequest) (resp *kitex_gen.GetVideoResponse, err error) {
	video, err := service.VideoCenter{}.GetVideo(req.Uuid)
	if err != nil {
		return nil, err
	}
	return &kitex_gen.GetVideoResponse{Video: video}, nil
}

// GetFavoriteVideo implements the VideoCenterImpl interface.
func (s *VideoCenterImpl) GetFavoriteVideo(ctx context.Context, req *kitex_gen.GetVideoRequest) (*kitex_gen.GetFavoriteVideosResponse, error) {
	videos, err := service.VideoCenter{}.FavoriteList(req.Uuid)
	if err != nil {
		return nil, err
	}
	return &kitex_gen.GetFavoriteVideosResponse{Videos: videos}, nil
}
