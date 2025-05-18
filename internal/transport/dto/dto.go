package dto

import (
	"time"
)

type CreateTweetRequest struct {
	AuthorID string `json:"author_id"`
	Content  string `json:"content"`
}

type UpdateTweetRequest struct {
	Content string `json:"content"`
}

type ListTweetsRequest struct {
	AuthorID string    `json:"author_id,omitempty"`
	Limit    int       `json:"limit,omitempty"`
	Offset   int       `json:"offset,omitempty"`
	Since    time.Time `json:"since,omitempty"`
}

type TweetResponse struct {
	ID        int64     `json:"id"`
	AuthorID  string    `json:"author_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	IsEdited  bool      `json:"is_edited"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ListTweetsResponse struct {
	Tweets []TweetResponse `json:"tweets"`
	Count  int             `json:"count"`
}
