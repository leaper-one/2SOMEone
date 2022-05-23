package core

import (
	"context"

	"gorm.io/gorm"
)

type (
	BiliUser struct {
		gorm.Model
		UserID      string `gorm:"size:36;unique_index"`
		Buid        int64  `json:"buid,omitempty"`
		LiveRoomID  int64  `json:"live_room_id,omitempty"`
		LiveRoomUrl string `json:"live_room_url,omitempty"`
	}

	BiliUserStore interface {
		Save(ctx context.Context, bili_user *BiliUser) error
		FindByUserID(ctx context.Context, user_id string) (*BiliUser, error)
	}
)
