package gorose

import (
	"context"
	"database/sql"
)

// ISession ...
type ISession interface {
	Close()
	//Table(bind interface{}) IOrm
	Bind(bind interface{}) ISession
	BeginTx(ctx context.Context, opts *sql.TxOptions) (err error)
	Begin() (err error)
	Rollback() (err error)
	Commit() (err error)
	Transaction(closer ...func(session ISession) error) (err error)
	Query(sqlstring string, args ...interface{}) ([]Data, error)
	Execute(sqlstring string, args ...interface{}) (int64, error)
	//GetDriver() string
	GetSessionId() uint64
	GetIEngin() IEngin
	LastInsertId() int64
	LastSql() string
	//SetIBinder(b IBinder)
	GetTableName() (string, error)
	SetIBinder(ib IBinder)
	GetIBinder() IBinder
	SetUnion(u interface{})
	GetUnion() interface{}
	SetTransaction(b bool)
	GetTransaction() bool
	//ResetBinder()
	GetBindAll() []Data
	GetErr() error
}
