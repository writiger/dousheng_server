package handler

import (
	"context"
	"dousheng_server/middleware"
	"dousheng_server/rpc"
	"dousheng_server/user_service/dal/model"
	"errors"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"strconv"
)

// CheckUser 登录校验
func CheckUser(ctx context.Context, c *app.RequestContext) {
	middleware.JwtMiddleware.LoginHandler(ctx, c)
}

// Register 注册
func Register(ctx context.Context, c *app.RequestContext) {
	// 1. 验证参数
	username := c.Query("username")
	password := c.Query("password")
	if username == "" || password == "" || len(username) > 32 || len(password) > 32 {
		c.JSON(consts.StatusBadRequest, utils.H{
			"status_code": -1,
			"status_msg":  errors.New("wrong request param"),
		})
		return
	}
	// 2. 调用rpc服务
	err := rpc.Register(username, password)
	// 3. 返回
	if err != nil {
		c.JSON(consts.StatusServiceUnavailable, utils.H{
			"status_code": -1,
			"status_msg":  err.Error(),
		})
		return
	}
	// 4. 登录
	middleware.JwtMiddleware.LoginHandler(ctx, c)
}

// Info 用户信息
func Info(ctx context.Context, c *app.RequestContext) {
	// TODO 判断是否关注

	// 1. 验证参数
	idStr := c.Query("user_id")
	requester, _ := c.Get("identity")
	fmt.Println(requester)

	parseInt, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(consts.StatusServiceUnavailable, utils.H{
			"status_code": -1,
			"status_msg":  "wrong request param" + err.Error(),
		})
		return
	}
	userModel, err := rpc.GetUserInfo(parseInt)
	if err != nil {
		c.JSON(consts.StatusServiceUnavailable, utils.H{
			"status_code": -1,
			"status_msg":  err.Error(),
		})
		return
	}
	c.JSON(consts.StatusOK, utils.H{
		"status_code": 0,
		"status_msg":  "success",
		"user":        userModel,
	})
}

// Follow 关注
func Follow(ctx context.Context, c *app.RequestContext) {
	//rpc.Follow()
	toId := c.Query("to_user_id")
	actionType := c.Query("action_type")
	requester, _ := c.Get("identity")
	fmt.Println(requester)
	userId := requester.(*model.User).UUID
	followId, err := strconv.ParseInt(toId, 10, 64)
	if err != nil {
		c.JSON(consts.StatusServiceUnavailable, utils.H{
			"status_code": -1,
			"status_msg":  "wrong request param" + err.Error(),
		})
		return
	}
	switch actionType {
	case "1":
		err := rpc.Follow(userId, followId)
		if err != nil {
			c.JSON(consts.StatusServiceUnavailable, utils.H{
				"status_code": -1,
				"status_msg":  "wrong request param" + err.Error(),
			})
			return
		}
	case "2":
		rpc.CancelFollow(userId, followId)
		if err != nil {
			c.JSON(consts.StatusServiceUnavailable, utils.H{
				"status_code": -1,
				"status_msg":  "wrong request param" + err.Error(),
			})
			return
		}
	default:
		c.JSON(consts.StatusServiceUnavailable, utils.H{
			"status_code": -1,
			"status_msg":  "wrong request param" + err.Error(),
		})
		return
	}
	c.JSON(consts.StatusOK, utils.H{
		"status_code": 0,
		"status_msg":  "success",
	})
	return
}

// 关注列表
func FollowList(ctx context.Context, c *app.RequestContext) {
	Id := c.Query("user_id")

	requester, _ := c.Get("identity")
	fmt.Println(Id)

	userId := requester.(*model.User).UUID

	userModel, err := rpc.FollowList(userId)
	if err != nil {
		c.JSON(consts.StatusServiceUnavailable, utils.H{
			"status_code": -1,
			"status_msg":  err.Error(),
			"user_list":   nil,
		})
		return
	}
	c.JSON(consts.StatusOK, utils.H{
		"status_code": 0,
		"status_msg":  "success",
		"user_list":   userModel,
	})
	return
}

// 粉丝列表
func FollowerList(ctx context.Context, c *app.RequestContext) {
	Id := c.Query("user_id")

	requester, _ := c.Get("identity")
	fmt.Println(Id)

	userId := requester.(*model.User).UUID

	userModel, err := rpc.FollowerList(userId)
	if err != nil {
		c.JSON(consts.StatusServiceUnavailable, utils.H{
			"status_code": -1,
			"status_msg":  err.Error(),
			"user_list":   nil,
		})
		return
	}
	c.JSON(consts.StatusOK, utils.H{
		"status_code": 0,
		"status_msg":  "success",
		"user_list":   userModel,
	})
	return
}

// 好友列表
func FriendList(ctx context.Context, c *app.RequestContext) {
	Id := c.Query("user_id")

	requester, _ := c.Get("identity")
	fmt.Println(Id)

	userId := requester.(*model.User).UUID

	userModel, err := rpc.FriendList(userId)
	if err != nil {
		c.JSON(consts.StatusServiceUnavailable, utils.H{
			"status_code": -1,
			"status_msg":  err.Error(),
			"user_list":   nil,
		})
		return
	}
	c.JSON(consts.StatusOK, utils.H{
		"status_code": 0,
		"status_msg":  "success",
		"user_list":   userModel,
	})
	return
}
