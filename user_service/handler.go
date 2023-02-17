package main

import (
	"context"
	"dousheng_server/user_service/kitex_gen"
	"dousheng_server/user_service/service"
	"fmt"
)

// UserCenterImpl implements the last service interface defined in the IDL.
type UserCenterImpl struct{}

// Ping implements the UserCenterImpl interface.
func (s *UserCenterImpl) Ping(ctx context.Context, req *kitex_gen.Request) (resp *kitex_gen.Response, err error) {
	resp = new(kitex_gen.Response)
	fmt.Println("Get Ping Message:", req.Ping)
	resp.Pong = "Hello Client"
	return
}

// Register implements the UserCenterImpl interface.
func (s *UserCenterImpl) Register(ctx context.Context, req *kitex_gen.RegisterRequest) (*kitex_gen.BasicResponse, error) {
	err := service.UserCenter{}.CreateUser(req.Username, req.Password)
	if err != nil {
		return &kitex_gen.BasicResponse{
			StatusCode: -1,
			StatusMsg:  "user create failed",
		}, err
	}
	return &kitex_gen.BasicResponse{
		StatusCode: 0,
		StatusMsg:  "success",
	}, nil
}

// Login implements the UserCenterImpl interface.
func (s *UserCenterImpl) Login(ctx context.Context, req *kitex_gen.LoginRequest) (resp *kitex_gen.LoginResponse, err error) {
	userModel, err := service.UserCenter{}.LoginByPassword(req.Username, req.Password)

	if err != nil {
		return &kitex_gen.LoginResponse{
			StatusCode: -1,
			StatusMsg:  "UserCenter{}.LoginByPassword wrong",
		}, err
	}
	if userModel.UUID == 0 {
		return &kitex_gen.LoginResponse{
			StatusCode: -1,
			StatusMsg:  "failed",
		}, nil
	}
	return &kitex_gen.LoginResponse{
		StatusCode: 0,
		StatusMsg:  "success",
		UserId:     userModel.UUID,
	}, nil
}

// GetInfo implements the UserCenterImpl interface.
func (s *UserCenterImpl) GetInfo(ctx context.Context, req *kitex_gen.GetInfoRequest) (*kitex_gen.GetInfoResponse, error) {
	userModel, err := service.UserCenter{}.GetInfo(req.Uuid)
	if err != nil {
		return &kitex_gen.GetInfoResponse{
			StatusCode: -1,
			StatusMsg:  "get user failed",
		}, err
	}
	user := &kitex_gen.User{
		Id:            userModel.UUID,
		Name:          userModel.Username,
		FollowCount:   userModel.FollowCount,
		FollowerCount: userModel.FollowerCount,

		// TODO 判断是否关注

		IsFollow: false,
	}
	return &kitex_gen.GetInfoResponse{
		StatusCode: 0,
		StatusMsg:  "success",
		User:       user,
	}, nil
}

// Follow implements the UserCenterImpl interface.
func (s *UserCenterImpl) Follow(ctx context.Context, req *kitex_gen.FollowRequest) (resp *kitex_gen.BasicResponse, err error) {
	// TODO: Your code here...
	err = service.UserCenter{}.Follow(req.UserId, req.FollowId)
	if err != nil {
		resp = &kitex_gen.BasicResponse{
			StatusCode: 1,
			StatusMsg:  "关注失败",
		}
		return
	}
	resp = &kitex_gen.BasicResponse{
		StatusCode: 0,
		StatusMsg:  "关注成功",
	}
	return
}

// CancelFollow implements the UserCenterImpl interface.
func (s *UserCenterImpl) CancelFollow(ctx context.Context, req *kitex_gen.FollowRequest) (resp *kitex_gen.BasicResponse, err error) {
	// TODO: Your code here...
	err = service.UserCenter{}.CancelFollow(req.UserId, req.FollowId)
	if err != nil {
		resp = &kitex_gen.BasicResponse{
			StatusCode: 1,
			StatusMsg:  "取消关注失败",
		}
		return
	}
	resp = &kitex_gen.BasicResponse{
		StatusCode: 0,
		StatusMsg:  "取消关注成功",
	}
	return
}

// FollowList implements the UserCenterImpl interface.
func (s *UserCenterImpl) FollowList(ctx context.Context, req *kitex_gen.GetInfoRequest) (*kitex_gen.FollowListResponse, error) {
	// TODO: Your code here...
	resp, err := service.UserCenter{}.FollowList(req.Uuid)
	return &kitex_gen.FollowListResponse{Followers: resp}, err
}

// FollowerList implements the UserCenterImpl interface.
func (s *UserCenterImpl) FollowerList(ctx context.Context, req *kitex_gen.GetInfoRequest) (*kitex_gen.FollowListResponse, error) {
	// TODO: Your code here...
	resp, err := service.UserCenter{}.FollowerList(req.Uuid)
	return &kitex_gen.FollowListResponse{Followers: resp}, err

}

// FriendList implements the UserCenterImpl interface.
func (s *UserCenterImpl) FriendList(ctx context.Context, req *kitex_gen.GetInfoRequest) (*kitex_gen.FollowListResponse, error) {
	// TODO: Your code here...
	resp, err := service.UserCenter{}.FriendList(req.Uuid)
	return &kitex_gen.FollowListResponse{Followers: resp}, err

}

// JudgeFollow implements the UserCenterImpl interface.
func (s *UserCenterImpl) JudgeFollow(ctx context.Context, req *kitex_gen.FollowRequest) (resp *kitex_gen.JudgeFollowResponse, err error) {
	// TODO: Your code here...
	var is bool
	is, err = service.UserCenter{}.JudgeFollow(req.UserId, req.FollowId)
	if err != nil {
		resp = &kitex_gen.JudgeFollowResponse{
			StatusCode: 1,
			StatusMsg:  "关系查询失败",
			Is:         is,
		}
		return
	}
	resp = &kitex_gen.JudgeFollowResponse{
		StatusCode: 0,
		StatusMsg:  "关系查询成功",
		Is:         is,
	}
	return
}
