package logic

import (
	"TikTokLite/app/message/cmd/rpc/message"
	"TikTokLite/common/ctxdata"
	"TikTokLite/common/dyerr"
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

// MessageAction 发送消息
func (l *MessageActionLogic) MessageAction(req *types.DouyinMessageActionRequest) (resp *types.DouyinMessageActionResponse, err error) {
	// todo: add your logic here and delete this line
	if req.ActionType != 1 {
		statusMsg := "参数错误"
		return &types.DouyinMessageActionResponse{
			StatusCode: int32(dyerr.NO_REQUEST),
			StatusMsg:  statusMsg,
		}, nil
	}

	in := &message.DouyinMessageActionRequest{
		FromUserId: ctxdata.GetUidFromCtx(l.ctx), // 从上下文中拿到的id就是发送请求的用户的id
		ToUserId:   req.ToUserId,
		ActionType: req.ActionType,
		Content:    req.Content,
	}
	res, err := l.svcCtx.MessageRpcClient.Action(l.ctx, in)
	return &types.DouyinMessageActionResponse{
		StatusCode: res.StatusCode,
		StatusMsg:  res.GetStatusMsg(),
	}, err
}
