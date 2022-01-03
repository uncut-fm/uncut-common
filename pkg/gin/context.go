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
	ginContextKey           ContextKey = "GIN_CONTEXT_KEY"
	AuthenticatedContextKey            = "authenticated"
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
		err := fmt.Errorf("could not retrieve gin.Context")
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
}

func (c ContextService) GetUserFromContext(ctx context.Context) (*model.User, error) {
	ginContext, err := c.getGinContextFromContext(ctx)
	if err != nil {
		return nil, err
	}

	if err := c.mustBeAuthenticatedGin(ginContext); err != nil {
		return nil, err
	}

	return &model.User{
		ID:                  ginContext.GetInt("id"),
		Name:                ginContext.GetString("name"),
		Email:               ginContext.GetString("email"),
		ProfileImageUrl:     ginContext.GetString("profile_image_url"),
		HasAdminPanelAccess: ginContext.GetBool("has_admin_panel_access"),
	}, nil
}

func (c ContextService) SetAuthenticatedFlag(ctx *gin.Context, authenticated bool) {
	ctx.Set(AuthenticatedContextKey, authenticated)
	return
}

func (c ContextService) IsAuthenticated(ctx context.Context) (bool, error) {
	ginContext, err := c.getGinContextFromContext(ctx)
	if err != nil {
		return false, err
	}

	return c.isAuthenticatedGin(ginContext), nil
}

func (c ContextService) MustBeAuthenticated(ctx context.Context) error {
	ginContext, err := c.getGinContextFromContext(ctx)
	if err != nil {
		return err
	}

	return c.mustBeAuthenticatedGin(ginContext)
}

func (c ContextService) mustBeAuthenticatedGin(ctx *gin.Context) error {
	authenticated := c.isAuthenticatedGin(ctx)
	if authenticated {
		return nil
	}

	return errors.New("not authenticated")
}

func (c ContextService) isAuthenticatedGin(ctx *gin.Context) bool {
	return ctx.GetBool(AuthenticatedContextKey)
}
