package bili_user

import (
	"context"
	"errors"

	"github.com/leaper-one/2SOMEone/core"
	"github.com/leaper-one/2SOMEone/util"

	"gorm.io/gorm"
)

func NewBiliUserStore(db *util.DB) core.BiliUserStore {
	return &biliUserStore{db: db}
}

type biliUserStore struct {
	db *util.DB
}

func toUpdateParams(user *core.BiliUser) map[string]interface{} {
	return map[string]interface{}{
		"user_id":       user.UserID,
		"buid":          user.Buid,
		"live_room_id":  user.LiveRoomID,
		"live_room_url": user.LiveRoomUrl,
	}
}

func update(db *util.DB, user *core.BiliUser) (int64, error) {
	updates := toUpdateParams(user)
	tx := db.Update().Model(user).Where("user_id = ?", user.UserID).Updates(updates)
	return tx.RowsAffected, tx.Error
}

func (s *biliUserStore) Save(_ context.Context, bili_user *core.BiliUser) error {
	return s.db.Tx(func(tx *util.DB) error {
		var rows int64
		var err error

		rows, err = update(tx, bili_user)
		if err != nil {
			return err
		}

		if rows == 0 {
			return tx.Update().Create(bili_user).Error
		}

		return nil
	})
}

func (s *biliUserStore) FindByUserID(_ context.Context, user_id string) (*core.BiliUser, error) {
	user := core.BiliUser{}
	if err := s.db.View().Where("user_id = ?", user_id).Take(&user).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &user, nil
}
