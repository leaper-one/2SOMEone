package service

import (
	"2Some/core"
	"2Some/store/note"
	"2Some/store/user"
	"2Some/util"
	"context"

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

func (n *NoteService) Create(ctx context.Context, tnote *core.Note, note_type int8, recipient_name string) error {
	noteStore := note.New(n.db)
	userStore := user.New(n.db)
	ruser, err := userStore.FindByName(ctx, recipient_name)
	if err != nil || ruser.ID == 0 {
		return err
	}
	tnote.Recipient = ruser.UserID
	tnote.Type = note_type
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
	notes,count,err:=noteStore.GetNotes(ctx, offset,limit,user_id,"")
	if err != nil {
		return nil,0,err
	}
	return notes, count, nil
}

func (n *NoteService) RecipientGet(ctx context.Context, offset, limit int, user_id string) ([]*core.Note, int64, error) {
	noteStore := note.New(n.db)
	notes,count,err:=noteStore.GetNotes(ctx, offset,limit,"",user_id)
	if err != nil {
		return nil,0,err
	}
	return notes, count, nil
}
