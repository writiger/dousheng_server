# dousheng_server

----字节跳动青训营大作业

# 运行方法

## 配置

* 修改`dousheng_server/video_service/service/video_center.go`中的saverIp(后续可能会使用配置文件)

* 调整各个服务下的etcd地址（main.go中）
* 调整各个服务中的MySQL地址（dal/query/init.go）
* 执行query中的测试

~~~ cmd
# wsl情况下
# 启动mysql
service mysql start
# 启动etcd
etcd
# 进入项目目录
# 启动uuid生成服务
sh snowflake_service/build.sh && sh snowflake_service/output/bootstrap.sh
# 启动用户服务
sh user_service/build.sh && sh user_service/output/bootstrap.sh
# 启动video服务
sh video_service/build.sh && sh video_service/output/bootstrap.sh
# 启动hertz
go build -o dousheng_server && ./dousheng_server
~~~

# TODO List

## 基础接口

- [x] 视频流
- [x] 用户注册
- [x] 用户登录
- [x] 用户信息
- [x] 视频投稿
- [x] 发布列表

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
| user | 8900 |
| snowflake | 8901 |
| video | 8902 |



## orm框架：gorm

### user模型

~~~ golang
type User struct {
	UUID          int64         `gorm:"primaryKey" json:"id"`
	Username      string        `json:"username"`
	Password      string        `json:"password"`
	FollowCount   int64         `json:"follow_count"`
	FollowerCount int64         `json:"follower_count"`
	Videos        []model.Video `gorm:"foreignKey:user_id"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}
~~~

### video模型

~~~ golang
type Video struct {
	UUID          int64  `json:"id"`
	UserID        int64  `json:"user_id"`
	PlayURL       string `json:"play_url"`
	CoverURL      string `json:"cover_url"`
	FavoriteCount int64  `json:"favorite_count"`
	CommentCount  int64  `json:"comment_count"`
	Title         string `json:"title"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}
~~~



##  优化

* UUID使用雪花算法生成
* 使用etcd的lease机制维护雪花算法的生成ID
* 密码使用bcrypt加盐