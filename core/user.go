package core

import (
	"context"

	"gorm.io/gorm"
)

type (
	User struct {
		gorm.Model
		UserID      string `gorm:"size:36;unique_index"`
		Name        string `gorm:"size:64; unique_index" json:"name,omitempty"`
		Phone       string `gorm:"size:14;index" json:"phone,omitempty"`
		Email       string `gorm:"unique_index" json:"email,omitempty"`
		Code        string `gorm:"size:6"`
		Password    string `gorm:"size:20" json:"password,omitempty"`
		Buid        string `json:"buid,omitempty"`
		MixinID     string `gorm:"size:36;unique_index" json:"mixin_id,omitempty"`
		Role        string `gorm:"size:24" json:"role,omitempty"`
		Lang        string `gorm:"size:36;default:'zh'" json:"lang,omitempty"`
		Avatar      string `gorm:"size:255" json:"avatar,omitempty"`
		AccessToken string `gorm:"size:512" json:"access_token,omitempty"`
		// Balence     decimal.Decimal `gorm:"precision:2"`
	}

	UserForShow struct {
		Name   string `gorm:"size:64; unique_index" json:"name,omitempty"`
		Avatar string `gorm:"size:255" json:"avatar,omitempty"`
	}

	UserForMe struct {
		Name   string `gorm:"size:64; unique_index" json:"name,omitempty"`
		Phone  string `gorm:"size:14;index" json:"phone,omitempty"`
		Email  string `gorm:"unique_index" json:"email,omitempty"`
		Buid   string `json:"buid,omitempty"`
		Role   string `gorm:"size:24" json:"role,omitempty"`
		Lang   string `gorm:"size:36;default:'zh'" json:"lang,omitempty"`
		Avatar string `gorm:"size:255" json:"avatar,omitempty"`
	}

	LoginUser struct {
		Phone string `json:"phone"`
		Code  string `json:"code,omitempty"`
		// Email    string `json:"email,omitempty"`
		Password string `json:"password,omitempty"`
	}

	SignUpUser struct {
		Phone    string `json:"phone,omitempty"`
		Code     string `json:"code,omitempty"`
		Password string `json:"password"`
	}

	UserStore interface {
		Save(ctx context.Context, user *User) error
		// SaveByEmail(_ context.Context, user *User) error
		FindByMixinID(ctx context.Context, mixinID string) (*User, error)
		FindByPhone(ctx context.Context, phone string) (*User, error)
		// FindByEmail(ctx context.Context, email string) (*User, error)
		FindByName(ctx context.Context, neme string) (*User, error)
		FindByUserID(_ context.Context, user_id string) (*User, error)
		FindByUserIDForShow(_ context.Context, user_id string) (*UserForShow, error)
		FindByUserIDForMe(_ context.Context, user_id string) (*UserForMe, error)
		DeleteByUserID(_ context.Context, email string) error
	}

	UserService interface {
		GetPhoneCode(ctx context.Context, phone string) error
		SignUpByPhone(ctx context.Context, l_user *SignUpUser) (*User, error)
		// Login(ctx context.Context, token string) (*User, error)
		Auth(ctx context.Context, login_user *LoginUser) (string, error)
		GetMe(ctx context.Context, user_id string) (*User, error)
	}
)
