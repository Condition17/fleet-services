package handler

import (
	"context"
	"fmt"

	proto "github.com/Condition17/fleet-services/user-service/proto/user-service"
	"github.com/micro/go-micro/v2/errors"
	"github.com/micro/go-micro/v2/metadata"
	"golang.org/x/crypto/bcrypt"
)

func (s *Service) Create(ctx context.Context, req *proto.User, res *proto.AuthResponse) error {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.BadRequest(s.Name, "Unapropriate password format.")
	}

	req.Password = string(hashedPass)
	if err := s.UserRepository.Create(req); err != nil {
		return errors.InternalServerError(s.Name, fmt.Sprintf("%v", err))
	}

	token, err := s.TokenService.Encode(req)
	if err != nil {
		return err
	}
	res.User = req
	res.Token = &proto.Token{Token: token}

	return nil
}

func (s *Service) Authenticate(ctx context.Context, req *proto.User, res *proto.AuthResponse) error {
	user, err := s.UserRepository.GetByEmail(req.Email)
	if err != nil {
		return err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return errors.Unauthorized(s.Name, "Invalid password.")
	}

	token, err := s.TokenService.Encode(req)
	if err != nil {
		return errors.InternalServerError(s.Name, fmt.Sprintf("%v", err))
	}
	res.Token = &proto.Token{Token: token}

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
	fmt.Println(ctx)
	meta, ok := metadata.FromContext(ctx)
	if !ok {
		return errors.New(s.Name, "Could not parse request headers (context metadata).", 500)
	}

	token := meta["Token"]

	claims, err := s.TokenService.Decode(token)
	if err != nil {
		return errors.Unauthorized(s.Name, "Invalid token")
	}
	res.User = claims.User

	return nil
}
