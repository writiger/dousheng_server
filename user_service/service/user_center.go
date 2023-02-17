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

// Follow 取消关注用户
func (uc UserCenter) CancelFollow(userId, followId int64) error {
	return query.CancelFollow(userId, followId)
}

// Follow 判断用户是否关注
func (uc UserCenter) JudgeFollow(userId, followId int64) (bool, error) {
	return query.JudgeFollow(userId, followId)
}

// GetInfo 通过UUID获取关注列表
func (uc UserCenter) FollowList(uuid int64) ([]*kitex_gen.Follower, error) {
	followers, err := query.FollowList(uuid)
	return modelToKitexFollower(followers), err
}

// GetInfo 通过UUID获取粉丝列表
func (uc UserCenter) FollowerList(uuid int64) ([]*kitex_gen.Follower, error) {
	followers, err := query.FollowerList(uuid)
	return modelToKitexFollower(followers), err
}

// GetInfo 通过UUID获取好友列表
func (uc UserCenter) FriendList(uuid int64) ([]*kitex_gen.Follower, error) {
	followers, err := query.FriendList(uuid)
	return modelToKitexFollower(followers), err
}

// 发消息
func (uc UserCenter) SendMessage(FromUserId, ToUserId int64, message string) error {
	id, err := uuidmaker.GetUUID()
	if err != nil {
		return err
	}
	messageModel := &model.Message{
		Id:         id,
		Messages:   message,
		FromUserId: FromUserId,
		ToUserId:   ToUserId,
	}
	err = query.SendMessage(messageModel)
	if err != nil {
		return err
	}
	return nil
}

// 获取消息列表
func (uc UserCenter) MessageList(FromUserId, ToUserId int64) ([]*kitex_gen.Message, error) {
	messageList, err := query.MessageList(FromUserId, ToUserId)
	return modelToKitexMessage(messageList), err
}

// 将model中的follower转换为kitex中生成的follwer
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

// 将model中的message转换为kitex中生成的message
func modelToKitexMessage(messages *[]model.Message) []*kitex_gen.Message {
	var messageList []*kitex_gen.Message
	for _, item := range *messages {
		createTime := item.CreatedAt.UnixMilli()
		messageList = append(messageList, &kitex_gen.Message{
			Id:         item.Id,
			FromUserId: item.FromUserId,
			ToUserId:   item.ToUserId,
			Message:    item.Messages,
			CreateTime: createTime,
		})
	}
	return messageList
}
