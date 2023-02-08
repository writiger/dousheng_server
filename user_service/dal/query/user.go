package query

import (
	"dousheng_server/user_service/dal/model"
	"dousheng_server/user_service/util"
	"errors"
)

// CreateUser 添加用户
func CreateUser(user *model.User) error {
	return GormClient.Create(user).Error
}

// IsUserNameExisted 用户是否存在
// 0  - 不存在
// 1  - 存在
// -1 - 系统错误
func IsUserNameExisted(name string) (int64, error) {
	var users []model.User
	res := GormClient.Where("username = ?", name).Find(&users)
	err := res.Error
	if err != nil {
		return -1, err
	}
	return res.RowsAffected, nil
}

// CheckPassword 判断用户名密码是否正确
func CheckPassword(name, password string) (*model.User, error) {
	var user model.User
	err := GormClient.Where("username = ?", name).Find(&user).Error
	if !util.ComparePasswords(user.Password, password) {
		return nil, errors.New("wrong password")
	}
	return &user, err
}
