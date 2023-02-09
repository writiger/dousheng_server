package handler

import (
	"context"
	"dousheng_server/middleware"
	"dousheng_server/rpc"
	"dousheng_server/user_service/dal/model"
	"errors"
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
	// 1. 验证参数
	user := middleware.JwtMiddleware.IdentityHandler(ctx, c).(*model.User)
	if strconv.FormatInt(user.UUID, 10) != c.Query("user_id") {
		c.JSON(consts.StatusServiceUnavailable, utils.H{
			"status_code": -1,
			"status_msg":  errors.New("wrong user_id or token"),
		})
		return
	}

	userModel, err := rpc.GetUserInfo(user.UUID)
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
