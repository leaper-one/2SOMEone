package core

import (
	"context"

	"gorm.io/gorm"
)

type (
	BasicUser struct {
		gorm.Model
		UserID   string `gorm:"size:36;unique_index"`
		Name     string `gorm:"size:64; unique_index" json:"name,omitempty"`
		Phone    string `gorm:"size:14;index" json:"phone,omitempty"`
		Email    string `gorm:"unique_index" json:"email,omitempty"`
		Password string `gorm:"size:20" json:"password,omitempty"`
		Lang     string `gorm:"size:36;default:'zh'" json:"lang,omitempty"`
		Avatar   string `gorm:"size:255" json:"avatar,omitempty"`
		State    string `gorm:"size:24;default:'formal'" json:"state,omitempty"`
	}

	BiliUserInfo struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Data    struct {
			Mid      int    `json:"mid,omitempty"`
			Face     string `json:"face,omitempty"`
			Name     string `json:"name,omitempty"`
			LiveRoom struct {
				RoomID int    `json:"roomid,omitempty"`
				Url    string `json:"url,omitempty"`
			} `json:"live_room,omitempty"`
		} `json:"data,omitempty"`
	}

	UserStore interface {
		Save(ctx context.Context, user *BasicUser) error
		// SaveByEmail(_ context.Context, user *User) error
		FindByMixinID(ctx context.Context, mixinID string) (*BasicUser, error)
		FindByPhone(ctx context.Context, phone string) (*BasicUser, error)
		// FindByEmail(ctx context.Context, email string) (*User, error)
		FindByName(ctx context.Context, neme string) (*BasicUser, error)
		FindByUserID(_ context.Context, user_id string) (*BasicUser, error)
		DeleteByUserID(_ context.Context, email string) error
	}

	UserService interface {
		GetPhoneCode(ctx context.Context, phone string) error
		SignUpByPhone(ctx context.Context, phone, code, password string, msg_id uint) error
		// Login(ctx context.Context, token string) (*User, error)
		Auth(ctx context.Context, phone, email, password string) (string, error)
		SetInfo(ctx context.Context, email, buid, avatar string) error
		GetMe(ctx context.Context, user_id string) (*BasicUser, error)
		FindByBuid(ctx context.Context, buid int64) (*BiliUser, error)
	}
)
