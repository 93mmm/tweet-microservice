package mapper

import (
	"github.com/93mmm/tweet-microservice/internal/service/domain"
	"github.com/93mmm/tweet-microservice/internal/storage/models"
	"github.com/93mmm/tweet-microservice/internal/transport/dto"
)

func CreateDTOToTweet(t *dto.TweetCreateRequest) *domain.Tweet {
	return &domain.Tweet{
		AuthorID: t.AuthorID,
		Content:  t.Content,
	}
}

func MongoModelToTweet(t *models.TweetMongo) *domain.Tweet {
	return &domain.Tweet{
		ID:        t.ID,
		AuthorID:  t.AuthorID,
		Content:   t.Content,
		CreatedAt: t.CreatedAt,
		IsEdited:  t.IsEdited,
		UpdatedAt: t.UpdatedAt,
	}
}

func TweetToMongoModel(t *domain.Tweet) *models.TweetMongo {
	return &models.TweetMongo{
		ID:        t.ID,
		AuthorID:  t.AuthorID,
		Content:   t.Content,
		CreatedAt: t.CreatedAt,
		IsEdited:  t.IsEdited,
		UpdatedAt: t.UpdatedAt,
	}
}

func TweetToTweetDTO(t *domain.Tweet) *dto.TweetResponse {
	return &dto.TweetResponse{
		ID:        t.ID,
		AuthorID:  t.AuthorID,
		Content:   t.Content,
		CreatedAt: t.CreatedAt,
		IsEdited:  t.IsEdited,
		UpdatedAt: t.UpdatedAt,
	}
}
