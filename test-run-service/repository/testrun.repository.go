package repository

import "gorm.io/gorm"

type TestRunRepository struct {
	Repository

	DB *gorm.DB
}
