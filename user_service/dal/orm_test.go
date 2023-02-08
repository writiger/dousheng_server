package dal

import (
	"dousheng_server/user_service/dal/model"
	"dousheng_server/user_service/dal/query"
	"dousheng_server/user_service/util"
	"fmt"
	"testing"
)

func TestCreateUser(t *testing.T) {
	fmt.Println(util.HashAndSalt("testing"))
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
				name:     "hello",
				password: "test",
			},
			wantRes: true,
		},
		{
			name: "2. 错误测试",
			args: args{
				name:     "hello",
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
			res, err := query.CheckPassword(tt.args.name, tt.args.password)
			if err != nil {
				t.Fatalf("%s 出错,err : %v", tt.name, err)
				return
			}
			if res != tt.wantRes {
				t.Fatal("结果出错")
				return
			}
			t.Log(tt.name, "Pass")
		})
	}
}
