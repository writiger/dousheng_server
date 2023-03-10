package query

import (
	"dousheng_server/video_service/dal/model"
	"errors"
	"gorm.io/gorm"
)

// Favorite 点赞
func Favorite(userId, videoId int64) error {
	favorite := model.Favorite{
		UserId:  userId,
		VideoId: videoId,
	}
	if ok := IsLiked(userId, videoId); ok {
		return errors.New("already liked")
	}
	return GormClient.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(&favorite).Error
		if err != nil {
			return err
		} else {
			if err := tx.Where("uuid = ?", videoId).Model(&model.Video{}).
				Update("favorite_count", gorm.Expr("favorite_count + 1")).
				Error; err != nil {
				return err
			}
		}
		return nil
	})
}

// UndoFavorite 取消点赞
func UndoFavorite(userId, videoId int64) error {
	return GormClient.Transaction(func(tx *gorm.DB) error {
		err := tx.Where("user_id = ? AND video_id = ?", userId, videoId).Delete(&model.Favorite{}).Error
		if err != nil {
			return err
		} else {
			if err := tx.Where("uuid = ?", videoId).Model(&model.Video{}).
				Update("favorite_count", gorm.Expr("favorite_count - 1")).
				Error; err != nil {
				return err
			}
		}
		return nil
	})
}

// IsLiked 判断此视频该用户是否点过赞
func IsLiked(userId, videoId int64) bool {
	favorite := model.Favorite{}
	//err := GormClient.Where("video_id = ?", videoId).Find(&favorite).Error
	result := GormClient.Where("user_id = ? AND video_id = ?", userId, videoId).Find(&favorite)
	return result.RowsAffected != 0
}

// FavoriteList 点赞过的视频uuid列表
func FavoriteList(userId int64) ([]int64, error) {
	res := make([]int64, 0)
	var favorites []model.Favorite
	err := GormClient.Where("user_id = ?", userId).Find(&favorites).Error
	for _, item := range favorites {
		res = append(res, item.VideoId)
	}
	return res, err
}
