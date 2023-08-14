package logic

import (
	"TikTokLite/app/message/cmd/rpc/message"
	"TikTokLite/common/ctxdata"
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
	in := &message.DouyinMessageChatRequest{
		FromUserId: ctxdata.GetUidFromCtx(l.ctx), // 从上下文中拿到id就是发送请求的用户id
		ToUserId:   req.ToUserId,
		PreMsgTime: req.PreMsgTime,
	}

	res, err := l.svcCtx.MessageRpcClient.Chat(l.ctx, in)

	if err != nil {
		return &types.DouyinMessageChatResponse{
			StatusCode:  res.StatusCode,
			StatusMsg:   res.GetStatusMsg(),
			MessageList: nil,
		}, err
	}

	messageList := make([]types.Message, len(res.MessageList))
	for i, v := range res.MessageList {
		messageList[i] = types.Message{
			Id:         v.Id,
			ToUserId:   v.ToUserId,
			FromUserId: v.FromUserId,
			Content:    v.Content,
			CreateTime: v.CreateTime,
		}
	}
	return &types.DouyinMessageChatResponse{
		StatusCode:  res.StatusCode,
		StatusMsg:   res.GetStatusMsg(),
		MessageList: messageList,
	}, err
}
