package query

import (
	"dousheng_server/video_service/dal/model"
	"gorm.io/gorm"
)

// CreateComment 新增评论
func CreateComment(comment *model.Comment) (*model.Comment, error) {
	err := GormClient.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(comment).Error; err != nil {
			return err
		}
		if err := tx.Where("uuid = ?", comment.VideoId).Model(&model.Video{}).
			Update("comment_count", gorm.Expr("comment_count + 1")).
			Error; err != nil {
			return err
		}
		return nil
	})
	return comment, err
}

// DeleteComment 删除评论
func DeleteComment(uuid int64) error {
	return GormClient.Transaction(func(tx *gorm.DB) error {
		comment := &model.Comment{CommentId: uuid}
		if err := tx.First(comment).Error; err != nil {
			return err
		}
		if err := tx.Where("uuid = ?", comment.VideoId).Model(&model.Video{}).
			Update("comment_count", gorm.Expr("comment_count - 1")).
			Error; err != nil {
			return err
		}
		if err := tx.Delete(comment).Error; err != nil {
			return err
		}
		return nil
	})
}

// GetComment 查看视频的评论
func GetComment(uuid int64) (*[]model.Comment, error) {
	var res []model.Comment
	err := GormClient.Order("created_at desc").Where("video_id = ?", uuid).Find(&res).Error
	return &res, err
}
