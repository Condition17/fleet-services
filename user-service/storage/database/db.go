package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func CreateConnection() (*gorm.DB, error) {
	// host := os.Getenv("DB_HOST")
	// user := os.Getenv("DB_USER")
	// dbName := os.Getenv("DB_NAME")
	// password := os.Getenv("DB_PASSWORD")
	conn := fmt.Sprintf("host=%s user=%s dbname=%s password=%s sslmode=disable", "127.0.0.1", "postgres", "fleet", "root")
	db, err := gorm.Open("postgres", conn)
	if err != nil {
		return nil, err
	}

	return db, nil
}
