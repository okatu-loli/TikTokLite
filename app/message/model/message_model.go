package model

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var _ MessageModel = (*customMessageModel)(nil)

type (
	// MessageModel is an interface to be customized, add more methods here,
	// and implement the added methods in customMessageModel.
	MessageModel interface {
		messageModel
		customMessageLogicModel
		myMessageModel
	}

	customMessageModel struct {
		*defaultMessageModel
	}

	customMessageLogicModel interface {
	}
)

// NewMessageModel returns a model for the database table.
func NewMessageModel(conn *gorm.DB, c cache.CacheConf) MessageModel {
	return &customMessageModel{
		defaultMessageModel: newMessageModel(conn, c),
	}
}

func (m *defaultMessageModel) customCacheKeys(data *Message) []string {
	if data == nil {
		return []string{}
	}

	return []string{
		m.getChatsKey(data.FromUserId, data.ToUserId),
	}
}

func (m *defaultMessageModel) getChatsKey(id1, id2 int64) string {
	maxId := func(a, b int64) int64 {
		if a > b {
			return a
		}
		return b
	}(id1, id2)
	chatsKey := fmt.Sprintf("%s%v%v", cacheChatsFromToIdPrefix, maxId, id1+id2-maxId)
	return chatsKey
}
