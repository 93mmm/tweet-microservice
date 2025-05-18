package mongo

import (
	"context"
	"errors"

	"github.com/93mmm/tweet-microservice/internal/storage/models"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func (s *mongoStorage) WriteNewTweet(ctx context.Context, tweet *models.TweetDocument) error {
	_, err := s.tweetsCollection().InsertOne(ctx, tweet)
	if err != nil {
		return err
	}

	return nil
}

func (s *mongoStorage) UpdateTweet(ctx context.Context, id int64, tweet *models.UpdateTweetDocument) (*models.TweetDocument, error) {
	result := s.tweetsCollection().FindOneAndUpdate(
		ctx,
		bson.M{"_id": id},
		bson.M{"$set": tweet},
		options.FindOneAndUpdate().SetReturnDocument(options.After),
	)

	model := &models.TweetDocument{}
	if err := result.Decode(model); err != nil {
		return nil, err
	}

	return model, nil
}

// TODO: implement
func (s *mongoStorage) DeleteTweet(ctx context.Context, id int64) error {
	return errors.New("unimplemented")
}

func (s *mongoStorage) GetTweet(ctx context.Context, id int64) (*models.TweetDocument, error) {
	response := s.tweetsCollection().FindOne(
		ctx,
		bson.M{"_id": id},
	)

	model := &models.TweetDocument{}
	err := response.Decode(model)
	if err != nil {
		return nil, err
	}

	return model, nil
}

// TODO: implement
func (s *mongoStorage) GetTweets(ctx context.Context, tweet *models.TweetDocument) error {
	return errors.New("unimplemented")
}
