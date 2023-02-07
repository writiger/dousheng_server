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
