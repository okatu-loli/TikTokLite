package types

type RegisterRequest struct {
	Username string `gorm:"unique;not null" ;json:"username" ;validate:"required"`
	Password string `gorm:"not null" ;json:"password" ;validate:"required"`
}

type RegisterResponse struct {
	StatusCode int    `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
	UserId     uint   `json:"user_id"`
	Token      string `json:"token"`
}
