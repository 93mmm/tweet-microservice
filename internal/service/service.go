package service

import (
	"context"
	"time"

	"github.com/93mmm/tweet-microservice/internal/mapper"
	"github.com/93mmm/tweet-microservice/internal/service/domain"
	"github.com/93mmm/tweet-microservice/internal/storage/models"
	"github.com/93mmm/tweet-microservice/internal/storage/mongo"

	"github.com/bwmarrin/snowflake"
)

type TweetService interface {
	CreateTweet(tweet *domain.Tweet) (*domain.Tweet, error)

	GetTweetByID(id string) (*domain.Tweet, error)
	// maybe tweet filter
	// GetTweets(tweet *domain.Tweet) ([]*domain.Tweet, error)

	UpdateTweet(id string, content string) (*domain.Tweet, error)
	DeleteTweet(id string) error
}

type tweetService struct {
	repo        mongo.Storage
	idGenerator *snowflake.Node
}

func NewTweetService(repo mongo.Storage) (TweetService, error) {
	sfnode, err := snowflake.NewNode(1)
	if err != nil {
		return nil, err
	}

	return &tweetService{
		repo:        repo,
		idGenerator: sfnode,
	}, nil
}

func (s *tweetService) CreateTweet(tweet *domain.Tweet) (*domain.Tweet, error) {
	tweet.ID = s.idGenerator.Generate().String()
	tweet.CreatedAt = time.Now()

	model := mapper.TweetToMongoModel(tweet)
	err := s.repo.WriteNewTweet(context.TODO(), model)
	if err != nil {
		return nil, err
	}

	return tweet, nil
}

func (s *tweetService) GetTweetByID(id string) (*domain.Tweet, error) {
	model, err := s.repo.GetTweet(context.TODO(), id)
	if err != nil {
		return nil, err
	}

	tweet := mapper.MongoModelToTweet(model)
	return tweet, nil
}

func (s *tweetService) GetTweets(tweet *domain.Tweet) ([]*domain.Tweet, error) {
	return make([]*domain.Tweet, 0), nil
}

func (s *tweetService) UpdateTweet(id string, content string) (*domain.Tweet, error) {
	model, err := s.repo.UpdateTweet(
		context.TODO(),
		id,
		&models.UpdateTweetDocument{
			Content:   content,
			IsEdited:  true,
			UpdatedAt: time.Now(),
		})
	if err != nil {
		return nil, err
	}

	tweet := mapper.MongoModelToTweet(model)
	return tweet, nil
}

func (s *tweetService) DeleteTweet(id string) error {
	err := s.repo.DeleteTweet(context.TODO(), id)
	if err != nil {
		return err
	}

	return nil
}
