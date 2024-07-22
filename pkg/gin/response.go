package gin

import (
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
)

func SetResponseError(ctx *gin.Context, statusCode int, err error, withSentryCapture bool) {
	ctx.JSON(statusCode, gin.H{"message": err.Error()})
	ctx.Abort()

	if withSentryCapture {
		if hub := sentrygin.GetHubFromContext(ctx); hub != nil {
			hub.CaptureException(err)
		}
	}

	return
}
