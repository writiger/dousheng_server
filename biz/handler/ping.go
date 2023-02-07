// Code generated by hertz generator.

package handler

import (
	"context"
	"dousheng_server/user_service/kitex_gen"
	"dousheng_server/user_service/kitex_gen/usercenter"
	"fmt"
	"github.com/cloudwego/kitex/client"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"

	"github.com/cloudwego/hertz/pkg/app"
)

// Ping .
func Ping(ctx context.Context, c *app.RequestContext) {
	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	if err != nil {
		panic(err)
	}

	cli, err := usercenter.NewClient(
		"userservice",
		client.WithResolver(r),
		client.WithSuite(tracing.NewClientSuite()))
	if err != nil {
		return
	}

	req := kitex_gen.Request{
		Ping: "i'm hertz",
	}
	resp, err := cli.Ping(context.Background(), &req)
	if err != nil {
		return
	}
	fmt.Println(resp)
}