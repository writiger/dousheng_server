package handler

import (
	"context"
	"dousheng_server/middleware"
	"dousheng_server/rpc"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// CheckUser 登录校验
func CheckUser(ctx context.Context, c *app.RequestContext) {
	middleware.JwtMiddleware.LoginHandler(ctx, c)
}

// Register 注册
func Register(c context.Context, ctx *app.RequestContext) {
	// 1. 验证参数
	username := ctx.Query("username")
	password := ctx.Query("password")
	if username == "" || password == "" || len(username) > 32 || len(password) > 32 {
		ctx.JSON(consts.StatusBadRequest, utils.H{
			"status_code": -1,
			"status_msg":  errors.New("wrong request param"),
		})
		return
	}
	// 2. 调用rpc服务
	err := rpc.Register(username, password)
	// 3. 返回
	if err != nil {
		ctx.JSON(consts.StatusServiceUnavailable, utils.H{
			"status_code": -1,
			"status_msg":  err.Error(),
		})
		return
	}
	// 4. 登录
	middleware.JwtMiddleware.LoginHandler(c, ctx)
}
