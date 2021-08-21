package mysql

import (
	"database/sql"
)

type connection interface {
	Prepare(query string) (*sql.Stmt, error)
}
