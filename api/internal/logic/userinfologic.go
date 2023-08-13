package logic

import (
	"TikTokLite/api/internal/model"
	"TikTokLite/api/internal/types"
	"context"

	"TikTokLite/api/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req *types.UserInfoRequest) (resp *types.UserInfoResponse, err error) {
	var user model.UserInfo
	err = l.svcCtx.DbEngin.Where("id = ?", req.UserId).First(&user).Error
	if err != nil {
		return nil, err
	}
	user = model.UserInfo{
		ID:              user.ID,
		Name:            user.Name,
		FollowCount:     0,
		FollowerCount:   0,
		IsFollow:        false,
		Avatar:          "https://cdn.cnqs.moe/qianshi-cdn/2023/08/f7f1430693f6cea767e812a7873c6cae.png",
		BackgroundImage: "https://cdn.cnqs.moe/qianshi-cdn/2023/08/4eb194704ead538d3aea1e17d1103956.png",
		Signature:       "Hello, World",
		TotalFavorited:  "66",
		WorkCount:       0,
		FavoriteCount:   0,
	}
	return &types.UserInfoResponse{
		StatusCode: 0,
		StatusMsg:  "success",
		User:       user,
	}, nil
}
