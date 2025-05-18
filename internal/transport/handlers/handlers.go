package handlers

import (
	"github.com/93mmm/tweet-microservice/internal/mapper"
	"github.com/93mmm/tweet-microservice/internal/service"
	"github.com/93mmm/tweet-microservice/internal/service/validator"
	"github.com/93mmm/tweet-microservice/internal/transport/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TweetHandler interface {
	CreateTweet(ctx *gin.Context)
	UpdateTweet(ctx *gin.Context)
	DeleteTweet(ctx *gin.Context)

	GetTweetById(ctx *gin.Context)
	GetTweets(ctx *gin.Context)
}

/*
									:::СВАГГЕР ДОКУМЕНТАЦИЯ ЧИСТА БЛЯ:::
POST /api/v1/tweets
TweetCreateRequest
TweetCreateResponce

GET /api/v1/tweets
null
TweetListResponse

PUT /api/v1/tweets/:id
TweetUpdateRequest
TweetResponse

DELETE /api/v1/tweets/:id
null
null
*/

type tweetHandler struct {
	svc service.TweetService
}

func NewTweetHandler(svc service.TweetService) *tweetHandler {
	return &tweetHandler{
		svc: svc,
	}
}

func (h *tweetHandler) CreateTweet(ctx *gin.Context) {
	dto := &dto.TweetCreateRequest{}
	if err := ctx.ShouldBind(dto); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err) // TODO: wrap to { "error": {error} }
		return
	}
	if err := validator.ValidateCreateRequest(dto); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	tweet := mapper.CreateDTOToTweet(dto)

	response, err := h.svc.CreateTweet(tweet)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	responseDTO := mapper.TweetToTweetDTO(response)

	ctx.JSON(http.StatusCreated, responseDTO)
}

func (h *tweetHandler) GetTweetById(ctx *gin.Context) {
	response, err := h.svc.GetTweetById(ctx.Param("id")) // TODO: param must be int64
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

// TODO: implement later
func (h *tweetHandler) GetTweets(ctx *gin.Context) {
	var dto dto.TweetListRequest

	if err := ctx.ShouldBind(&dto); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, dto)
}

func (h *tweetHandler) UpdateTweet(ctx *gin.Context) {
	var dto dto.TweetUpdateRequest

	if err := ctx.ShouldBind(&dto); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response, err := h.svc.UpdateTweet(ctx.Param("id"), dto.Content)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (h *tweetHandler) DeleteTweet(ctx *gin.Context) {
	err := h.svc.DeleteTweet(ctx.Param("id"))

	if err != nil {
		ctx.AbortWithError(http.StatusNotFound, err)
		return
	}

	ctx.Status(http.StatusOK)
}
