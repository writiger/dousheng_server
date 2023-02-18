// Package rpc
// 封装与service的请求与返回
package rpc

import (
	"context"
	"dousheng_server/conf"
	"dousheng_server/user_service/dal/model"
	"dousheng_server/user_service/dal/query"
	"dousheng_server/user_service/kitex_gen"
	"dousheng_server/user_service/kitex_gen/usercenter"
	"errors"
	"fmt"
	"github.com/cloudwego/kitex/client"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
	"time"
)

var userClient usercenter.Client

func init() {
	// 通过etcd发现服务
	r, err := etcd.NewEtcdResolver([]string{conf.Conf.EtcdConfig.Url})
	if err != nil {
		panic(err)
	}

	userClient, err = usercenter.NewClient(
		"userservice",
		client.WithResolver(r),
		client.WithSuite(tracing.NewClientSuite()))
	if err != nil {
		panic(err)
	}
}

type FollowerUserInfo struct {
	UUID          int64  `json:"uuid"`
	UserName      string `json:"username"`
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `json:"follower_count"`
	isFollow      bool
}

// Register .
func Register(username, password string) error {
	req := kitex_gen.RegisterRequest{
		Username: username,
		Password: password,
	}
	resp, err := userClient.Register(context.Background(), &req)
	if err != nil {
		return err
	}
	if resp.StatusCode != 0 {
		return errors.New("wrong in userClient.Register()")
	}
	return nil
}

// LoginByPassword .
func LoginByPassword(username, password string) (int64, error) {
	req := kitex_gen.LoginRequest{
		Username: username,
		Password: password,
	}
	resp, err := userClient.Login(context.Background(), &req)
	if err != nil {
		return 0, err
	}
	return resp.UserId, nil
}

// GetUserInfo .
func GetUserInfo(uuid int64) (*model.User, error) {
	req := kitex_gen.GetInfoRequest{Uuid: uuid}
	resp, err := userClient.GetInfo(context.Background(), &req)
	if err != nil {
		return nil, err
	}
	userModel := model.User{
		UUID:          resp.User.Id,
		Username:      resp.User.Name,
		FollowCount:   resp.User.FollowCount,
		FollowerCount: resp.User.FollowerCount,
	}
	return &userModel, nil
}

// Follow 关注
func Follow(userId, followId int64) error {
	//userClient.
	req := kitex_gen.FollowRequest{
		UserId:   userId,
		FollowId: followId,
	}
	_, err := userClient.Follow(context.Background(), &req)
	if err != nil {
		return err
	}

	return nil
}

// CancelFollow 取消关注
func CancelFollow(userId, followId int64) error {
	//userClient.
	req := kitex_gen.FollowRequest{
		UserId:   userId,
		FollowId: followId,
	}
	_, err := userClient.CancelFollow(context.Background(), &req)
	if err != nil {
		return err
	}

	return nil
}

// 关注列表
func FollowList(userId int64) ([]FollowerUserInfo, error) {
	var res []FollowerUserInfo
	req := kitex_gen.GetInfoRequest{Uuid: userId}
	resp, err := userClient.FollowList(context.Background(), &req)
	if err != nil {
		return nil, err
	}
	if resp == nil {
		return nil, err
	}
	for _, item := range resp.Followers {
		userInfo, err := GetUserInfo(item.FollowId)
		if err != nil {
			return nil, err
		}
		//is, err := query.JudgeFollow(userId, userInfo.UUID)//关注列表里的人肯定关注了,所以不用判断了
		if err != nil {
			return nil, err
		}
		res = append(res, FollowerUserInfo{
			UUID:          userInfo.UUID,
			UserName:      userInfo.Username,
			FollowCount:   userInfo.FollowCount,
			FollowerCount: userInfo.FollowerCount,
			isFollow:      true,
		})
	}
	return res, nil
}

// 粉丝列表
func FollowerList(userId int64) ([]FollowerUserInfo, error) {
	var res []FollowerUserInfo
	req := kitex_gen.GetInfoRequest{Uuid: userId}
	resp, err := userClient.FollowerList(context.Background(), &req)
	if err != nil {
		return nil, err
	}
	if resp == nil {
		return nil, err
	}
	for _, item := range resp.Followers {
		userInfo, err := GetUserInfo(item.UserId)
		if err != nil {
			return nil, err
		}

		is, err := query.JudgeFollow(userId, userInfo.UUID)
		if err != nil {
			return nil, err
		}
		res = append(res, FollowerUserInfo{
			UUID:          userInfo.UUID,
			UserName:      userInfo.Username,
			FollowCount:   userInfo.FollowCount,
			FollowerCount: userInfo.FollowerCount,
			isFollow:      is,
		})
	}
	return res, nil
}

// 好友列表
func FriendList(userId int64) ([]FollowerUserInfo, error) {
	var res []FollowerUserInfo
	req := kitex_gen.GetInfoRequest{Uuid: userId}
	resp, err := userClient.FriendList(context.Background(), &req)
	if err != nil {
		return nil, err
	}
	if resp == nil {
		return nil, err
	}
	for _, item := range resp.Followers {
		userInfo, err := GetUserInfo(item.FollowId)
		if err != nil {
			return nil, err
		}
		fmt.Println("takonid:", userId)
		fmt.Println("查出来的id:", userInfo.UUID)
		//如果关系表userInof中的选取的id等于传进来的userId就说明好友应该获取另一个id
		if userId == userInfo.UUID {
			userInfo, err = GetUserInfo(item.UserId)
			fmt.Println("真正查出来的id:", userInfo.UUID)
			if err != nil {
				return nil, err
			}
		}

		is, err := query.JudgeFollow(userId, userInfo.UUID)
		if err != nil {
			return nil, err
		}
		res = append(res, FollowerUserInfo{
			UUID:          userInfo.UUID,
			UserName:      userInfo.Username,
			FollowCount:   userInfo.FollowCount,
			FollowerCount: userInfo.FollowerCount,
			isFollow:      is,
		})
	}
	return res, nil
}

// 发送消息
func SendMessage(fromUserId, ToUserId int64, message string) error {
	req := kitex_gen.SendMessageRequest{
		UserId:  fromUserId,
		ToId:    ToUserId,
		Message: message,
	}
	_, err := userClient.SendMessage(context.Background(), &req)
	if err != nil {
		return err
	}
	return nil
}

// 消息列表
func MessageList(fromUserId, ToUserId, lastTime int64) ([]model.Message, error) {
	var messageList []model.Message
	req := kitex_gen.MessageListRequest{
		UserId:   fromUserId,
		ToId:     ToUserId,
		LastTime: lastTime,
	}
	resp, err := userClient.MessageList(context.Background(), &req)
	if err != nil {
		return nil, err
	}
	if resp == nil {
		return nil, err
	}
	for _, item := range resp.MessageList {
		createAt := time.UnixMilli(item.CreateTime)
		messageList = append(messageList, model.Message{
			Id:         item.Id,
			Messages:   item.Content,
			FromUserId: item.FromUserId,
			ToUserId:   item.ToUserId,
			CreatedAt:  createAt,
		})
	}

	return messageList, err
}
