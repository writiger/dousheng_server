package model

type Follower struct {
	// UserId 关注了 FollowID
	UserId   int64 `json:"userId"`
	FollowId int64 `json:"followerId"`
}
