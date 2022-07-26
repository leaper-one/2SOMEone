package bili_user

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ BiliUsersModel = (*customBiliUsersModel)(nil)

type (
	// BiliUsersModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBiliUsersModel.
	BiliUsersModel interface {
		biliUsersModel
	}

	customBiliUsersModel struct {
		*defaultBiliUsersModel
	}
)

// NewBiliUsersModel returns a model for the database table.
func NewBiliUsersModel(conn sqlx.SqlConn) BiliUsersModel {
	return &customBiliUsersModel{
		defaultBiliUsersModel: newBiliUsersModel(conn),
	}
}
