package auth

import (
	"time"

	proto "github.com/Condition17/fleet-services/user-service/proto/user-service"
	"github.com/dgrijalva/jwt-go"
)

var key = []byte("mySuperSecretKey")

// custom metadata which will be hashed
// and sent as the second segment in our JWT
type CustomClaims struct {
	User *proto.User
	jwt.StandardClaims
}

type AuthChecker interface {
	Decode(token string) (*CustomClaims, error)
	Encode(user *proto.User) (string, error)
	ValidateToken(token string) error
}

type TokenService struct {
	Issuer string
}

// Decode a token string into a token object
func (s *TokenService) Decode(tokenStr string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		// TODO: THIS IS NOT CORRECT!! FIX THIS;
		return nil, err
	}
}

// Encode a claim into a JWT
func (s *TokenService) Encode(user *proto.User) (string, error) {
	tokenExpireTime := time.Now().Add(time.Hour * 24 * 30 * 12).Unix()

	// Create the Claimsxww
	claims := CustomClaims{
		user,
		jwt.StandardClaims{
			ExpiresAt: tokenExpireTime,
			Issuer:    s.Issuer,
		},
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token and return
	return token.SignedString(key)
}

// Validate token
// Returns nil if the passed token is valid
// Error otherwise
func (s *TokenService) ValidateToken(tokenStr string) error {
	if _, err := s.Decode(tokenStr); err != nil {
		return err
	}
	return nil
}
