# dousheng_server

字节跳动青训营大作业

# 运行方法

~~~ cmd

~~~

# 项目详情

## http框架：hertz

~~~ cmd
hz new 项目名
~~~

## 微服务内框架：kitex

~~~ cmd
kitex -module 模块名 -service 服务名 idl文件路径
~~~

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