package query

import (
	"dousheng_server/user_service/dal/model"
	"dousheng_server/user_service/util"
	"errors"
	"gorm.io/gorm"
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

// GetUser 通过UUID获取用户
func GetUser(uuid int64) (*model.User, error) {
	user := model.User{}
	err := GormClient.Where("uuid = ?", uuid).
		Select("uuid", "username", "follow_count", "follower_count").
		Find(&user).Error
	return &user, err
}

// Follow 关注
func Follow(userId, followId int64) error {
	err := GormClient.Transaction(func(tx *gorm.DB) error {
		//tx.Create()
		// 1. 新增

		// followID 的 follower + 1
		// 2. 增加

		// userId 的 follow + 1
		//
		return nil
	})
	return err
}
