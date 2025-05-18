package middleware

import "github.com/gin-gonic/gin"

func SetupMiddleware(router *gin.Engine) {
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
}
