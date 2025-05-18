package validator

import (
	"errors"
	"github.com/93mmm/tweet-microservice/internal/transport/dto"
)

var (
	InvalidAuthorID   = errors.New("Invalid AuthorID")
	InvalidContent    = errors.New("Invalid Content")
	InvalidTypePassed = errors.New("Invalid Type Passed")
)

func ValidateRequest(t any) error {
	switch v := t.(type) {
	case *dto.CreateTweetRequest:
		if len(v.AuthorID) == 0 {
			return InvalidAuthorID
		}
		if len(v.Content) == 0 || len(v.Content) >= 280 {
			return InvalidContent
		}
		return nil
	case *dto.UpdateTweetRequest:
		if len(v.Content) == 0 || len(v.Content) >= 280 {
			return InvalidContent
		}
		return nil
	case *dto.ListTweetsRequest:
		return nil
	default:
		return InvalidTypePassed
	}
}
