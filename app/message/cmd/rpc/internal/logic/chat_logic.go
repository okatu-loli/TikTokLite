package logic

import (
	"context"

	"TikTokLite/app/message/cmd/rpc/internal/svc"
	"TikTokLite/app/message/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChatLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewChatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChatLogic {
	return &ChatLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ChatLogic) Chat(in *pb.DouyinMessageChatRequest) (*pb.DouyinMessageChatResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.DouyinMessageChatResponse{}, nil
}
