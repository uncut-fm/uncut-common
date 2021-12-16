package jwt

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/uncut-fm/uncut-common/pkg/config"
	"time"
)

// New generates new JWT service necessery for auth middleware
func New(jwtConfigs config.JWTConfigs) *Service {
	signingMethod := jwt.GetSigningMethod(jwtConfigs.SigningMethod)
	if signingMethod == nil {
		panic("invalid jwt signing method")
	}
	return &Service{
		accessKey:            []byte(jwtConfigs.AccessSecret),
		refreshKey:           []byte(jwtConfigs.RefreshSecret),
		algo:                 signingMethod,
		accessTokenDuration:  time.Duration(jwtConfigs.AccessDuration) * time.Minute,
		refreshTokenDuration: time.Duration(jwtConfigs.RefreshDuration) * time.Minute,
	}
}

// Service provides a Json-Web-Token authentication implementation
type Service struct {
	// Secret accessKey used for signing.
	accessKey  []byte
	refreshKey []byte

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
	}).SignedString(s.accessKey)
}
