package service

import (
	"dousheng_server/user_service/dal/model"
	"dousheng_server/user_service/dal/query"
	"dousheng_server/user_service/util"
	"dousheng_server/uuidmaker"
	"errors"
)

type UserCenter struct {
}

// CreateUser 注册用户
func (uc UserCenter) CreateUser(userName, password string) error {
	// 1. 判断用户名是否存在
	existed, err := query.IsUserNameExisted(userName)
	if err != nil {
		return err
	}

	if existed == 1 {
		return errors.New("用户名已存在")
	}
	// 2. 生成用户
	// 获取uuid
	uuid, err := uuidmaker.GetUUID()
	if err != nil {
		return err
	}

	// 加密
	passwd, err := util.HashAndSalt(password)
	if err != nil {
		return err
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
		return err
	}
	// 4. 生成Token返回
	return nil
}

// LoginByPassword 使用用户名-密码登录
func (uc UserCenter) LoginByPassword(userName, password string) (*model.User, error) {
	return query.CheckPassword(userName, password)
}

// GetInfo 通过UUID获取用户信息
func (uc UserCenter) GetInfo(uuid int64) (*model.User, error) {
	return query.GetUser(uuid)
}
