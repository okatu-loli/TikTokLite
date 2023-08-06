package model

import (
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
	return []string{}
}
