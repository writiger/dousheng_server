package model

type Follower struct {
	// UserId 关注了 FollowID
	UserId   int64 `gorm:"primaryKey" json:"userId"`
	FollowId int64 `gorm:"primaryKey" json:"followerId"`
}
