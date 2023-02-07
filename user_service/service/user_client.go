package service

import (
	"context"
	snowflake_gen "dousheng_server/snowflake_service/kitex_gen"
	"dousheng_server/snowflake_service/kitex_gen/snowflake"
	"dousheng_server/user_service/dal/model"
	"dousheng_server/user_service/dal/query"
	"dousheng_server/user_service/util"
	"errors"
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
)

type UserCenter struct {
}

// CreateUser 注册用户
func (ucs UserCenter) CreateUser(userName, password string) (int64, string, error) {
	// 1. 判断用户名是否存在
	existed, err := query.IsUserNameExisted(userName)
	if err != nil || existed == -1 {
		return -1, "", err
	}

	if existed == 1 {
		return -1, "", errors.New("用户名已存在")
	}
	// 2. 生成用户
	// 获取uuid
	uuid, err := getUUID()

	if err != nil {
		return -1, "", err
	}
	// 加密
	passwd, err := util.HashAndSalt(password)
	if err != nil {
		return -1, "", err
	}
	user := &model.User{
		UUID:          uuid,
		Username:      userName,
		Password:      passwd,
		FollowCount:   0,
		FollowerCount: 0,
	}
	// 3. 保存
	err = query.CreateUser(user)
	if err != nil {
		return -1, "", err
	}
	// 4. 生成Token返回

	return uuid, "token", nil
}

func getUUID() (int64, error) {
	// 1. 通过etcd发现服务
	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"}) // r不应重复使用。
	snowClient, err := snowflake.NewClient("snowflakeservice", client.WithResolver(r))
	if err != nil {
		return -1, err
	}
	var req snowflake_gen.NewIDRequest
	resp, err := snowClient.NewID(context.Background(), &req)
	if err != nil {
		return -1, err
	}
	return resp.ID, nil
}
