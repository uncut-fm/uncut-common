package gin

import (
	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
)

func SetResponseError(ctx *gin.Context, statusCode int, err error) {
	ctx.JSON(statusCode, gin.H{"message": err.Error()})
	ctx.Abort()
	if hub := sentry.GetHubFromContext(ctx); hub != nil {
		hub.CaptureException(err)
	}

	return
}
