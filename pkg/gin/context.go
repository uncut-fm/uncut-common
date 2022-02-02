package gin

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/uncut-fm/uncut-common/model"
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

func (c ContextService) SetUserToGinContext(ctx *gin.Context, user *model.User) {
	ctx.Set("email", user.Email)
	ctx.Set("id", user.UserId)
	ctx.Set("name", user.Name)
	ctx.Set("profile_image_url", user.ProfileImageUrl)
	ctx.Set("has_admin_panel_access", user.HasAdminPanelAccess)
	ctx.Set("wallet_address", user.WalletAddress)
}

func (c ContextService) GetUserFromContext(ctx context.Context) (*model.User, error) {
	ginContext, err := c.getGinContextFromContext(ctx)
	if err != nil {
		return nil, err
	}

	return c.GetUserFromGinContext(ginContext)
}

func (c ContextService) GetUserFromGinContext(ginContext *gin.Context) (*model.User, error) {
	if err := c.mustBeAuthenticatedUserGin(ginContext); err != nil {
		return nil, err
	}

	return &model.User{
		ID:                  ginContext.GetInt("id"),
		Name:                ginContext.GetString("name"),
		Email:               ginContext.GetString("email"),
		ProfileImageUrl:     ginContext.GetString("profile_image_url"),
		HasAdminPanelAccess: ginContext.GetBool("has_admin_panel_access"),
		WalletAddress:       ginContext.GetString("wallet_address"),
	}, nil
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
