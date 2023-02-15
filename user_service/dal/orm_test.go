package dal

import (
	"dousheng_server/user_service/dal/model"
	"dousheng_server/user_service/dal/query"
	"fmt"
	"testing"
)

func TestCreateUser(t *testing.T) {
	user := model.User{
		Username:      "test",
		Password:      "test",
		FollowCount:   0,
		FollowerCount: 0,
	}
	for i := 0; i < 10; i++ {
		err := query.CreateUser(&user)

		if err != nil {
			return
		}
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
	// TODO
}
