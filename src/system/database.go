package system

import (
	"database/sql"
	"log"
)

var DB *sql.DB

func InitDB(dsn string) (*sql.DB, error) {

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)

	if err := db.Ping(); err != nil {
		return nil, err
	}

	DB = db

	log.Println("Connected to database")

	return db, nil
}
