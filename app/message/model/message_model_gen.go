// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"fmt"
	"time"

	"github.com/SpectatorNan/gorm-zero/gormc"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var (
	cacheTiktokliteMessageIdPrefix = "cache:tiktoklite:message:id:"
)

type (
	messageModel interface {
		Insert(ctx context.Context, tx *gorm.DB, data *Message) error

		FindOne(ctx context.Context, id int64) (*Message, error)
		Update(ctx context.Context, tx *gorm.DB, data *Message) error

		Delete(ctx context.Context, tx *gorm.DB, id int64) error
		Transaction(ctx context.Context, fn func(db *gorm.DB) error) error
	}

	defaultMessageModel struct {
		gormc.CachedConn
		table string
	}

	Message struct {
		Id         int64     `gorm:"column:id"`           // 主键
		FromUserId int64     `gorm:"column:from_user_id"` // 发送人id
		ToUserId   int64     `gorm:"column:to_user_id"`   // 接收方id
		Content    string    `gorm:"column:content"`      // 发送内容
		CreateTime time.Time `gorm:"column:create_time"`  // 创建时间
		UpdateTime time.Time `gorm:"column:update_time"`  // 更新时间
	}
)

func (Message) TableName() string {
	return "`message`"
}

func newMessageModel(conn *gorm.DB, c cache.CacheConf) *defaultMessageModel {
	return &defaultMessageModel{
		CachedConn: gormc.NewConn(conn, c),
		table:      "`message`",
	}
}

func (m *defaultMessageModel) Insert(ctx context.Context, tx *gorm.DB, data *Message) error {

	err := m.ExecCtx(ctx, func(conn *gorm.DB) error {
		db := conn
		if tx != nil {
			db = tx
		}
		return db.Save(&data).Error
	}, m.getCacheKeys(data)...)
	return err
}

func (m *defaultMessageModel) FindOne(ctx context.Context, id int64) (*Message, error) {
	tiktokliteMessageIdKey := fmt.Sprintf("%s%v", cacheTiktokliteMessageIdPrefix, id)
	var resp Message
	err := m.QueryCtx(ctx, &resp, tiktokliteMessageIdKey, func(conn *gorm.DB, v interface{}) error {
		return conn.Model(&Message{}).Where("`id` = ?", id).First(&resp).Error
	})
	switch err {
	case nil:
		return &resp, nil
	case gormc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultMessageModel) Update(ctx context.Context, tx *gorm.DB, data *Message) error {
	old, err := m.FindOne(ctx, data.Id)
	if err != nil && err != ErrNotFound {
		return err
	}
	err = m.ExecCtx(ctx, func(conn *gorm.DB) error {
		db := conn
		if tx != nil {
			db = tx
		}
		return db.Save(data).Error
	}, m.getCacheKeys(old)...)
	return err
}

func (m *defaultMessageModel) getCacheKeys(data *Message) []string {
	if data == nil {
		return []string{}
	}
	tiktokliteMessageIdKey := fmt.Sprintf("%s%v", cacheTiktokliteMessageIdPrefix, data.Id)
	cacheKeys := []string{
		tiktokliteMessageIdKey,
	}
	cacheKeys = append(cacheKeys, m.customCacheKeys(data)...)

	return cacheKeys
}

func (m *defaultMessageModel) Delete(ctx context.Context, tx *gorm.DB, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		if err == ErrNotFound {
			return nil
		}
		return err
	}
	err = m.ExecCtx(ctx, func(conn *gorm.DB) error {
		db := conn
		if tx != nil {
			db = tx
		}
		return db.Delete(&Message{}, id).Error
	}, m.getCacheKeys(data)...)
	return err
}

func (m *defaultMessageModel) Transaction(ctx context.Context, fn func(db *gorm.DB) error) error {
	return m.TransactCtx(ctx, fn)
}

func (m *defaultMessageModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheTiktokliteMessageIdPrefix, primary)
}

func (m *defaultMessageModel) queryPrimary(conn *gorm.DB, v, primary interface{}) error {
	return conn.Model(&Message{}).Where("`id` = ?", primary).Take(v).Error
}
