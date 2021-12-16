package jwt

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// New generates new JWT service necessery for auth middleware
func New(secret, algo string, accessTokenDuration, refreshTokenDuration int) *Service {
	signingMethod := jwt.GetSigningMethod(algo)
	if signingMethod == nil {
		panic("invalid jwt signing method")
	}
	return &Service{
		key:                  []byte(secret),
		algo:                 signingMethod,
		accessTokenDuration:  time.Duration(accessTokenDuration) * time.Minute,
		refreshTokenDuration: time.Duration(refreshTokenDuration) * time.Minute,
	}
}

// Service provides a Json-Web-Token authentication implementation
type Service struct {
	// Secret key used for signing.
	key []byte

	// JWT signing algorithm
	algo jwt.SigningMethod

	// Duration represents token expiration time
	accessTokenDuration  time.Duration
	refreshTokenDuration time.Duration
}

// GenerateToken generates new JWT token and populates it with user data
func (s Service) GenerateAccessToken(ctx context.Context, userID int, email, profileImageURL string) (string, error) {
	return jwt.NewWithClaims(s.algo, jwt.MapClaims{
		"exp":               time.Now().Add(s.accessTokenDuration).Unix(),
		"user_id":           userID,
		"email":             email,
		"profile_image_url": profileImageURL,
	}).SignedString(s.key)
}
