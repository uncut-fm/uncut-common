package gin

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
)

type ContextKey = string

const (
	ginContextKey ContextKey = "GIN_CONTEXT_KEY"
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
