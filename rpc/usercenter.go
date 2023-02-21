// Package rpc
// 封装与service的请求与返回
package rpc

import (
	"context"
	"dousheng_server/conf"
	"dousheng_server/user_service/kitex_gen"
	"dousheng_server/user_service/kitex_gen/usercenter"
	"errors"
	"fmt"
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

type UserInfo struct {
	UUID            int64  `gorm:"primaryKey" json:"id"`
	UserName        string `json:"name"`
	FollowCount     int64  `json:"follow_count"`
	FollowerCount   int64  `json:"follower_count"`
	IsFollow        bool   `json:"is_follow"`
	Avatar          string `json:"avatar"`
	BackgroundImage string `json:"background_image"`
	Signature       string `json:"signature"`
	TotalFavorited  int64  `json:"total_favorited"`
	WorkCount       int64  `json:"work_count"`
	FavoriteCount   int64  `json:"favorite_count"`
}
type Message struct {
	// UserId 关注了 FollowID
	Id         int64  `json:"id"`
	ToUserId   int64  `json:"to_user_id"`
	FromUserId int64  `json:"from_user_id"`
	Messages   string `json:"content"`
	CreatedAt  int64  `json:"create_time"`
}
type Friend struct {
	Message string `json:"message"`
	MsgType int64  `json:"msgType"`
	*UserInfo
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
	totalFavorited, err := BePraisedCounts(userId)
	if err != nil {
		return nil, err
	}
	workCount, err := WorkCounts(userId)
	if err != nil {
		return nil, err
	}
	favoriteCount, err := FavouriteCounts(userId)
	if err != nil {
		return nil, err
	}
	userInfo := UserInfo{
		UUID:            resp.User.Id,
		UserName:        resp.User.Name,
		FollowCount:     resp.User.FollowCount,
		FollowerCount:   resp.User.FollowerCount,
		IsFollow:        isFollowed,
		Avatar:          "http://192.168.101.112:8080/static/covers/img.png",
		BackgroundImage: "http://192.168.101.112:8080/static/covers/img.png",
		Signature:       "个人简介为空",
		TotalFavorited:  totalFavorited,
		WorkCount:       workCount,
		FavoriteCount:   favoriteCount,
	}
	fmt.Println("前端作品数:", workCount)
	fmt.Println("前端喜欢数:", favoriteCount)
	fmt.Println("前端获赞数:", totalFavorited)
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
func FollowList(userId int64) ([]UserInfo, error) {
	var res []UserInfo
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
		res = append(res, *userInfo)
	}
	return res, nil
}

// FollowerList 粉丝列表
func FollowerList(userId int64) ([]UserInfo, error) {
	var res []UserInfo
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
		if err != nil {
			return nil, err
		}
		res = append(res, *userInfo)
	}
	return res, nil
}

// FriendList 好友列表
func FriendList(userId int64) ([]Friend, error) {
	var res []Friend
	var msg Message
	var msgList []Message
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
		msgList, err = MessageList(userId, userInfo.UUID, 0)
		if err != nil {
			return nil, err
		}
		var msgType int64
		msgType = 0
		content := ""
		if msgList != nil {
			msg = msgList[len(msgList)-1]
			if msg.FromUserId == userId {
				msgType = 1
			}
			content = msg.Messages
		}
		res = append(res, Friend{
			UserInfo: userInfo,
			Message:  content,
			MsgType:  msgType,
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
func MessageList(fromUserId, ToUserId, lastTime int64) ([]Message, error) {
	var messageList []Message
	req := kitex_gen.MessageListRequest{
		UserId:   fromUserId,
		ToId:     ToUserId,
		LastTime: lastTime,
	}
	resp, err := userClient.MessageList(context.Background(), &req)
	if err != nil {
		return nil, err
	}
	if resp == nil || len(resp.MessageList) == 0 {
		return nil, err
	}
	for _, item := range resp.MessageList {
		messageList = append(messageList, Message{
			Id:         item.Id,
			FromUserId: item.FromUserId,
			ToUserId:   item.ToUserId,
			Messages:   item.Content,
			CreatedAt:  item.CreateTime,
		})
	}

	return messageList, err
}

// 作品数
func WorkCounts(uuid int64) (int64, error) {
	var count int64
	req := kitex_gen.GetInfoRequest{Uuid: uuid}
	item, err := userClient.WorkCounts(context.Background(), &req)
	if err != nil {
		return 0, err
	}
	if item != nil {
		count = item.Counts
	}
	return count, err
}

// 喜欢数
func FavouriteCounts(uuid int64) (int64, error) {
	var count int64
	req := kitex_gen.GetInfoRequest{Uuid: uuid}
	item, err := userClient.FavouriteCounts(context.Background(), &req)
	if err != nil {
		return 0, err
	}
	if item != nil {
		count = item.Counts
	}
	return count, err
}

// 获赞数
func BePraisedCounts(uuid int64) (int64, error) {
	var count int64
	req := kitex_gen.GetInfoRequest{Uuid: uuid}
	item, err := userClient.BePraisedCounts(context.Background(), &req)
	if err != nil {
		return 0, err
	}
	if item != nil {
		count = item.Counts
	}
	return count, err
}
