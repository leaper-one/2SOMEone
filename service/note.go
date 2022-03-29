package service

import (
	"2SOMEone/core"
	"2SOMEone/store/note"
	"2SOMEone/store/user"
	"2SOMEone/util"
	"context"
	"errors"

	"github.com/gofrs/uuid"
)

func NewNoteService(
	db *util.DB,
) *NoteService {
	return &NoteService{
		db: db,
	}
}

type NoteService struct {
	db *util.DB
}

func (n *NoteService) Create(ctx context.Context, tnote *core.Note, recipient_name string) error {
	noteStore := note.New(n.db)
	userStore := user.New(n.db)
	ruser, err := userStore.FindByName(ctx, recipient_name)
	if err != nil {
		return err
	} else if ruser == nil && err == nil {
		return errors.New("无此用户")
	}
	tnote.Recipient = ruser.UserID
	note_id, _ := uuid.NewV1()
	tnote.NoteID = note_id.String()

	err = noteStore.Save(ctx, tnote)
	if err != nil {
		return err
	}
	return nil
}

func (n *NoteService) Delete(ctx context.Context, note_id string) error {
	noteStore := note.New(n.db)
	err := noteStore.DeleteByNoteID(ctx, note_id)
	if err != nil {
		return err
	}
	return nil
}
func (n *NoteService) SenderGet(ctx context.Context, offset, limit int, user_id string) ([]*core.Note, int64, error) {
	noteStore := note.New(n.db)
	notes, count, err := noteStore.GetSent(ctx, offset, limit, user_id)
	if err != nil {
		return nil, 0, err
	} else if notes == nil && err == nil {
		return nil, 0, errors.New("无记录")
	}
	return notes, count, nil
}

func (n *NoteService) RecipientGet(ctx context.Context, offset, limit int, user_id string) ([]*core.Note, int64, error) {
	noteStore := note.New(n.db)
	notes, count, err := noteStore.GetReceived(ctx, offset, limit, user_id)
	if err != nil {
		return nil, 0, err
	} else if notes == nil && err == nil {
		return nil, 0, errors.New("无记录")
	}
	return notes, count, nil
}

func (n *NoteService) GetByID(ctx context.Context, note_id string) (*core.Note, error) {
	noteStore := note.New(n.db)
	userStore := user.New(n.db)
	note, err := noteStore.FindByNoteID(ctx, note_id)
	if err != nil {
		return nil, err
	} else if note == nil && err == nil {
		return nil, errors.New("无此 Note")
	}

	// 隐藏 user_id
	sender, err := userStore.FindByUserID(ctx, note.Sender)
	if err != nil {
		return nil, err
	} else if sender == nil && err == nil {
		return nil, errors.New("无发送者")
	}

	recipient, err := userStore.FindByUserID(ctx, note.Recipient)
	if err != nil {
		return nil, err
	} else if sender == nil && err == nil {
		return nil, errors.New("无接收者")
	}

	note.Sender = sender.Name
	note.Recipient = recipient.Name

	return note, nil
}
