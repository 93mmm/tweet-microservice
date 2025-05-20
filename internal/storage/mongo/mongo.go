package mongo

import (
	"context"
	"github.com/93mmm/tweet-microservice/internal/storage/models"
	"log"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

type Storage interface {
	WriteNewTweet(ctx context.Context, tweet *models.TweetDocument) error

	GetTweet(ctx context.Context, id string) (*models.TweetDocument, error)
	GetTweets(ctx context.Context, tweet *models.TweetDocument) error

	UpdateTweet(ctx context.Context, id string, tweet *models.UpdateTweetDocument) (*models.TweetDocument, error)

	DeleteTweet(ctx context.Context, id string) error

	Disconnect()
}

type mongoStorage struct {
	db *mongo.Client
}

func NewMongoStorage(url string) (Storage, error) {
	client, err := mongo.Connect(options.Client().ApplyURI(url))
	if err != nil {
		return nil, err
	}
	if err = client.Ping(context.TODO(), readpref.Nearest()); err != nil {
		return nil, err
	}

	return &mongoStorage{
		db: client,
	}, nil
}

func (s *mongoStorage) Disconnect() {
	if err := s.db.Disconnect(context.Background()); err != nil {
		log.Fatal(err)
	}
}

func (s *mongoStorage) tweetsCollection() *mongo.Collection {
	return s.db.Database("tweet").Collection("tweets")
}
