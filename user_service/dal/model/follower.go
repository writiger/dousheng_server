package model

import "gorm.io/gorm"

type Follower struct {
	gorm.Model
	User   User
	UserID uint
}
