package contract

import "github.com/diegoclair/bank-transfer/domain/entity"

// MysqlTransaction holds the methods that manipulates the main data, from within a transaction.
type MysqlTransaction interface {
	MySQLRepo
	Rollback() error
	Commit() error
}

//MySQLRepo defines the repository aggregator interface
type MySQLRepo interface {
	Begin() (MysqlTransaction, error)
	User() UserRepo
}

type UserRepo interface {
	CreateUser(user entity.User) (err error)
	GetUserByDocument(encryptedDocumentNumber string) (user entity.User, err error)
}
