package repository

import "gorm.io/gorm"

type FileSystemRepository struct {
	Repository

	DB *gorm.DB
}
