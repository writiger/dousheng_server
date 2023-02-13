// Code generated by hertz generator.

package main

import (
	"dousheng_server/biz/handler"
	"dousheng_server/middleware"
	"github.com/cloudwego/hertz/pkg/app/server"
)

// customizeRegister registers customize routers.
func customizedRegister(r *server.Hertz) {
	// 静态资源映射
	r.Static("/static/", "/home/writiger/go/src/github.com/cloudwego/dousheng_server")

	dy := r.Group("/douyin")
	auth := r.Group("/douyin")
	auth.Use(middleware.JwtMiddleware.MiddlewareFunc())

	// 基础接口
	dy.GET("/feed/", handler.Feed)
	dy.POST("/user/register/", handler.Register)
	dy.POST("/user/login/", handler.CheckUser)
	auth.GET("/user/", handler.Info)
	auth.POST("/publish/action/", handler.Publish)
	auth.GET("/publish/list/", handler.VideoList)

	// 互动接口
	auth.POST("/favorite/action/", handler.VideoLike)
	auth.GET("/favorite/list/", handler.FavoriteList)
}
