package models

import "time"

type TweetMongo struct {
	ID        int64       `bson:"_id"`
	AuthorID  string    `bson:"author_id"`
	Content   string    `bson:"content"`
	CreatedAt time.Time `bson:"created_at"`
	IsEdited  bool      `bson:"is_edited"`
	UpdatedAt time.Time `bson:"updated_at"`
}

type UpdateTweetMongo struct {
	Content   string    `bson:"content"`
	IsEdited  bool      `bson:"is_edited"`
	UpdatedAt time.Time `bson:"updated_at"`
}
