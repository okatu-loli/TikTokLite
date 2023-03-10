package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName      string `json:"username"`       // 昵称
	Password      string `json:"password"`       // 密码, MD5 加密后不可逆
	FollowCount   int    `json:"follow_count"`   // 关注人数
	FollowerCount int    `json:"follower_count"` // 粉丝人数
}
