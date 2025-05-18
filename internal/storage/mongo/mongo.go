package mongo

import (
	"context"
	"github.com/93mmm/tweet-microservice/internal/config"
	"github.com/93mmm/tweet-microservice/internal/storage/models"
	"log"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

type Storage interface {
	WriteNewTweet(tweet *models.TweetMongo) error

	UpdateTweet(id int64, tweet *models.UpdateTweetMongo) (*models.TweetMongo, error)

	DeleteTweet(id int64) error

	GetTweet(id int64) (*models.TweetMongo, error)
	Disconnect()
	// GetTweets(tweet *models.TweetMongo) error
}

type mongoStorage struct {
	db *mongo.Client
}

type bsonId struct {
	ID int64 `bson:"_id"`
}

func NewMongoStorage() (Storage, error) {
	url := config.Mongo().ConnectionString()
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

func IdToBson(id int64) any {
	return &bsonId{ID: id}
}
