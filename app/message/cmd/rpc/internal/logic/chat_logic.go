package logic

import (
	"TikTokLite/app/message/cmd/rpc/internal/svc"
	"TikTokLite/app/message/cmd/rpc/pb"
	"TikTokLite/common/dyerr"
	"context"

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

// Chat 获取消息
func (l *ChatLogic) Chat(in *pb.DouyinMessageChatRequest) (*pb.DouyinMessageChatResponse, error) {
	// todo: add your logic here and delete this line
	fromUserId := in.FromUserId
	messages, err := l.svcCtx.MessageModel.FindChats(l.ctx, fromUserId, in.ToUserId, in.PreMsgTime)

	if err != nil {
		statusMsg := err.Error()
		return &pb.DouyinMessageChatResponse{
			StatusCode:  int32(dyerr.DB_ERROR),
			StatusMsg:   &statusMsg,
			MessageList: nil,
		}, err
	}

	messageList := make([]*pb.Message, 0)
	for _, m := range messages {
		t := &pb.Message{
			Id:         m.Id,
			ToUserId:   m.ToUserId,
			FromUserId: m.FromUserId,
			Content:    m.Content,
			CreateTime: m.CreateTime.Unix(),
		}
		messageList = append(messageList, t)
	}
	return &pb.DouyinMessageChatResponse{
		StatusCode:  int32(dyerr.OK),
		StatusMsg:   nil,
		MessageList: messageList,
	}, nil
}
