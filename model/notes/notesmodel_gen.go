// Code generated by goctl. DO NOT EDIT!

package notes

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	notesFieldNames          = builder.RawFieldNames(&Notes{})
	notesRows                = strings.Join(notesFieldNames, ",")
	notesRowsExpectAutoSet   = strings.Join(stringx.Remove(notesFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), ",")
	notesRowsWithPlaceHolder = strings.Join(stringx.Remove(notesFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), "=?,") + "=?"
)

type (
	notesModel interface {
		Insert(ctx context.Context, data *Notes) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Notes, error)
		Update(ctx context.Context, data *Notes) error
		Delete(ctx context.Context, id int64) error
	}

	defaultNotesModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Notes struct {
		Id        int64          `db:"id"`
		CreatedAt sql.NullTime   `db:"created_at"`
		UpdatedAt sql.NullTime   `db:"updated_at"`
		DeletedAt sql.NullTime   `db:"deleted_at"`
		NoteId    sql.NullString `db:"note_id"`
		Type      int64          `db:"type"`
		Title     sql.NullString `db:"title"`
		Content   sql.NullString `db:"content"`
		Atts      sql.NullString `db:"atts"`
		Sender    sql.NullString `db:"sender"`
		Recipient sql.NullString `db:"recipient"`
		Read      int64          `db:"read"`
		Archived  int64          `db:"archived"`
	}
)

func newNotesModel(conn sqlx.SqlConn) *defaultNotesModel {
	return &defaultNotesModel{
		conn:  conn,
		table: "`notes`",
	}
}

func (m *defaultNotesModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultNotesModel) FindOne(ctx context.Context, id int64) (*Notes, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", notesRows, m.table)
	var resp Notes
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultNotesModel) Insert(ctx context.Context, data *Notes) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, notesRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.CreatedAt, data.UpdatedAt, data.DeletedAt, data.NoteId, data.Type, data.Title, data.Content, data.Atts, data.Sender, data.Recipient, data.Read, data.Archived)
	return ret, err
}

func (m *defaultNotesModel) Update(ctx context.Context, data *Notes) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, notesRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.CreatedAt, data.UpdatedAt, data.DeletedAt, data.NoteId, data.Type, data.Title, data.Content, data.Atts, data.Sender, data.Recipient, data.Read, data.Archived, data.Id)
	return err
}

func (m *defaultNotesModel) tableName() string {
	return m.table
}
