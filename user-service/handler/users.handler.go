package handler

import (
	"context"
	"errors"
	"fmt"

	"gorm.io/gorm"

	"github.com/Condition17/fleet-services/user-service/model"
	proto "github.com/Condition17/fleet-services/user-service/proto/user-service"
	microErrors "github.com/micro/go-micro/v2/errors"
	"github.com/micro/go-micro/v2/metadata"
	"golang.org/x/crypto/bcrypt"
)

func (s *Service) Create(ctx context.Context, req *proto.User, res *proto.AuthResponse) error {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return microErrors.BadRequest(s.Name, "Unapropriate password format.")
	}

	// Check if user already exists
	if user, _ := s.UserRepository.GetByEmail(req.Email); user != nil {
		return microErrors.Conflict(s.Name, "An user with this email was already created. Please try another email.")
	}

	// Create user
	req.Password = string(hashedPass)
	if err := s.UserRepository.Create(model.MarshalUser(req)); err != nil {
		return microErrors.InternalServerError(s.Name, fmt.Sprintf("%v", err))
	}
	res.User = req

	// Create JWT token
	if res.Token, err = s.generateToken(req); err != nil {
		return microErrors.InternalServerError(s.Name, fmt.Sprintf("%v", err))
	}

	return nil
}

func (s *Service) Authenticate(ctx context.Context, req *proto.User, res *proto.AuthResponse) error {
	user, err := s.UserRepository.GetByEmail(req.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return microErrors.Unauthorized(s.Name, "User with this email not found")
		}
		return microErrors.InternalServerError(s.Name, fmt.Sprintf("%v", err))
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return microErrors.Unauthorized(s.Name, "Invalid password.")
	}

	if res.Token, err = s.generateToken(req); err != nil {
		return microErrors.InternalServerError(s.Name, fmt.Sprintf("%v", err))
	}

	return nil
}

func (s *Service) ValidateToken(ctx context.Context, req *proto.Token, res *proto.TokenValidationResponse) error {
	_, err := s.TokenService.Decode(req.Token)

	if err != nil {
		res.Valid = false
		return err
	} else {
		res.Valid = true
	}

	return nil
}

func (s *Service) GetProfile(ctx context.Context, req *proto.EmptyRequest, res *proto.AuthResponse) error {
	meta, ok := metadata.FromContext(ctx)
	if !ok {
		return microErrors.InternalServerError(s.Name, "Could not parse request headers (context metadata).")
	}
	token := meta["Token"]
	claims, err := s.TokenService.Decode(token)

	if err != nil {
		return microErrors.Unauthorized(s.Name, "Invalid token")
	}
	res.User = claims.User

	return nil
}

func (s *Service) generateToken(payload *proto.User) (*proto.Token, error) {
	token, err := s.TokenService.Encode(payload)
	if err != nil {
		return nil, err
	}

	return &proto.Token{Token: token}, nil
}
