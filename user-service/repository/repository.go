package repository

import proto "github.com/Condition17/fleet-services/user-service/proto/user-service"

type Repository interface {
	Get(id string) (*proto.User, error)
	Create(user *proto.User) error
}
