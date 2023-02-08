# dousheng_server

字节跳动青训营大作业

# 运行方法

~~~ cmd

~~~

# TODO List

## 基础接口

- [ ] 视频流
- [x] 用户注册
- [x] 用户登录
- [ ] 用户信息
- [ ] 视频投稿
- [ ] 发布列表

## 互动接口

- [ ] 赞
- [ ] 喜欢列表
- [ ] 评论
- [ ] 视频评论列表

## 社交接口

- [ ] 关系操作
- [ ] 用户关注列表
- [ ] 用户粉丝列表
- [ ] 用户好友列表



# 项目详情

## http框架：hertz

~~~ cmd
go build -o dousheng_server && ./dousheng_server
~~~

## 微服务内框架：kitex

~~~ cmd
kitex -module dousheng_server -service 服务名 idl文件路径
~~~

## 服务启动命令

~~~ cmd
sh build.sh && sh output/bootstrap.sh
~~~



### 服务端口预分配

| 服务 | 端口 |
| --- | ----------- |
| hertz | 8080 |
| userservice | 8900 |
| snowflakeservice | 8901 |

## orm框架：gorm

### user模型

~~~ golang
type User struct {
	gorm.Model
	Username      string `json:"username"`
	Password      string `json:"password"`
	FollowCount   int64  `json:"followCount"`
	FollowerCount int64  `json:"followerCount"`
}
~~~

##  优化

* UUID使用雪花算法生成
* 使用etcd的lease机制维护雪花算法的生成ID
* 密码使用bcrypt加盐