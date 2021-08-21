package contract

//MySQLRepo defines the repository aggregator interface
type MySQLRepo interface {
	Begin() (MysqlTransaction, error)
	Auth() AuthRepo
}

// MysqlTransaction holds the methods that manipulates the main data, from within a transaction.
type MysqlTransaction interface {
	MySQLRepo
	Rollback() error
	Commit() error
}

type AuthRepo interface {
}
