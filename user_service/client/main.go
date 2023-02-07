package main

import (
	"context"
	"dousheng_server/user_service/kitex_gen"
	"dousheng_server/user_service/kitex_gen/usercenter"
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	"time"
)

func main() {
	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"}) // r不应重复使用。
	userClient, err := usercenter.NewClient("userservice", client.WithResolver(r))
	if err != nil {
		log.Fatal(err)
	}
	for {
		req := kitex_gen.Request{Ping: "Hello server"}
		resp, err := userClient.Ping(context.Background(), &req)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Get Pong Message:", resp.Pong)
		time.Sleep(time.Second)
	}

}
