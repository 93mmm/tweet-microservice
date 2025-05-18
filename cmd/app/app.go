package main

import (
	"github.com/93mmm/tweet-microservice/internal/config"
	"github.com/93mmm/tweet-microservice/internal/service"
	"github.com/93mmm/tweet-microservice/internal/storage/mongo"
	"github.com/93mmm/tweet-microservice/internal/transport/application_server"
	"github.com/93mmm/tweet-microservice/internal/transport/handlers"
	"github.com/93mmm/tweet-microservice/internal/transport/middleware"

	"github.com/gin-gonic/gin"
)

func Run() error {
	config.Load()
	router := gin.New()
	middleware.SetupMiddleware(router)

	storage, err := mongo.NewMongoStorage()
	if err != nil {
		return err
	}
	defer storage.Disconnect()

	service, err := service.NewTweetService(storage)
	if err != nil {
		return err
	}

	handler := handlers.NewTweetHandler(service)

	server := application_server.NewServer(router, handler)

	return server.StartApiServer("app:8080") // TODO: import host and port from config
}
