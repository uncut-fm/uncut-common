package gin

import (
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
)

func SetResponseError(ctx *gin.Context, statusCode int, err error) {
	ctx.JSON(statusCode, gin.H{"message": err.Error()})
	ctx.Abort()
	if hub := sentrygin.GetHubFromContext(ctx); hub != nil {
		hub.CaptureException(err)
	}

	return
}
