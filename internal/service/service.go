package service

import (
	"github.com/93mmm/tweet-microservice/internal/mapper"
	"github.com/93mmm/tweet-microservice/internal/service/domain"
	"github.com/93mmm/tweet-microservice/internal/storage/models"
	"github.com/93mmm/tweet-microservice/internal/storage/mongo"
	"time"

	"github.com/bwmarrin/snowflake"
)

type TweetService interface {
	CreateTweet(tweet *domain.Tweet) (*domain.Tweet, error)

	GetTweetByID(id int64) (*domain.Tweet, error)
	// maybe tweet filter
	// GetTweets(tweet *domain.Tweet) ([]*domain.Tweet, error)

	UpdateTweet(id int64, content string) (*domain.Tweet, error)
	DeleteTweet(id int64) error
}

type tweetService struct {
	repo   mongo.Storage
	sfnode *snowflake.Node
}

func NewTweetService(repo mongo.Storage) (TweetService, error) {
	sfnode, err := snowflake.NewNode(1)
	if err != nil {
		return nil, err
	}

	// TODO: snowflake generated id == int64, rewrite all fields to int:)
	return &tweetService{
		repo:   repo,
		sfnode: sfnode,
	}, nil
}

func (s *tweetService) CreateTweet(tweet *domain.Tweet) (*domain.Tweet, error) {
	tweet.ID = int64(s.sfnode.Generate())
	tweet.CreatedAt = time.Now()

	err := s.repo.WriteNewTweet(mapper.TweetToMongoModel(tweet))
	if err != nil {
		return nil, err
	}

	return tweet, nil
}

func (s *tweetService) GetTweetByID(id int64) (*domain.Tweet, error) {
	tweet, err := s.repo.GetTweet(id)

	if err != nil {
		return nil, err
	}
	return mapper.MongoModelToTweet(tweet), nil
}

func (s *tweetService) GetTweets(tweet *domain.Tweet) ([]*domain.Tweet, error) {
	return make([]*domain.Tweet, 0), nil
}

func (s *tweetService) UpdateTweet(id int64, content string) (*domain.Tweet, error) {
	tweet, err := s.repo.UpdateTweet(
		id,
		&models.UpdateTweetMongo{
			Content:   content,
			IsEdited:  true,
			UpdatedAt: time.Now(),
		})
	if err != nil {
		return nil, err
	}

	return mapper.MongoModelToTweet(tweet), nil
}

func (s *tweetService) DeleteTweet(id int64) error {
	err := s.repo.DeleteTweet(id)

	if err != nil {
		return err
	}
	return nil
}
