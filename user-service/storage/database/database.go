package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func CreateConnection() (*gorm.DB, error) {
	// host := os.Getenv("DB_HOST")
	// user := os.Getenv("DB_USER")
	// dbName := os.Getenv("DB_NAME")
	// password := os.Getenv("DB_PASSWORD")
	dsn := fmt.Sprintf("host=%s user=%s dbname=%s password=%s sslmode=disable", "127.0.0.1", "postgres", "fleet", "root")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
