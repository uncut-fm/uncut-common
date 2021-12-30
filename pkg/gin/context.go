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

func CreateWithGinContext(ctx context.Context, gin *gin.Context) context.Context {
	return context.WithValue(ctx, ginContextKey, gin)
}

func GetGinContextFromContext(ctx context.Context) (*gin.Context, error) {
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

func SetUserToGinContext(ctx *gin.Context, user *model.User) {
	ctx.Set("email", user.Email)
	ctx.Set("id", user.UserId)
	ctx.Set("name", user.Name)
	ctx.Set("profile_image_url", user.ProfileImageUrl)
	ctx.Set("has_admin_panel_access", user.HasAdminPanelAccess)
}

func GetUserFromContext(ctx context.Context) (*model.User, error) {
	ginContext, err := GetGinContextFromContext(ctx)
	if err != nil {
		return nil, err
	}

	if !ginContext.GetBool(AuthenticatedContextKey) {
		return nil, errors.New("user is not authenticated")
	}

	return &model.User{
		ID:                  ginContext.GetInt("id"),
		Name:                ginContext.GetString("name"),
		Email:               ginContext.GetString("email"),
		ProfileImageUrl:     ginContext.GetString("profile_image_url"),
		HasAdminPanelAccess: ginContext.GetBool("has_admin_panel_access"),
	}, nil
}
