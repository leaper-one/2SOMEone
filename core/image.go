package core

import (
	"context"
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type (
	Image struct {
		gorm.Model
		ImageName    string          `json:"image_name,omitempty"`
		Fix          string          `json:"fix,omitempty"`
		Size         decimal.Decimal `gorm:"type:float" json:"size"`
		AttachmentID string          `gorm:"index" json:"attachment_id"`
		IsLocal      bool            `gorm:"default:false" json:"is_local"`
		ViewUrl      string          `json:"view_url"`
		UserID       string          `gorm:"size:36;unique_index;index" json:"mixin_id,omitempty"`
	}

	ImageForUser struct {
		CreatedAt    time.Time `json:"created_at"`
		ImageName    string    `json:"image_name,omitempty"`
		AttachmentID string    `gorm:"index" json:"attachment_id"`
	}

	ImageStore interface {
		Save(ctx context.Context, img *Image) error
		Find(ctx context.Context, img_id string) (*Image, error)
		GetImgs(ctx context.Context, offset, limit int, user_id string) ([]*ImageForUser, int64, error)
		GetImg(ctx context.Context, img_id, user_id string) (*Image, error)
		DeleteByImgID(ctx context.Context, img_id string) error
		DeleteByUserID(ctx context.Context, user_id string) error
		HasImgByID(ctx context.Context, img_id string) (bool, *Image, error)
	}

	ImageService interface {
		Find(ctx context.Context, imgID string) (*Image, error)
		Save(ctx context.Context, img []byte) error
	}
)
