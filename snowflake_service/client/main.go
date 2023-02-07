package main

import (
	"context"
	"dousheng_server/snowflake_service/kitex_gen"
	"dousheng_server/snowflake_service/kitex_gen/snowflake"
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	"time"
)

func main() {
	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"}) // r不应重复使用。
	snowClient, err := snowflake.NewClient("snowflakeservice", client.WithResolver(r))
	if err != nil {
		log.Fatal(err)
	}
	for {
		req := kitex_gen.NewIDRequest{}
		resp, err := snowClient.NewID(context.Background(), &req)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Get NewID:", resp.ID)
		time.Sleep(time.Second)
	}
}
