package logic

import (
	"context"

	"TikTokLite/app/message/cmd/rpc/internal/svc"
	"TikTokLite/app/message/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ActionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ActionLogic {
	return &ActionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ActionLogic) Action(in *pb.DouyinMessageActionRequest) (*pb.DouyinMessageActionResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.DouyinMessageActionResponse{}, nil
}
