package query

import (
	"dousheng_server/conf"
	"dousheng_server/user_service/dal/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var GormClient *gorm.DB

func init() {
	dsn := conf.Conf.MysqlConfig.User + ":" +
		conf.Conf.MysqlConfig.Password + "@tcp(" +
		conf.Conf.MysqlConfig.Url + ")/" +
		conf.Conf.MysqlConfig.Database + "?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	GormClient, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	err = GormClient.AutoMigrate(&model.User{}, &model.Follower{})
	if err != nil {
		panic("gorm init table failed ")
	}
}
