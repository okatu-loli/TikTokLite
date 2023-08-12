package logic

import (
	"TikTokLite/api/internal/middleware"
	"TikTokLite/api/internal/model"
	"TikTokLite/api/internal/types"
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"TikTokLite/api/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	if req.Username == "" || req.Password == "" {
		return nil, errors.New("username and password must not be empty")
	}

	// Find user by username
	var user model.AuthUser
	err = l.svcCtx.DbEngin.Where("username = ?", req.Username).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("username not found")
		}
		l.Error("Error retrieving user:", err)
		return nil, err
	}

	// Validate password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		// Passwords do not match
		l.Error("Error comparing passwords:", err)
		return nil, errors.New("incorrect password")
	}

	// Generate token (this part needs your actual token generation logic)
	token, err := middleware.GenerateToken(user.ID)
	if err != nil {
		l.Error("Error generating token:", err)
		return nil, err
	}

	return &types.LoginResponse{
		StatusCode: 0,
		StatusMsg:  "success",
		UserId:     user.ID,
		Token:      token,
	}, nil
}
