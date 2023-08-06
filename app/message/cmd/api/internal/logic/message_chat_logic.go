package logic

import (
	"context"

	"TikTokLite/app/message/cmd/api/internal/svc"
	"TikTokLite/app/message/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MessageChatLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMessageChatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MessageChatLogic {
	return &MessageChatLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// MessageChat 获取消息列表
func (l *MessageChatLogic) MessageChat(req *types.DouyinMessageChatRequestt) (resp *types.DouyinMessageChatResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
