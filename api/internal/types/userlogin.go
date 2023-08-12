package types

type LoginRequest struct {
	Username string `validate:"required"`
	Password string `validate:"required"`
}

type LoginResponse struct {
	StatusCode int    `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
	UserId     uint   `json:"user_id"`
	Token      string `json:"token"`
}
