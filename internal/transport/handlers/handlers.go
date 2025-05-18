package handlers

import (
	"net/http"

	"github.com/93mmm/tweet-microservice/internal/httputil"
	"github.com/93mmm/tweet-microservice/internal/mapper"
	"github.com/93mmm/tweet-microservice/internal/service"
	"github.com/93mmm/tweet-microservice/internal/service/validator"
	"github.com/93mmm/tweet-microservice/internal/transport/dto"

	"github.com/gin-gonic/gin"
)

type TweetHandler interface {
	CreateTweet(ctx *gin.Context)
	UpdateTweet(ctx *gin.Context)
	DeleteTweet(ctx *gin.Context)

	GetTweetByID(ctx *gin.Context)
	GetTweets(ctx *gin.Context)
}

/*
::: HAHA SWAGGER LOLOLOL :::
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
	req := &dto.CreateTweetRequest{}
	if err := ctx.ShouldBind(req); err != nil {
		httputil.Error(ctx, err, http.StatusBadRequest)
		return
	}
	if err := validator.ValidateRequest(req); err != nil {
		httputil.Error(ctx, err, http.StatusBadRequest)
		return
	}

	tweet := mapper.CreateRequestToTweet(req)

	createdTweet, err := h.svc.CreateTweet(tweet)
	if err != nil {
		httputil.Error(ctx, err, http.StatusBadRequest)
		return
	}
	resp := mapper.TweetToTweetResponse(createdTweet)

	ctx.JSON(http.StatusCreated, resp)
}

func (h *tweetHandler) GetTweetByID(ctx *gin.Context) {
	id, err := httputil.GetIDParam(ctx)
	if err != nil {
		httputil.Error(ctx, err, http.StatusBadRequest)
		return
	}

	tweet, err := h.svc.GetTweetByID(id)
	if err != nil {
		httputil.Error(ctx, err, http.StatusBadRequest)
		return
	}

	resp := mapper.TweetToTweetResponse(tweet)
	ctx.JSON(http.StatusOK, resp)
}

// TODO: implement later
func (h *tweetHandler) GetTweets(ctx *gin.Context) {
	req := &dto.ListTweetsRequest{}
	if err := ctx.ShouldBind(req); err != nil {
		httputil.Error(ctx, err, http.StatusBadRequest)
		return
	}
	if err := validator.ValidateRequest(req); err != nil {
		httputil.Error(ctx, err, http.StatusBadRequest)
		return
	}

	ctx.JSON(http.StatusNotImplemented, req)
}

func (h *tweetHandler) UpdateTweet(ctx *gin.Context) {
	req := &dto.UpdateTweetRequest{}
	if err := ctx.ShouldBind(req); err != nil {
		httputil.Error(ctx, err, http.StatusBadRequest)
		return
	}
	if err := validator.ValidateRequest(req); err != nil {
		httputil.Error(ctx, err, http.StatusBadRequest)
		return
	}

	id, err := httputil.GetIDParam(ctx)
	if err != nil {
		httputil.Error(ctx, err, http.StatusBadRequest)
		return
	}

	tweet, err := h.svc.UpdateTweet(id, req.Content)
	if err != nil {
		httputil.Error(ctx, err, http.StatusBadRequest)
		return
	}

	resp := mapper.TweetToTweetResponse(tweet)

	ctx.JSON(http.StatusOK, resp)
}

func (h *tweetHandler) DeleteTweet(ctx *gin.Context) {
	id, err := httputil.GetIDParam(ctx)
	if err != nil {
		httputil.Error(ctx, err, http.StatusBadRequest)
		return
	}

	err = h.svc.DeleteTweet(id)
	if err != nil {
		httputil.Error(ctx, err, http.StatusNotFound)
		return
	}

	ctx.Status(http.StatusNoContent)
}
