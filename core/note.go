package core

import (
	"context"
	"encoding/json"

	"gorm.io/gorm"
)

type (
	Note struct {
		gorm.Model
		NoteID    string   `gorm:"size:36;unique_index"`
		Type      int8     `gorm:"default:0"`
		Context   string   `gorm:"size:140"`
		AttsArry  []string `gorm:"-" json:"atts_arry,omitempty"`
		Atts      string
		ImgsArry  []string `gorm:"-" json:"imgs_arry,omitempty"`
		Imgs      string
		Sender    string `gorm:"size:36;unique_index"`
		Recipient string `gorm:"size:36;unique_index"`
	}

	NoteConvers interface {
		ForStore() *Note
		ForRead() *Note
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
		Create(ctx context.Context, tnote *Note, recipient_name string) error
		Delete(ctx context.Context, note_id string) error
		// Get(ctx context.Context, offset, limit int, user_id string) ([]*Note, int64, error)
		SenderGet(ctx context.Context, offset, limit int, user_id string) ([]*Note, int64, error)
		RecipientGet(ctx context.Context, offset, limit int, user_id string) ([]*Note, int64, error)
	}
)

func (n *Note) ForStore() *Note {
	if len(n.AttsArry) != 0 {
		atts, _ := json.Marshal(n.AttsArry)
		n.Atts = string(atts)
	}
	if len(n.ImgsArry) != 0 {
		imgs, _ := json.Marshal(n.ImgsArry)
		n.Imgs = string(imgs)
	}

	return n
}

func (n *Note) ForRead() (*Note, error) {
	if n.Atts != "" {
		err := json.Unmarshal([]byte(n.Atts), &n.AttsArry)
		if err != nil {
			return n, err
		}
		n.Atts = ""
	}
	if n.Imgs != "" {
		err := json.Unmarshal([]byte(n.Imgs), &n.ImgsArry)
		if err != nil {
			return n, err
		}
		n.Imgs = ""
	}

	return n, nil
}
