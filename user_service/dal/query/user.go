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

// Follow 关注操作
func Follow(userId, followId int64) error {
	err := GormClient.Transaction(func(tx *gorm.DB) error {
		//tx.Create()
		// 1. 新增
		follow := model.Follower{
			UserId:   userId,
			FollowId: followId,
		}
		if err := tx.Create(&follow).Error; err != nil {
			return err
		}

		// followID 的 follower + 1
		if err := tx.Model(&model.User{}).Where("uuid = ?", followId).
			Update("follower_count", gorm.Expr("follower_count + 1")).Error; err != nil {
			return err
		}
		// userId 的 follow + 1
		if err := tx.Model(&model.User{}).Where("uuid = ?", userId).
			Update("follow_count", gorm.Expr("follow_count + 1")).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// Cancel Follow取消关注操作
func CancelFollow(userId, followId int64) error {
	err := GormClient.Transaction(func(tx *gorm.DB) error {
		follow := model.Follower{
			UserId:   userId,
			FollowId: followId,
		}
		if err := tx.Delete(&follow).Error; err != nil {
			return err
		}
		if err := tx.Model(&model.User{}).Where("uuid = ?", followId).
			Update("follower_count", gorm.Expr("follower_count - 1")).Error; err != nil {
			return err
		}
		// userId 的 follow + 1
		if err := tx.Model(&model.User{}).Where("uuid = ?", userId).
			Update("follow_count", gorm.Expr("follow_count - 1")).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// 判断userId是否关注followId
func JudgeFollow(userId, followId int64) (bool, error) {
	follow := model.Follower{
		UserId:   userId,
		FollowId: followId,
	}
	result := GormClient.Find(&follow)
	return result.RowsAffected != 0, result.Error
}

// 通过uuid查询关注列表
func FollowList(uuid int64) (*[]model.Follower, error) {
	followers := []model.Follower{}
	err := GormClient.Where("user_id = ?", uuid).Find(&followers).Error
	return &followers, err
}

// // 通过uuid查询粉丝列表
func FollowerList(uuid int64) (*[]model.Follower, error) {
	followers := []model.Follower{}
	err := GormClient.Where("follow_id = ?", uuid).Find(&followers).Error
	return &followers, err
}

// // 通过uuid查询好友列表
func FriendList(uuid int64) (*[]model.Follower, error) {
	friends := []model.Follower{}
	err := GormClient.Raw("SELECT DISTINCT t1.* FROM (SELECT * FROM followers WHERE `user_id` = ?)  AS t1 INNER JOIN followers t2 ON t1.follow_id = t2.user_id", uuid).Scan(&friends).Error
	return &friends, err
}
