package gin

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

type ContextKey = string

const (
	ginContextKey                ContextKey = "GIN_CONTEXT_KEY"
	AuthenticatedAdminContextKey            = "authenticated-token"
	AuthenticatedUserContextKey             = "authenticated-user"
)

type ContextService struct{}

func NewContextService() *ContextService {
	return &ContextService{}
}

func CreateWithGinContext(ctx context.Context, gin *gin.Context) context.Context {
	return context.WithValue(ctx, ginContextKey, gin)
}

func (c ContextService) getGinContextFromContext(ctx context.Context) (*gin.Context, error) {
	ginContext := ctx.Value(ginContextKey)

	if ginContext == nil {
		err := fmt.Errorf("cannot retrieve gin.Context")
		return nil, err
	}

	gc, ok := ginContext.(*gin.Context)
	if !ok {
		err := fmt.Errorf("gin.Context has wrong type")
		return nil, err
	}

	return gc, nil
}

func (c ContextService) SetUserIDToGinContext(ctx *gin.Context, userID int) {
	ctx.Set("id", userID)
}

func (c ContextService) GetUserIDFromContext(ctx context.Context) (int, error) {
	ginContext, err := c.getGinContextFromContext(ctx)
	if err != nil {
		return 0, err
	}

	userID, err := c.GetUserIDFromGinContext(ginContext)

	return userID, err
}

func (c ContextService) GetUserIDFromGinContext(ginContext *gin.Context) (int, error) {
	if err := c.mustBeAuthenticatedUserGin(ginContext); err != nil {
		return 0, err
	}

	return ginContext.GetInt("id"), nil
}

func (c ContextService) SetAuthenticatedUserKey(ctx *gin.Context, authenticated bool) {
	ctx.Set(AuthenticatedUserContextKey, authenticated)
	return
}

func (c ContextService) SetAuthenticatedAdminKey(ctx *gin.Context, authenticated bool) {
	ctx.Set(AuthenticatedAdminContextKey, authenticated)
	return
}

func (c ContextService) IsAuthenticatedAdmin(ctx context.Context) (bool, error) {
	ginContext, err := c.getGinContextFromContext(ctx)
	if err != nil {
		return false, err
	}

	return c.isAuthenticatedAdminGin(ginContext), nil
}

func (c ContextService) MustBeAuthenticatedAdmin(ctx context.Context) error {
	ginContext, err := c.getGinContextFromContext(ctx)
	if err != nil {
		return err
	}

	return c.mustBeAuthenticatedAdminGin(ginContext)
}

func (c ContextService) mustBeAuthenticatedAdminGin(ctx *gin.Context) error {
	authenticated := c.isAuthenticatedAdminGin(ctx)
	if authenticated {
		return nil
	}

	return errors.New("not authenticated")
}

func (c ContextService) MustBeAuthenticatedUser(ctx context.Context) error {
	ginContext, err := c.getGinContextFromContext(ctx)
	if err != nil {
		return err
	}

	return c.mustBeAuthenticatedUserGin(ginContext)
}

func (c ContextService) mustBeAuthenticatedUserGin(ctx *gin.Context) error {
	authenticated := c.isAuthenticatedUserGin(ctx)
	if authenticated {
		return nil
	}

	return errors.New("not authenticated")
}

func (c ContextService) isAuthenticatedAdminGin(ctx *gin.Context) bool {
	return ctx.GetBool(AuthenticatedAdminContextKey)
}

func (c ContextService) isAuthenticatedUserGin(ctx *gin.Context) bool {
	return ctx.GetBool(AuthenticatedUserContextKey)
}

func (c ContextService) GetIpFromContext(ctx context.Context) (string, error) {
	ginContext, err := c.getGinContextFromContext(ctx)
	if err != nil {
		log.Println("failed retrieving gin context, err: ", err.Error())
		return "", err
	}

	if ginContext.Request == nil {
		return "", nil
	}

	return ginContext.ClientIP(), nil
}
