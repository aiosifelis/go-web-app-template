package system

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	connectionPool *sql.DB
}

func CreateConnectionPool() (Database, error) {

	// Create connection pool
	dsn, ok := os.LookupEnv("DATABASE_DSN")
	if !ok {
		panic("DATABASE_DSN is not set")
	}

	connectionPool, err := sql.Open("mysql", dsn)
	if err != nil {
		return Database{}, err
	}

	connectionPool.SetMaxOpenConns(10)
	connectionPool.SetMaxIdleConns(5)

	if err := connectionPool.Ping(); err != nil {
		return Database{}, err
	}

	fmt.Println("Database connection established")

	return Database{connectionPool: connectionPool}, nil
}
