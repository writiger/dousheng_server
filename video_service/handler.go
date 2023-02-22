package main

import (
	"context"
	kitex_gen "dousheng_server/video_service/kitex_gen"
	"dousheng_server/video_service/service"
	"fmt"
)

// VideoCenterImpl implements the last service interface defined in the IDL.
type VideoCenterImpl struct{}

// Publish implements the VideoCenterImpl interface.
func (s *VideoCenterImpl) Publish(ctx context.Context, req *kitex_gen.PublishRequest) (*kitex_gen.PublishResponse, error) {
	uuid, err := service.VideoCenter{}.Publish(req)
	if err != nil {
		return nil, err
	}
	return &kitex_gen.PublishResponse{
		StatusCode: 0,
		StatusMsg:  "success",
		Uuid:       uuid,
	}, nil
}

// Delete implements the VideoCenterImpl interface.
func (s *VideoCenterImpl) Delete(ctx context.Context, req *kitex_gen.DeleteRequest) (*kitex_gen.BasicResponse, error) {
	err := service.VideoCenter{}.Delete(req.VideoId)
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

// PostComment implements the VideoCenterImpl interface.
func (s *VideoCenterImpl) PostComment(ctx context.Context, req *kitex_gen.PostCommentRequest) (*kitex_gen.PostCommentResponse, error) {
	comment, err := service.VideoCenter{}.PostComment(req)
	if err != nil {
		return nil, err
	}
	return &kitex_gen.PostCommentResponse{Comment: comment}, nil
}

// DeleteComment implements the VideoCenterImpl interface.
func (s *VideoCenterImpl) DeleteComment(ctx context.Context, req *kitex_gen.DeleteCommentRequest) (*kitex_gen.BasicResponse, error) {
	err := service.VideoCenter{}.DeleteComment(req.Uuid)
	if err != nil {
		return &kitex_gen.BasicResponse{
			StatusCode: -1,
			StatusMsg:  "delete comment wrong",
		}, err
	}
	return &kitex_gen.BasicResponse{
		StatusCode: 0,
		StatusMsg:  "success",
	}, nil
}

// GetComment implements the VideoCenterImpl interface.
func (s *VideoCenterImpl) GetComment(ctx context.Context, req *kitex_gen.GetCommentRequest) (*kitex_gen.GetCommentResponse, error) {
	comments, err := service.VideoCenter{}.GetComment(req.VideoId)
	if err != nil {
		return nil, err
	}
	return &kitex_gen.GetCommentResponse{Comments: comments}, nil
}

// IsFavorite implements the VideoCenterImpl interface.
func (s *VideoCenterImpl) IsFavorite(ctx context.Context, req *kitex_gen.IsFavoriteRequest) (*kitex_gen.BasicResponse, error) {
	res := service.VideoCenter{}.IsFavorite(req.UserId, req.VideoId)
	return &kitex_gen.BasicResponse{StatusMsg: fmt.Sprint(res), StatusCode: 0}, nil
}
