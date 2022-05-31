package core

import (
	"context"

	"gorm.io/gorm"
)

type (
	Note struct {
		gorm.Model
		NoteID    string `gorm:"size:36;unique_index"`
		Type      int8   `gorm:"default:0"`
		Title     string `gorm:"size:20"`
		Content   string `gorm:"size:255"`
		Atts      string
		Sender    string `gorm:"size:36;unique_index"`
		Recipient string `gorm:"size:36;unique_index"`
		Read      bool   `gorm:"default:false"`
		Archived  bool   `gorm:"default:false"`
	}

	NoteStore interface {
		Save(ctx context.Context, note *Note) error
		FindByNoteID(ctx context.Context, note_id string) (*Note, error)
		DeleteByNoteID(ctx context.Context, note_id string) error
		GetNotes(ctx context.Context, offset, limit int64, is_sender bool, user_id string) ([]*Note, int64, error)
	}

	NoteService interface {
		Create(ctx context.Context, tnote *Note, recipient_name string) error
		Delete(ctx context.Context, note_id string) error
		GetNotes(ctx context.Context, offset, limit int64, is_sender bool, user_id string) ([]*Note, int64, error)
	}
)
