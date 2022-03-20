package core

import (
	"context"

	"gorm.io/gorm"
)

type (
	User struct {
		gorm.Model
		UserID      string `gorm:"size:36;unique_index"`
		Password    string `gorm:"size:20" json:"password,omitempty"`
		MixinID     string `gorm:"size:36;unique_index" json:"mixin_id,omitempty"`
		Role        string `gorm:"size:24" json:"role,omitempty"`
		Lang        string `gorm:"size:36;default:'zh'" json:"lang,omitempty"`
		Name        string `gorm:"size:64; unique_index" json:"name,omitempty"`
		Buid        string `json:"buid,omitempty"`
		Avatar      string `gorm:"size:255" json:"avatar,omitempty"`
		AccessToken string `gorm:"size:512" json:"access_token,omitempty"`
		Phone       string `gorm:"size:14;index" json:"phone,omitempty"`
		Code        string `gorm:"size:6"`
		// Balence     decimal.Decimal `gorm:"precision:2"`
		// Email       string          `gorm:"unique_index" json:"email,omitempty"`
	}

	LoginUser struct {
		Phone string `json:"phone,omitempty"`
		Code  string `json:"code,omitempty"`
		// Email    string `json:"email,omitempty"`
		Password string `json:"password"`
	}

	SignUp struct {
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
		DeleteByUserID(_ context.Context, email string) error
	}

	UserService interface {
		GetPhoneCode(ctx context.Context, phone string) error
		SignUp(ctx context.Context, l_user *SignUp) (*User, string, error)
		// Login(ctx context.Context, token string) (*User, error)
		Auth(ctx context.Context, login_user *LoginUser) (string, error)
	}
)
