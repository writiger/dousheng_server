// Package rpc
// 封装与service的请求与返回
package rpc

import (
	"context"
	"dousheng_server/conf"
	"dousheng_server/user_service/dal/model"
	"dousheng_server/user_service/kitex_gen"
	"dousheng_server/user_service/kitex_gen/usercenter"
	"errors"
	"github.com/cloudwego/kitex/client"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var userClient usercenter.Client

func init() {
	// 通过etcd发现服务
	r, err := etcd.NewEtcdResolver([]string{conf.Conf.EtcdConfig.Url})
	if err != nil {
		panic(err)
	}

	userClient, err = usercenter.NewClient(
		"userservice",
		client.WithResolver(r),
		client.WithSuite(tracing.NewClientSuite()))
	if err != nil {
		panic(err)
	}
}

// Register .
func Register(username, password string) error {
	req := kitex_gen.RegisterRequest{
		Username: username,
		Password: password,
	}
	resp, err := userClient.Register(context.Background(), &req)
	if err != nil {
		return err
	}
	if resp.StatusCode != 0 {
		return errors.New("wrong in userClient.Register()")
	}
	return nil
}

// LoginByPassword .
func LoginByPassword(username, password string) (int64, error) {
	req := kitex_gen.LoginRequest{
		Username: username,
		Password: password,
	}
	resp, err := userClient.Login(context.Background(), &req)
	if err != nil {
		return 0, err
	}
	return resp.UserId, nil
}

// GetUserInfo .
func GetUserInfo(uuid int64) (*model.User, error) {
	req := kitex_gen.GetInfoRequest{Uuid: uuid}
	resp, err := userClient.GetInfo(context.Background(), &req)
	if err != nil {
		return nil, err
	}
	userModel := model.User{
		UUID:          resp.User.Id,
		Username:      resp.User.Name,
		FollowCount:   resp.User.FollowCount,
		FollowerCount: resp.User.FollowerCount,
	}
	return &userModel, nil
}
