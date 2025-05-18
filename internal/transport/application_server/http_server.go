package application_server

import (
	"github.com/93mmm/tweet-microservice/internal/transport/handlers"
	"github.com/93mmm/tweet-microservice/internal/transport/routers"

	"github.com/gin-gonic/gin"
)

type apiServer struct {
	root *gin.Engine
	svc  handlers.TweetHandler
}

func NewServer(root *gin.Engine, svc handlers.TweetHandler) *apiServer {
	return &apiServer{
		root: root,
		svc:  svc,
	}
}

func (s *apiServer) StartApiServer(listenAddr string) error {
	routers.InitRoutes(s.root, s.svc)
	// TODO: raise handlers
	return s.root.Run(listenAddr)
}
