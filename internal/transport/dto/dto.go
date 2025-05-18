package dto

import (
	"time"
)

type TweetCreateRequest struct {
	AuthorID string `json:"author_id"`
	Content  string `json:"content"`
}

type TweetUpdateRequest struct {
	Content string `json:"content"`
}

type TweetListRequest struct {
	AuthorID string    `json:"author_id,omitempty"`
	Limit    int       `json:"limit,omitempty"`
	Offset   int       `json:"offset,omitempty"`
	Since    time.Time `json:"since,omitempty"`
}

type TweetResponse struct {
	ID        string    `json:"id"`
	AuthorID  string    `json:"author_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	IsEdited  bool      `json:"is_edited"`
	UpdatedAt time.Time `json:"updated_at"`
}

type TweetListResponse struct {
	Tweets []TweetResponse `json:"tweets"`
	Count  int             `json:"count"`
}
