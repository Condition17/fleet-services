package handler

import (
	"context"
	"errors"
	"fmt"

	"gorm.io/gorm"

	"github.com/Condition17/fleet-services/lib/auth"
	"github.com/Condition17/fleet-services/user-service/model"
	proto "github.com/Condition17/fleet-services/user-service/proto/user-service"
	microErrors "github.com/micro/go-micro/v2/errors"
	"golang.org/x/crypto/bcrypt"
)

func (h *Handler) Create(ctx context.Context, req *proto.User, res *proto.EmptyResponse) error {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return microErrors.BadRequest(h.Service.Name(), "Unapropriate password format.")
	}

	// Check if user already exists
	if user, _ := h.UserRepository.GetByEmail(req.Email); user != nil {
		return microErrors.Conflict(h.Service.Name(), "An user with this email was already created. Please try another email.")
	}

	// Create user entry
	req.Password = string(hashedPass)
	if err := h.UserRepository.Create(model.MarshalUser(req)); err != nil {
		return microErrors.InternalServerError(h.Service.Name(), fmt.Sprintf("%v", err))
	}

	return nil
}

func (h *Handler) Authenticate(ctx context.Context, req *proto.User, res *proto.AuthResponse) error {
	user, err := h.GetUserByEmail(req.Email)
	if err != nil {
		return err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return microErrors.Unauthorized(h.Service.Name(), "Invalid password.")
	}

	if res.Token, err = h.generateToken(model.UnmarshalUser(user)); err != nil {
		return microErrors.InternalServerError(h.Service.Name(), fmt.Sprintf("%v", err))
	}
	res.User = model.UnmarshalUser(user)

	return nil
}

func (h *Handler) ValidateToken(ctx context.Context, req *proto.Token, res *proto.TokenValidationResponse) error {
	if err := h.TokenService.ValidateToken(req.Token); err != nil {
		res.Valid = false
		return err
	}
	res.Valid = true

	return nil
}

func (h *Handler) GetProfile(ctx context.Context, req *proto.EmptyRequest, res *proto.AuthResponse) error {
	// TODO: refactor here. Use logic existent in lib/auth
	var tokenBytes []byte = auth.GetTokenBytesFromContext(ctx)
	claims, err := h.TokenService.Decode(string(tokenBytes))

	if err != nil {
		return microErrors.Unauthorized(h.Service.Name(), "Invalid token")
	}
	// -- end refactor

	userEntry, err := h.GetUserByEmail(ctx, claims.User.Email)
	if err != nil {
		return err
	}

	res.User = model.UnmarshalUser(userEntry)

	return nil
}

func (h *Handler) generateToken(payload *proto.User) (*proto.Token, error) {
	token, err := h.TokenService.Encode(payload)
	if err != nil {
		return nil, err
	}

	return &proto.Token{Token: token}, nil
}

func (h *Handler) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	user, err := h.UserRepository.GetByEmail(email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, microErrors.Unauthorized(h.Service.Name(), "User with this email not found")
		}
		return nil, microErrors.InternalServerError(h.Service.Name(), fmt.Sprintf("%v", err))
	}

	return user, nil
}
