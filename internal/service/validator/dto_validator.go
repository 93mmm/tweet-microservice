package validator

import (
	"errors"
	"github.com/93mmm/tweet-microservice/internal/transport/dto"
)

var (
	InvalidAuthorID = errors.New("Invalid AuthorID")
	InvalidContent = errors.New("Invalid Content")
)

func ValidateCreateRequest(t *dto.TweetCreateRequest) error {
	if t.AuthorID == "" { // == 0
		return InvalidAuthorID
	}
	if len(t.Content) == 0 || len(t.Content) >= 280 {
		return InvalidContent
	}
	return nil
}
