package database

import (
	"fmt"

	"github.com/Condition17/fleet-services/user-service/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func CreateConnection() (*gorm.DB, error) {
	configs := config.GetConfig()
	host := configs.DbHost
	user := configs.DbUser
	dbName := configs.DbName
	password := configs.DbPassword

	dsn := fmt.Sprintf("host=%s user=%s dbname=%s password=%s sslmode=disable", host, user, dbName, password)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
