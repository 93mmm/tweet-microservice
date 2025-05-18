package mongo

import (
	"context"
	"github.com/93mmm/tweet-microservice/internal/storage/models"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func (s *mongoStorage) WriteNewTweet(tweet *models.TweetMongo) error {
	_, err := s.tweetsCollection().InsertOne(context.TODO(), tweet)
	if err != nil {
		return err
	}

	return nil
}

func (s *mongoStorage) UpdateTweet(id int64, tweet *models.UpdateTweetMongo) (*models.TweetMongo, error) {
	objId := bsonId{ID: id}

	result := s.tweetsCollection().FindOneAndUpdate(context.TODO(),
		objId,
		bson.M{"$set": tweet},
		options.FindOneAndUpdate().SetReturnDocument(options.After),
	)

	editedTweet := &models.TweetMongo{}
	if err := result.Decode(editedTweet); err != nil {
		return nil, err
	}

	return editedTweet, nil
}

func (s *mongoStorage) DeleteTweet(id int64) error {
	return nil
}

func (s *mongoStorage) GetTweet(id int64) (*models.TweetMongo, error) {
	objId := &bsonId{ID: id}
	response := s.tweetsCollection().FindOne(context.TODO(), objId)

	tweet := &models.TweetMongo{}
	err := response.Decode(tweet)
	if err != nil {
		return nil, err
	}

	return tweet, nil
}

func (s *mongoStorage) GetTweets(tweet *models.TweetMongo) error {
	return nil
}
