package types

import "TikTokLite/api/internal/model"

type UserInfoRequest struct {
	UserId int    `json:"user_id" validate:"required"`
	Token  string `json:"token" validate:"required"`
}

type UserInfoResponse struct {
	StatusCode int            `json:"status_code"`
	StatusMsg  string         `json:"status_msg"`
	User       model.UserInfo `json:"user"`
}
