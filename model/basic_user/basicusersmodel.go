package basic_user

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ BasicUsersModel = (*customBasicUsersModel)(nil)

type (
	// BasicUsersModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBasicUsersModel.
	BasicUsersModel interface {
		basicUsersModel
	}

	customBasicUsersModel struct {
		*defaultBasicUsersModel
	}
)

// NewBasicUsersModel returns a model for the database table.
func NewBasicUsersModel(conn sqlx.SqlConn) BasicUsersModel {
	return &customBasicUsersModel{
		defaultBasicUsersModel: newBasicUsersModel(conn),
	}
}
