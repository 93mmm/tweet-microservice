package routers

import (
	"github.com/93mmm/tweet-microservice/internal/transport/handlers"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine, svc handlers.TweetHandler) {
	{
		router := r.Group("/api/v1/tweets")

		router.POST("", svc.CreateTweet)

		router.GET("/:id", svc.GetTweetByID)
		router.GET("", svc.GetTweets)

		router.PUT("/:id", svc.UpdateTweet)

		router.DELETE("/:id", svc.DeleteTweet)
	}
}
