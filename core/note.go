package core

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type (
	Note struct {
		gorm.Model
		NoteID    string `gorm:"size:36;unique_index"`
		Type      int8   `gorm:"default:0"`
		Context   string `gorm:"size:140"`
		AttID     []string
		Imgs      []string
		Sender    string `gorm:"size:36;unique_index"`
		Recipient string `gorm:"size:36;unique_index"`
	}

	NoteStore interface {
		Save(ctx context.Context, note *Note) error
		// SaveByEmail(_ context.Context, user *User) error
		// FindByMixinID(ctx context.Context, mixinID string) (*User, error)
		FindByNoteID(ctx context.Context, note_id string) (*Note, error)
		// FindByPhone(ctx context.Context, phone string) (*User, error)
		// FindByEmail(ctx context.Context, email string) (*User, error)
		// FindByName(ctx context.Context, phone string) (*User, error)
		// FindByUserID(_ context.Context, user_id string) (*User, error)
		DeleteByNoteID(ctx context.Context, note_id string) error
		GetNotes(ctx context.Context, offset, limit int, sender, recipient string) ([]*Note, int64, error)
	}

	NoteService interface {
		Create(ctx context.Context, tnote *Note, note_type int8, recipient_name string) error
		Delete(ctx context.Context, note_id string) error
		// Get(ctx context.Context, offset, limit int, user_id string) ([]*Note, int64, error)
		SenderGet(ctx context.Context, offset, limit int, user_id string) ([]*Note, int64, error)
		RecipientGet(ctx context.Context, offset, limit int, user_id string) ([]*Note, int64, error)
	}
)
