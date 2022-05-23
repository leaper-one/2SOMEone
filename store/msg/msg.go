package msg

import (
	"2SOMEone/core"
	"2SOMEone/util"
	"context"
	"errors"

	"gorm.io/gorm"
)

func New(db *util.DB) core.MessageStore {
	return &msgStore{db: db}
}

type msgStore struct {
	db *util.DB
}

func toUpdateParms(msg *core.Message) map[string]interface{} {
	return map[string]interface{}{
		"phone":   msg.Phone,
		"type":    msg.Type,
		"content": msg.Content,
		"code":    msg.Code,
	}
}

func updateByphone(db *util.DB, msg *core.Message) (int64, error) {
	updates := toUpdateParms(msg)
	tx := db.Update().Model(msg).Where("phone = ?", msg.Phone).Updates(updates)
	return tx.RowsAffected, tx.Error
}

// Create a new message
func (m *msgStore) Create(ctx context.Context, msg *core.Message) error {
	return m.db.Tx(func(tx *util.DB) error {
		return tx.Update().Create(msg).Error
	})
}

// Save message
func (m *msgStore) Save(ctx context.Context, msg *core.Message) error {
	return m.db.Tx(func(tx *util.DB) error {
		var rows int64
		var err error
		rows, err = updateByphone(m.db, msg)
		if err != nil {
			return err
		}
		if rows == 0 {
			return tx.Update().Create(msg).Error
		}
		return nil
	})
}

// Find message by message_id
func (m *msgStore) Find(_ context.Context, message_id uint, phone string) (*core.Message, error) {
	msg := core.Message{}
	if err := m.db.View().Where("id = ? AND phone = ?", message_id, phone).Take(&msg).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &msg, nil
}
