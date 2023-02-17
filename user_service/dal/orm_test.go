package dal

import (
	"dousheng_server/user_service/dal/model"
	"dousheng_server/user_service/dal/query"
	"dousheng_server/user_service/util"
	"fmt"
	"testing"
)

func TestCreateUser(t *testing.T) {
	passwd, _ := util.HashAndSalt("test")
	user := model.User{
		Username:      "test",
		Password:      passwd,
		FollowCount:   0,
		FollowerCount: 0,
	}
	err := query.CreateUser(&user)
	if err != nil {
		return
	}
}

func TestCheckPassword(t *testing.T) {
	type args struct {
		name     string
		password string
	}

	tests := []struct {
		name    string
		args    args
		wantRes bool
	}{
		{
			name: "1. 正确测试",
			args: args{
				name:     "test",
				password: "test",
			},
			wantRes: true,
		},
		{
			name: "2. 错误测试",
			args: args{
				name:     "test",
				password: "world",
			},
			wantRes: false,
		},
		{
			name: "3. 账号不存在测试",
			args: args{
				name:     "hello2",
				password: "test",
			},
			wantRes: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := query.CheckPassword(tt.args.name, tt.args.password)
			if err != nil && tt.wantRes {
				t.Fatalf("%s 出错,err : %v", tt.name, err)
				return
			}
			t.Log(tt.name, "Pass")
		})
	}
}

func TestGetInfo(t *testing.T) {
	t.Run("用户存在的测试", func(t *testing.T) {
		user1, err := query.GetUser(629648834103869440)
		if err != nil {
			t.Fatalf("从用户存在的测试失败")
			return
		}
		fmt.Println(user1)
	})
	t.Run("用户不存在的测试", func(t *testing.T) {
		user2, err := query.GetUser(629440)
		if err != nil {
			t.Fatalf("从用户不存在的测试失败")
			return
		}
		fmt.Println(user2)
	})
}

func TestForeignKey(t *testing.T) {
	var users []model.User
	err := query.GormClient.Model(&model.User{}).Preload("Videos").Find(&users).Error
	if err != nil {
		t.Fatalf(err.Error())
		return
	}
	fmt.Println(users)
}

func TestFollow(t *testing.T) {
	//asd关注了asdasd
	var userId int64
	var followId int64
	userId = 631992324594860032
	followId = 632238110008147968
	err := query.Follow(userId, followId)
	err2 := query.Follow(followId, userId)
	if err != nil {
		fmt.Println("第一步互相关注出错", err)
		return
	}
	if err2 != nil {
		fmt.Println("第一步互相关注出错", err)
		return
	}
	//test1关注了test123
	//userId = 631640850035638273
	//followId = 631640850035638274
	//err := query.Follow(userId, followId)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
}

func TestCancelFollow(t *testing.T) {
	//asd取消关注了asdasd
	var userId int64
	var followId int64
	//userId = 631992324594860032
	//followId = 632238110008147968
	//err := query.CancelFollow(userId, followId)
	//if err != nil {
	//	fmt.Println("第一步互相关注出错", err)
	//	return
	//}
	//test1取消关注了test123
	userId = 631640850035638273
	followId = 631640850035638274
	err := query.CancelFollow(userId, followId)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func TestFollowList(t *testing.T) {

	user, err := query.FollowList(631640850035638272)
	if err != nil {
		fmt.Println("错误")
	}

	for _, i := range *user {
		fmt.Println(i)
	}
}

func TestFollowerList(t *testing.T) {

	user, err := query.FollowerList(631640850035638273)
	if err != nil {
		fmt.Println("错误")
	}

	for _, i := range *user {
		fmt.Println(i)
	}
}

func TestFriendList(t *testing.T) {

	user, err := query.FriendList(631640850035638272)
	if err != nil {
		fmt.Println("错误")
	}

	for _, i := range *user {
		fmt.Println(i)
	}
}
