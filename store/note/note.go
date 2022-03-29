package note

import (
	"2SOMEone/core"
	"2SOMEone/util"
	"context"
	"errors"

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
		"context":   note.Context,
		"imgs":      note.Imgs,
		"atts":      note.Atts,
		"sender":    note.Sender,
		"recipient": note.Recipient,
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
func (s *noteStore) GetNotes(ctx context.Context, offset, limit int, sender, recipient string) ([]*core.Note, int64, error) {
	var notedb *gorm.DB
	if recipient != "" && sender != "" {
		notedb = s.db.View().Where(&core.Note{Sender: sender, Recipient: recipient}).Model(&core.Note{})
	} else if recipient == "" && sender != "" {
		notedb = s.db.View().Where(&core.Note{Sender: sender}).Model(&core.Note{})
	} else if recipient != "" && sender == "" {
		notedb = s.db.View().Where(&core.Note{Recipient: recipient}).Model(&core.Note{})
	} else if recipient == "" && sender == "" {
		return nil, 0, errors.New("未指定对象")
	}
	var count int64
	notedb.Count(&count) //总行数

	notes := []*core.Note{}
	notedb.Offset((offset - 1) * limit).Limit(limit).Find(&notes)

	return notes, count, nil
}

func (s *noteStore) GetReceived(ctx context.Context, offset, limit int, user_id string) ([]*core.Note, int64, error) {
	notedb := s.db.View().Where(&core.Note{Recipient: user_id}).Model(&core.Note{})
	if notedb.Error != nil {
		return nil, 0, notedb.Error
	} else if errors.Is(notedb.Error, gorm.ErrRecordNotFound) {
		return nil, 0, nil
	}
	var count int64
	notedb.Count(&count) //总行数

	notes := []*core.Note{}
	err := notedb.Offset((offset - 1) * limit).Limit(limit).Find(&notes).Error
	if err != nil {
		return nil, 0, err
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, 0, nil
	}

	for _, note := range notes {
		note.ForRead()
	}
	return notes, count, nil
}

func (s *noteStore) GetSent(ctx context.Context, offset, limit int, user_id string) ([]*core.Note, int64, error) {
	notedb := s.db.View().Where(&core.Note{Sender: user_id}).Model(&core.Note{})
	if notedb.Error != nil {
		return nil, 0, notedb.Error
	} else if errors.Is(notedb.Error, gorm.ErrRecordNotFound) {
		return nil, 0, nil
	}
	var count int64
	notedb.Count(&count) //总行数

	notes := []*core.Note{}
	err := notedb.Offset((offset - 1) * limit).Limit(limit).Find(&notes).Error
	if err != nil {
		return nil, 0, err
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, 0, nil
	}

	for _, note := range notes {
		note.ForRead()
	}

	return notes, count, nil
}
