package mysql

import (
	"database/sql"
	"fmt"
	"sync"

	"github.com/GuiaBolso/darwin"
	"github.com/diegoclair/bank-transfer/contract"
	"github.com/diegoclair/bank-transfer/infra/data/migrations"
	"github.com/diegoclair/bank-transfer/util/config"
	"github.com/diegoclair/go_utils-lib/v2/logger"
	mysqlDriver "github.com/go-sql-driver/mysql"
)

var (
	conn    *mysqlConn
	onceDB  sync.Once
	connErr error
)

type mysqlConn struct {
	db *sql.DB
}

//Instance returns an instance of a MySQLRepo
func Instance() (contract.MySQLRepo, error) {
	onceDB.Do(func() {
		cfg := config.GetConfigEnvironment()

		dataSourceName := fmt.Sprintf("%s:root@tcp(%s:%s)/%s?charset=utf8",
			cfg.MySQL.Username, cfg.MySQL.Host, cfg.MySQL.Port, cfg.MySQL.DBName,
		)

		logger.Info("Connecting to database...")
		db, connErr := sql.Open("mysql", dataSourceName)
		if connErr != nil {
			return
		}

		logger.Info("Database Ping...")
		connErr = db.Ping()
		if connErr != nil {
			return
		}

		logger.Info("Creating database...")
		if _, connErr = db.Exec("CREATE DATABASE IF NOT EXISTS sampamodas_db;"); connErr != nil {
			logger.Error("Create Database error: ", connErr)
			return
		}

		if _, connErr = db.Exec("USE sampamodas_db;"); connErr != nil {
			logger.Error("Default Database error: ", connErr)
			return
		}

		connErr = mysqlDriver.SetLogger(logger.GetLogger())
		if connErr != nil {
			return
		}
		logger.Info("Database successfully configured")

		logger.Info("Running the migrations")
		driver := darwin.NewGenericDriver(db, darwin.MySQLDialect{})
		d := darwin.New(driver, migrations.Migrations, nil)

		connErr = d.Migrate()
		if connErr != nil {
			logger.Error("Migrate Error: ", connErr)
			return
		}

		logger.Info("Migrations executed")

		conn = &mysqlConn{
			db: db,
		}
	})

	return conn, connErr
}

// Begin starts a mysql transaction
func (c *mysqlConn) Begin() (contract.MysqlTransaction, error) {
	tx, err := c.db.Begin()
	if err != nil {
		return nil, err
	}

	return newTransaction(tx), nil
}

func (c *mysqlConn) Close() (err error) {
	return c.db.Close()
}

func (c *mysqlConn) Auth() contract.AuthRepo {
	return newAuthRepo(c.db)
}
