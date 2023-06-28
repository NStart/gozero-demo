package mysql_nocache

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ OsOrgInfoModel = (*customOsOrgInfoModel)(nil)

type (
	// OsOrgInfoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOsOrgInfoModel.
	OsOrgInfoModel interface {
		osOrgInfoModel
	}

	customOsOrgInfoModel struct {
		*defaultOsOrgInfoModel
	}
)

// NewOsOrgInfoModel returns a model for the database table.
func NewOsOrgInfoModel(conn sqlx.SqlConn) OsOrgInfoModel {
	return &customOsOrgInfoModel{
		defaultOsOrgInfoModel: newOsOrgInfoModel(conn),
	}
}
