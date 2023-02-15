package main

import (
	"dousheng_server/conf"
	"dousheng_server/video_service/kitex_gen/videocenter"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	"net"
)

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", ":8902")
	r, err := etcd.NewEtcdRegistry([]string{conf.Conf.EtcdConfig.Url})
	if err != nil {
		log.Fatal(err)
	}
	svr := videocenter.NewServer(new(VideoCenterImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "videoservice"}),
		server.WithRegistry(r),
		server.WithServiceAddr(addr))

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
