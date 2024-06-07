package database

import (
	"database/sql"
	"heartvoice/src/config"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
	db, connectError := sql.Open("mysql", config.StrConn)

	if connectError != nil {
		return nil, connectError
	}

	if connectError = db.Ping(); connectError != nil {
		db.Close()
		return nil, connectError
	}

	return db, nil
}
