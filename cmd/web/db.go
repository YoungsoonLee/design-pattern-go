package main

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	masxOpenDbConn = 25
	maxIdleDbConn  = 25
	maxDbLifetime  = 5 * time.Minute
)

func initMySQL(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// test our connection
	if err = db.Ping(); err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(masxOpenDbConn)
	db.SetMaxIdleConns(maxIdleDbConn)
	db.SetConnMaxLifetime(maxDbLifetime)

	return db, nil
}
