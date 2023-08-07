package model

import (
	"context"
	"gorm.io/gorm"
)

var (
	// 缓存的键，二者较大的id在前
	cacheChatsFromToIdPrefix = "cache:tiktoklite:message:one:another:id:"
)

type (
	myMessageModel interface {
		FindChats(ctx context.Context, fromUserId, toUserId int64, preMsgTime int64) ([]Message, error)
	}
)

func (m *defaultMessageModel) FindChats(ctx context.Context, fromUserId, toUserId int64, preMsgTime int64) ([]Message, error) {
	chatsKey := m.getChatsKey(fromUserId, toUserId)
	var messageList []Message
	err := m.QueryCtx(ctx, &messageList, chatsKey, func(conn *gorm.DB, v interface{}) error {
		err := conn.Model(&Message{}).Where(&Message{FromUserId: fromUserId, ToUserId: toUserId}).
			Or(&Message{FromUserId: toUserId, ToUserId: fromUserId}).Find(&messageList).Error
		return err
	})
	return messageList, err
}
