package logic

import (
	"context"

	"TikTokLite/app/message/cmd/api/internal/svc"
	"TikTokLite/app/message/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MessageActionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMessageActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MessageActionLogic {
	return &MessageActionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MessageActionLogic) MessageAction(req *types.DouyinMessageActionRequest) (resp *types.DouyinMessageActionResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
