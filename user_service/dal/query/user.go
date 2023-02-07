package query

import (
	"dousheng_server/snowflake_service/kitex_gen/kitex_gen/snowflake"
	"dousheng_server/user_service/dal/model"
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
)

func CreateUser(user *model.User) error {
	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"}) // r不应重复使用。
	snowClient, err := snowflake.NewClient("snowflakeservice", client.WithResolver(r))
	if err != nil {
		log.Fatal(err)
	}

	if err := GormClient.Create(user).Error; err != nil {
		return err
	}
	return nil
}
