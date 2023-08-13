package model

type AuthUser struct {
	ID       uint   `gorm:"primaryKey" ;json:"user_id"`
	Username string `gorm:"unique;not null;index" ;json:"username"` // 添加索引，并且是唯一的
	Password string `gorm:"not null" ;json:"password"`
}

type UserInfo struct {
	ID              uint   `json:"id"` // 设置外键约束
	Name            string `json:"name"`
	FollowCount     int    `json:"follow_count"`
	FollowerCount   int    `json:"follower_count"`
	IsFollow        bool   `json:"is_follow"`
	Avatar          string `json:"avatar"`
	BackgroundImage string `json:"background_image"`
	Signature       string `json:"signature"`
	TotalFavorited  string `json:"total_favorited"`
	WorkCount       int    `json:"work_count"`
	FavoriteCount   int    `json:"favorite_count"`
}
