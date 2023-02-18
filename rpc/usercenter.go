// Package rpc
// 封装与service的请求与返回
package rpc

import (
	"context"
	"dousheng_server/conf"
	"dousheng_server/user_service/dal/query"
	"dousheng_server/user_service/kitex_gen"
	"dousheng_server/user_service/kitex_gen/usercenter"
	"errors"
	"github.com/cloudwego/kitex/client"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
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

type UserInfo struct {
	UUID          int64  `gorm:"primaryKey" json:"id"`
	Username      string `json:"name"`
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
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
func GetUserInfo(tokenId, userId int64) (*UserInfo, error) {
	req := kitex_gen.GetInfoRequest{Uuid: userId}
	resp, err := userClient.GetInfo(context.Background(), &req)
	if err != nil {
		return nil, err
	}
	isFollowed := false
	// 没登陆的用户
	if tokenId != 0 {
		isFollowed, err = IsFollowed(tokenId, userId)
		if err != nil {
			return nil, err
		}
	}
	userInfo := UserInfo{
		UUID:          resp.User.Id,
		Username:      resp.User.Name,
		FollowCount:   resp.User.FollowCount,
		FollowerCount: resp.User.FollowerCount,
		IsFollow:      isFollowed,
	}
	return &userInfo, nil
}

// Follow 关注
func Follow(userId, followId int64) error {
	req := kitex_gen.FollowRequest{
		UserId:   userId,
		FollowId: followId,
	}
	_, err := userClient.Follow(context.Background(), &req)
	return err
}

// CancelFollow 取消关注
func CancelFollow(userId, followId int64) error {
	req := kitex_gen.FollowRequest{
		UserId:   userId,
		FollowId: followId,
	}
	_, err := userClient.CancelFollow(context.Background(), &req)
	return err
}

// FollowList 关注列表
func FollowList(userId int64) ([]FollowerUserInfo, error) {
	var res []FollowerUserInfo
	req := kitex_gen.GetInfoRequest{Uuid: userId}
	resp, err := userClient.FollowList(context.Background(), &req)
	if err != nil {
		return nil, err
	}
	for _, item := range resp.Followers {
		userInfo, err := GetUserInfo(userId, item.FollowId)
		if err != nil {
			return nil, err
		}
		//is, err := query.JudgeFollow(userId, userInfo.UUID)
		//关注列表里的人肯定关注了,所以不用判断了
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

// FollowerList 粉丝列表
func FollowerList(userId int64) ([]FollowerUserInfo, error) {
	var res []FollowerUserInfo
	req := kitex_gen.GetInfoRequest{Uuid: userId}
	resp, err := userClient.FollowerList(context.Background(), &req)
	if err != nil {
		return nil, err
	}
	for _, item := range resp.Followers {
		userInfo, err := GetUserInfo(userId, item.UserId)
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

// FriendList 好友列表
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
		userInfo, err := GetUserInfo(userId, item.FollowId)
		if err != nil {
			return nil, err
		}
		//如果关系表userInfo中的选取的id等于传进来的userId就说明好友应该获取另一个id
		if userId == userInfo.UUID {
			userInfo, err = GetUserInfo(userId, item.UserId)
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

// IsFollowed 是否关注
func IsFollowed(userId, followId int64) (bool, error) {
	req := kitex_gen.FollowRequest{UserId: userId, FollowId: followId}
	resp, err := userClient.JudgeFollow(context.Background(), &req)
	if err != nil {
		return false, err
	}
	return resp.Is, nil
}
