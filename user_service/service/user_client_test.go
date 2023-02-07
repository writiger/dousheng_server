package service

import (
	"context"
	"dousheng_server/user_service/kitex_gen/kitex_gen"
	"dousheng_server/user_service/kitex_gen/kitex_gen/usercenter"
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	"testing"
)

func TestUserCenter_CreateUser(t *testing.T) {
	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"}) // r不应重复使用。
	userClient, err := usercenter.NewClient("userservice", client.WithResolver(r))
	if err != nil {
		log.Fatal(err)
	}
	req := kitex_gen.RegisterRequest{
		Username: "hello",
		Password: "test",
	}
	_, err = userClient.Register(context.Background(), &req)
	if err != nil {
		return
	}
}
