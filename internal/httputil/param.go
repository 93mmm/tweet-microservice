package httputil

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

var (
	InvalidId = errors.New("invalid id passed")
)

func GetIDParam(ctx *gin.Context) (int64, error) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		return 0, InvalidId
	}
	return id, nil
}
