package httputil

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

var (
	InvalidId = errors.New("invalid id passed")
)

func GetIDParam(ctx *gin.Context) (string, error) {
	id := ctx.Param("id")
	
	if _, err := strconv.ParseInt(id, 10, 64); err != nil {
		return "", fmt.Errorf("invalid id passed: %v", id)
	}
	return id, nil
}
