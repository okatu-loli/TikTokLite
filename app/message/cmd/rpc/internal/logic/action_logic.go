package logic

import (
	"TikTokLite/app/message/model"
	"TikTokLite/common/dyerr"
	"context"
	"strconv"
	"time"

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

// Action 发送消息
func (l *ActionLogic) Action(in *pb.DouyinMessageActionRequest) (*pb.DouyinMessageActionResponse, error) {
	// todo: add your logic here and delete this line
	// todo: 仅供测试
	a, _ := strconv.Atoi(in.Token)
	fromUserId := int64(a)

	//fromUserId := ctxdata.GetUidFromCtx(l.ctx)
	msg := &model.Message{
		ToUserId:   in.ToUserId,
		FromUserId: fromUserId,
		Content:    in.Content,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	err := l.svcCtx.MessageModel.Insert(l.ctx, nil, msg)
	if err != nil {
		statusMsg := "数据库错误: " + err.Error()
		return &pb.DouyinMessageActionResponse{
			StatusCode: int32(dyerr.DB_ERROR),
			StatusMsg:  &statusMsg,
		}, err
	}

	statusMsg := "发送成功"
	return &pb.DouyinMessageActionResponse{
		StatusCode: int32(dyerr.OK),
		StatusMsg:  &statusMsg,
	}, nil
}
