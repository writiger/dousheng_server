package dal

import (
	"dousheng_server/user_service/util"
	"fmt"
	"testing"
)

func TestCreateUser(t *testing.T) {
	fmt.Println(util.HashAndSalt("testing"))

	//user := model.User{
	//	Username:      "test",
	//	Password:      "test",
	//	FollowCount:   0,
	//	FollowerCount: 0,
	//}
	//err := query.CreateUser()
	//if err != nil {
	//	return
	//}
}
