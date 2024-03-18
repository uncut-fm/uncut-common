package jwt

import (
	"context"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/uncut-fm/uncut-common/pkg/config"
	"strings"
	"time"
)

const (
	authorizationHeader = "Authorization"
)

// New generates new JWT service necessery for auth middleware
func New(jwtConfigs config.JWTConfigs, ctx Context) *Service {
	signingMethod := jwt.GetSigningMethod(jwtConfigs.SigningMethod)
	if signingMethod == nil {
		panic("invalid jwt signing method")
	}
	return &Service{
		accessKey:           []byte(jwtConfigs.AccessSecret),
		algo:                signingMethod,
		accessTokenDuration: time.Duration(jwtConfigs.AccessDuration) * time.Minute,
		context:             ctx,
	}
}

// Service provides a Json-Web-Token authentication implementation
type Service struct {
	// Secret accessKey used for signing.
	accessKey []byte

	// JWT signing algorithm
	algo jwt.SigningMethod

	// Duration represents token expiration time
	accessTokenDuration time.Duration

	context Context
}

type Context interface {
	SetUserIDToGinContext(ctx *gin.Context, userID int)
	SetAuthenticatedUserKey(ctx *gin.Context, authenticated bool)
}

// GenerateAccessToken generates new JWT token and populates it with user data
func (s Service) GenerateAccessToken(ctx context.Context, userID int, expirable bool) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
	}

	if expirable {
		claims["exp"] = time.Now().Add(s.accessTokenDuration).Unix()
	}

	return jwt.NewWithClaims(s.algo, claims).SignedString(s.accessKey)
}

func (s Service) MWFunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := s.parseTokenFromHeader(c)
		if err != nil {
			s.context.SetAuthenticatedUserKey(c, false)
			c.Next()
			return
		}

		userID, err := s.GetUserIDFromToken(token)
		if err != nil {
			s.context.SetAuthenticatedUserKey(c, false)
			c.Next()
			return
		}

		s.context.SetAuthenticatedUserKey(c, true)

		s.context.SetUserIDToGinContext(c, userID)

		c.Next()
	}
}

// parseToken parses token from Authorization header
func (s Service) parseTokenFromHeader(c *gin.Context) (string, error) {
	token := c.GetHeader(authorizationHeader)
	if token == "" {
		return "", errors.New("authorization header is missing")
	}

	parts := strings.SplitN(token, " ", 2)
	if len(parts) < 2 {
		return "", errors.New("token is missing")
	}

	return parts[1], nil
}

func (s Service) GetUserIDFromToken(token string) (int, error) {
	claims, err := s.getClaimsFromJwtToken(token)
	if err != nil {
		return 0, err
	}

	userIDFloat64, ok := claims["user_id"].(float64)
	if !ok {
		return 0, errors.New("invalid user_id")
	}

	return int(userIDFloat64), err
}

func (s Service) getClaimsFromJwtToken(token string) (jwt.MapClaims, error) {
	claims := jwt.MapClaims{}

	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		// Check the signing method
		if token.Method.Alg() != s.algo.Alg() {
			return nil, errors.New("invalid signing method")
		}

		return s.accessKey, nil
	})

	if err != nil {
		return nil, err
	}

	return claims, nil
}
