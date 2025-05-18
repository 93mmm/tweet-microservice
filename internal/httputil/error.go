package httputil

import (
	"github.com/gin-gonic/gin"
)

func Error(ctx *gin.Context, err error, status int) {
	ctx.AbortWithStatusJSON(
		status,
		gin.H{"error": err},
	)
}
