package rpc

import (
	"context"
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
	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
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
	// 2. 请求注册服务
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
	// 2. 请求登录服务
	req := kitex_gen.LoginRequest{
		Username: username,
		Password: password,
	}
	resp, err := userClient.Login(context.Background(), &req)
	if err != nil || resp.StatusCode != 0 {
		return 0, err
	}
	return resp.UserId, nil
}
