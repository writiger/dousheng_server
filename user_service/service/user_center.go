package service

import (
	"dousheng_server/user_service/dal/model"
	"dousheng_server/user_service/dal/query"
	"dousheng_server/user_service/kitex_gen"
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

// Follow 关注用户
func (uc UserCenter) Follow(userId, followId int64) error {
	return query.Follow(userId, followId)
}

// CancelFollow 取消关注用户
func (uc UserCenter) CancelFollow(userId, followId int64) error {
	return query.CancelFollow(userId, followId)
}

// JudgeFollow 判断用户是否关注
func (uc UserCenter) JudgeFollow(userId, followId int64) (bool, error) {
	return query.JudgeFollow(userId, followId)
}

// FollowList 通过UUID获取关注列表
func (uc UserCenter) FollowList(uuid int64) ([]*kitex_gen.Follower, error) {
	followers, err := query.FollowList(uuid)
	return modelToKitexFollower(followers), err
}

// FollowerList 通过UUID获取粉丝列表
func (uc UserCenter) FollowerList(uuid int64) ([]*kitex_gen.Follower, error) {
	followers, err := query.FollowerList(uuid)
	return modelToKitexFollower(followers), err
}

// FriendList  通过UUID获取好友列表
func (uc UserCenter) FriendList(uuid int64) ([]*kitex_gen.Follower, error) {
	followers, err := query.FriendList(uuid)
	return modelToKitexFollower(followers), err
}

// SendMessages 发消息
func SendMessages(FromUser, ToUser int64, message string) error {
	id, err := uuidmaker.GetUUID()
	if err != nil {
		return err
	}
	messageModel := &model.Message{
		Id:         id,
		Messages:   message,
		FromUserId: FromUser,
		ToUserId:   ToUser,
	}
	err = query.SendMessage(messageModel)
	if err != nil {
		return err
	}
	return nil
}

// 将model中的follower转换为kitex中生成的follower
func modelToKitexFollower(follower *[]model.Follower) []*kitex_gen.Follower {
	var followerList []*kitex_gen.Follower
	for _, item := range *follower {
		followerList = append(followerList, &kitex_gen.Follower{
			UserId:   item.UserId,
			FollowId: item.FollowId,
		})
	}
	return followerList
}
