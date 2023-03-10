# dousheng_server

字节跳动青训营大作业~~

# 运行方法

## 配置

* 新建conf.ini
* 复制conf.template.ini到conf.ini
* 修改conf.template.ini

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

- [x] 赞
- [x] 喜欢列表
- [x] 评论
- [x] 视频评论列表

## 社交接口

- [x] 关系操作
- [x] 用户关注列表
- [x] 用户粉丝列表
- [x] 用户好友列表
- [x] 发送消息
- [x] 消息列表



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

#### prometheus

~~~ cmd
prometheus --config.file=prometheus.yml
~~~

#### mysql-export

~~~ cmd
mysqld_exporter --config.my-cnf=my.cnf
~~~

#### grafana

~~~ cmd
sudo systemctl start grafana-server
~~~



### 服务端口预分配

| 服务         | 端口   |
|------------|------|
| hertz      | 8080、9980 |
| user       | 8900、9900 |
| snowflake  | 8901、9901 |
| video      | 8902、9902 |
| prometheus | 9090  |



## orm框架：gorm

### user模型

~~~ golang
type User struct {
	UUID          int64      `gorm:"primaryKey" json:"id"`
	Username      string     `json:"name"`
	Password      string     `json:"password"`
	FollowCount   int64      `json:"follow_count"`
	FollowerCount int64      `json:"follower_count"`
	Followers     []Follower `gorm:"foreignKey:user_id;foreignKey:follow_id"`
	Messages      []Message  `gorm:"foreignKey:from_user_id;foreignKey:to_user_id"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}
~~~

### Follower模型
~~~ golang
type Follower struct {
	// UserId 关注了 FollowID
	UserId   int64 `gorm:"primaryKey" json:"userId"`
	FollowId int64 `gorm:"primaryKey" json:"followerId"`
}
~~~

### Message模型
~~~ golang
type Message struct {
	//FromUserId对ToUserId发消息
	Id         int64  `gorm:"primaryKey" json:"id"`
	ToUserId   int64  `gorm:"to_user_id,index:idx_to" json:"to_user_id"`
	FromUserId int64  `gorm:"from_user_id,index:idx_from" json:"from_user_id"`
	Messages   string `gorm:"message" json:"content"`
	CreatedAt  int64  `gorm:"index:idx_create" json:"create_time"`
}
~~~

### video模型

~~~ golang
type Video struct {
	UUID          int64  `gorm:"primaryKey" json:"id"`
	UserID        int64  `json:"user_id"`
	PlayURL       string `json:"play_url"`
	CoverURL      string `json:"cover_url"`
	FavoriteCount int64  `json:"favorite_count"`
	CommentCount  int64  `json:"comment_count"`
	Title         string `json:"title"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
	//外键
	User model.User `gorm:"foreignKey:user_id"`
}
~~~

### favorite模型

~~~ golang
type Favorite struct {
	UserId  int64 `gorm:"index:idx_user" json:"user_id"`
	VideoId int64 `json:"video_id"`
	//外键
	User model.User `gorm:"foreignKey:user_id"`
}
~~~

### comment模型

~~~ golang
type Comment struct {
	CommentId  int64  `gorm:"primaryKey" json:"comment_id"`
	UserId     int64  `gorm:"index:idx_user" json:"user_id"`
	VideoId    int64  `gorm:"index:idx_video" json:"video_id"`
	Content    string `json:"content"`
	CreateDate string `gorm:"index:idx_create,sort:desc" json:"create_date"`
	CreatedAt  time.Time
}
~~~

##  优化

* UUID使用雪花算法生成
* 使用etcd的lease机制维护雪花算法
* 密码使用bcrypt加盐
* 使用令牌桶限流
