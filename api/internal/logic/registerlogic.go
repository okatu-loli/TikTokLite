package logic

import (
	"TikTokLite/api/internal/middleware"
	"TikTokLite/api/internal/model"
	"TikTokLite/api/internal/svc"
	"TikTokLite/api/internal/types"
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterRequest) (resp *types.RegisterResponse, err error) {
	if req.Username == "" || req.Password == "" {
		return nil, errors.New("username and password must not be empty")
	}

	// Check if username already exists
	var exists model.AuthUser
	err = l.svcCtx.DbEngin.Where("username = ?", req.Username).First(&exists).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			l.Error("Error checking existing username:", err)
			return nil, err // 其他数据库错误
		}
		// 用户名未找到，可以继续注册流程
	} else {
		return nil, errors.New("username already exists") // 用户名已存在
	}

	// Hashing the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		l.Error("Error hashing password:", err)
		return nil, err
	}

	// Using a transaction
	tx := l.svcCtx.DbEngin.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return nil, err
	}

	// Create new user in AuthUser table
	authUser := model.AuthUser{
		Username: req.Username,
		Password: string(hashedPassword),
	}
	if err = tx.Create(&authUser).Error; err != nil {
		tx.Rollback()
		l.Error("Error creating auth user:", err)
		return nil, err
	}

	// Create corresponding user info in UserInfo table
	userInfo := model.UserInfo{
		ID:   authUser.ID,
		Name: authUser.Username,
	}
	if err = tx.Create(&userInfo).Error; err != nil {
		tx.Rollback()
		l.Error("Error creating user info:", err)
		return nil, err
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		l.Error("Error committing transaction:", err)
		return nil, err
	}

	token, err := middleware.GenerateToken(authUser.ID)
	if err != nil {
		l.Error("Error generating token:", err)
		return nil, err
	}

	return &types.RegisterResponse{
		StatusCode: 0,
		StatusMsg:  "success",
		UserId:     authUser.ID,
		Token:      token,
	}, nil
}
