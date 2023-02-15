package main

import (
	"dousheng_server/conf"
	"dousheng_server/user_service/kitex_gen/usercenter"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	"net"
)

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", ":8900")
	r, err := etcd.NewEtcdRegistry([]string{conf.Conf.EtcdConfig.Url})
	if err != nil {
		log.Fatal(err)
	}
	svr := usercenter.NewServer(new(UserCenterImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "userservice"}),
		server.WithRegistry(r),
		server.WithServiceAddr(addr))

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
