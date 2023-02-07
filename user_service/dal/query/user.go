package query

import (
	"dousheng_server/user_service/dal/model"
)

func CreateUser(user *model.User) error {
	return GormClient.Create(user).Error
}
