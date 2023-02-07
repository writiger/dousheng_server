package query

import (
	"dousheng_server/user_service/dal/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var GormClient *gorm.DB

func init() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/dousheng?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	GormClient, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	err = GormClient.AutoMigrate(&model.User{})
	if err != nil {
		panic("gorm init table failed ")
	}
}
