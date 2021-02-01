package repository

import (
	"github.com/Condition17/fleet-services/user-service/model"
	"gorm.io/gorm"
	"math/rand"
)

type UserRepository struct {
	Repository

	DB *gorm.DB
}

var PictureUrls = []string{
	"https://storage.googleapis.com/assets-pics/profile_1.png",
	"https://storage.googleapis.com/assets-pics/profile_2.png",
	"https://storage.googleapis.com/assets-pics/profile_3.png",
}

func (r *UserRepository) Create(user *model.User) error {
	user.Picture = PictureUrls[rand.Intn(len(PictureUrls))]
	if err := r.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) GetByEmail(email string) (*model.User, error) {
	user := &model.User{}
	if err := r.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) Get(id uint) (*model.User, error) {
	var user *model.User
	user.ID = id
	if err := r.DB.Find(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
