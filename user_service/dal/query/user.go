package query

import "dousheng_server/user_service/dal/model"

func CreateUser(user *model.User) error {
	if err := GormClient.Create(user).Error; err != nil {
		return err
	}
	return nil
}
