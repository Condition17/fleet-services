package repository

import (
	proto "github.com/Condition17/fleet-services/user-service/proto/user-service"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type UserRepository struct {
	Repository

	DB *gorm.DB
}

func (r *UserRepository) Create(user *proto.User) error {
	user.Id = uuid.NewV4().String()
	if err := r.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) GetByEmail(email string) (*proto.User, error) {
	user := &proto.User{}
	if err := r.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) Get(id string) (*proto.User, error) {
	var user *proto.User
	user.Id = id
	if err := r.DB.Find(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
