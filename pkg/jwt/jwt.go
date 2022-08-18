package jwt

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/uncut-fm/uncut-common/model"
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
	SetUserToGinContext(ctx *gin.Context, user *model.User)
	SetAuthenticatedUserKey(ctx *gin.Context, authenticated bool)
}

// GenerateAccessToken generates new JWT token and populates it with user data
func (s Service) GenerateAccessToken(ctx context.Context, user model.User) (string, error) {
	return jwt.NewWithClaims(s.algo, jwt.MapClaims{
		"exp":               time.Now().Add(s.accessTokenDuration).Unix(),
		"user_id":           user.UserId,
		"name":              user.Name,
		"email":             user.Email,
		"profile_image_url": user.ProfileImageUrl,
		"wallet_addresses":  user.WalletAddresses,
		"twitter_handle":    user.TwitterHandle,
		"is_nft_creator":    user.IsNftCreator,
	}).SignedString(s.accessKey)
}

func (s Service) MWFunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := s.parseTokenFromHeader(c)
		if err != nil {
			s.context.SetAuthenticatedUserKey(c, false)
			c.Next()
			return
		}

		user, err := s.getUserFromToken(token)
		if err != nil {
			s.context.SetAuthenticatedUserKey(c, false)
			c.Next()
			return
		}

		s.context.SetAuthenticatedUserKey(c, true)

		s.context.SetUserToGinContext(c, user)

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

func (s Service) getUserFromToken(token string) (*model.User, error) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return s.accessKey, nil
	})

	if err != nil {
		return nil, err
	}

	claimsBytes, err := json.Marshal(claims)
	if err != nil {
		return nil, err
	}

	user := new(model.User)

	err = json.Unmarshal(claimsBytes, user)
	if err != nil {
		return nil, err
	}

	return user, err
}
