package db

import (
	"database/sql"
	"go_social/config"

	_ "github.com/go-sql-driver/mysql" // MySQL driver
)

func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", config.DatabaseStringConection)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}
