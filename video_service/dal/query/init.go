package query

import (
	"dousheng_server/conf"
	zaplog "dousheng_server/deploy/log"
	"dousheng_server/video_service/dal/model"
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
		zaplog.ZapLogger.Error("failed when opening gorm:%s err:%v", conf.Conf.Database, err)
		panic(err.Error())
	}
	err = GormClient.AutoMigrate(&model.Comment{}, &model.Video{}, &model.Favorite{})
	if err != nil {
		zaplog.ZapLogger.Error("failed when init gorm table err:%v", err)
		panic("gorm init table failed ")
	}
}
