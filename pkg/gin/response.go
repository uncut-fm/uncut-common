package gin

import "github.com/gin-gonic/gin"

func SetResponseError(ctx *gin.Context, statusCode int, err error) {
	ctx.JSON(statusCode, gin.H{"message": err.Error()})
	ctx.Abort()
	return
}
