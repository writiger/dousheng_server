package conf

import (
	"gopkg.in/ini.v1"
)

var Conf DConfig

type DConfig struct {
	EtcdConfig   `ini:"etcd"`
	MysqlConfig  `ini:"mysql"`
	StaticConfig `ini:"static"`
}

type EtcdConfig struct {
	Url string `ini:"url"`
}

type MysqlConfig struct {
	Url      string `ini:"url"`
	User     string `ini:"user"`
	Password string `ini:"password"`
	Database string `ini:"database"`
}

type StaticConfig struct {
	SaverIp string `ini:"saverIp"`
	Reflect string `ini:"reflect"`
}

func init() {
	// 先实例化结构体，将指针传入MapTo方法中
	err := ini.MapTo(&Conf, "./conf.ini")
	if err != nil {
		// 服务使用的相对路径
		err = ini.MapTo(&Conf, "../../conf.ini")
	}
	if err != nil {
		panic("failed when init conf err : " + err.Error())
	}
}
