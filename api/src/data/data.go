package data

import (
	"DevBookAPI/src/config"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func Connecting() (*sql.DB, error) {
	db, err := sql.Open("mysql", config.ConnectionString)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
