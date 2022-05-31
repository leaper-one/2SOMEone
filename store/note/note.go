package note

import (
	"context"
	"errors"

	"github.com/leaper-one/2SOMEone/core"
	"github.com/leaper-one/2SOMEone/util"

	"gorm.io/gorm"
)

func New(db *util.DB) core.NoteStore {
	return &noteStore{db: db}
}

type noteStore struct {
	db *util.DB
}

func toUpdateParams(note *core.Note) map[string]interface{} {
	return map[string]interface{}{
		"note_id":   note.NoteID,
		"type":      note.Type,
		"title":     note.Title,
		"content":   note.Content,
		"atts":      note.Atts,
		"sender":    note.Sender,
		"recipient": note.Recipient,
		"read":      note.Read,
		"archived":  note.Archived,
	}
}

func update(db *util.DB, note *core.Note) (int64, error) {
	updates := toUpdateParams(note)
	tx := db.Update().Model(note).Where("note_id = ?", note.NoteID).Updates(updates)
	return tx.RowsAffected, tx.Error
}

func (s *noteStore) Save(_ context.Context, note *core.Note) error {
	return s.db.Tx(func(tx *util.DB) error {
		rows, err := update(tx, note)
		if err != nil {
			return err
		}

		if rows == 0 {
			return tx.Update().Create(note).Error
		}

		return nil
	})
}

func (s *noteStore) FindByNoteID(ctx context.Context, note_id string) (*core.Note, error) {
	note := core.Note{}
	if err := s.db.View().Where("note_id = ?", note_id).Take(&note).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &note, nil
}

func (s *noteStore) DeleteByNoteID(ctx context.Context, note_id string) error {
	if err := s.db.Update().Where("note = ?", note_id).Delete(&core.Note{}).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	return nil
}

// 指定查询类型，需要提前获得userID
func (s *noteStore) GetNotes(ctx context.Context, offset, limit int64, is_sender bool, user_id string) ([]*core.Note, int64, error) {
	var notedb *gorm.DB
	if is_sender {
		notedb = s.db.View().Where("sender = ?", user_id).Model(&core.Note{})
	} else {
		notedb = s.db.View().Where("recipient = ?", user_id).Model(&core.Note{})
	}
	var count int64
	notedb.Count(&count) //总行数

	notes := []*core.Note{}
	// notedb.Offset((offset - 1) * limit).Limit(limit).Find(&notes)
	notedb.Offset(int((offset - 1) * limit)).Limit(int(limit)).Find(&notes)
	return notes, count, nil
}
