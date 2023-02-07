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
	// TODO: Your code here...
	resp = new(kitex_gen.Response)
	fmt.Println("Get Ping Message:", req.Ping)
	resp.Pong = "Hello Client"
	return
}

// Register implements the UserCenterImpl interface.
func (s *UserCenterImpl) Register(ctx context.Context, req *kitex_gen.RegisterRequest) (*kitex_gen.RegisterResponse, error) {
	// TODO: Your code here...
	userID, token, err := service.UserCenter{}.CreateUser(req.Username, req.Password)
	if err != nil {
		return &kitex_gen.RegisterResponse{
			StatusCode: -1,
			StatusMsg:  "user create failed",
			UserId:     -1,
			Token:      "",
		}, err
	}
	return &kitex_gen.RegisterResponse{
		StatusCode: 0,
		StatusMsg:  "success",
		UserId:     userID,
		Token:      token,
	}, nil
}
