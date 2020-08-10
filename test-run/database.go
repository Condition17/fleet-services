package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"os"
)

func CreateConnection() (*gorm.DB, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASSWORD")
	conn := fmt.Sprintf("host=%s user=%s dbname=%s password=%s sslmode=disable", host, user, dbName, password)
	db, err := gorm.Open("postgres", conn)

	if err != nil {
		return nil, err
	}

	return db, nil
}