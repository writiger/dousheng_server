package uuidmaker

import (
	"context"
	snowflake_gen "dousheng_server/snowflake_service/kitex_gen"
	"dousheng_server/snowflake_service/kitex_gen/snowflake"
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var snowClient snowflake.Client

func init() {
	// 1. 通过etcd发现服务
	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"}) // r不应重复使用。
	snowClient, err = snowflake.NewClient("snowflakeservice", client.WithResolver(r))
	if err != nil {
		panic("wrong on snowflake_service")
	}
}

func GetUUID() (int64, error) {
	var req snowflake_gen.NewIDRequest
	resp, err := snowClient.NewID(context.Background(), &req)
	if err != nil {
		return -1, err
	}
	return resp.ID, nil
}
