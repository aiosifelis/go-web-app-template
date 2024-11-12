package system

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitializeDatabase() (*gorm.DB, error) {

	// Create connection pool
	dsn := getDSN()

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// Set connection pool
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(5 * 60 * 60)

	return db, nil
}

func getDSN() string {
	dbUser, ok := os.LookupEnv("DB_USER")
	if !ok {
		panic("DB_USER is not set")
	}
	dbPass, ok := os.LookupEnv("DB_PASS")
	if !ok {
		panic("DB_PASS is not set")
	}
	dbHost, ok := os.LookupEnv("DB_HOST")
	if !ok {
		panic("DB_HOST is not set")
	}
	dbName, ok := os.LookupEnv("DB_NAME")
	if !ok {
		panic("DB_NAME is not set")
	}

	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName)
}
