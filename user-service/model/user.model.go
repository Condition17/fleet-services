package model

import (
	proto "github.com/Condition17/fleet-services/user-service/proto/user-service"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"not null"`
	Email    string `gorm:"unique;uniqueIndex"`
	Company  string `gorm:"not null"`
	Password string `gorm:"not null"`
	Picture  string `gorm:"default:null"`
}

func MarshalUser(user *proto.User) *User {
	return &User{
		Name:     user.Name,
		Email:    user.Email,
		Company:  user.Company,
		Password: user.Password,
		Picture:  user.Picture,
	}
}

func UnmarshalUser(user *User) *proto.User {
	return &proto.User{
		Id:       uint32(user.ID),
		Name:     user.Name,
		Email:    user.Email,
		Company:  user.Company,
		Password: "",
		Picture:  user.Picture,
	}
}
