package domain

import (
	"time"
)

type Tweet struct {
	ID        string
	AuthorID  string
	Content   string
	CreatedAt time.Time
	IsEdited  bool
	UpdatedAt time.Time
	IsDeleted bool
}
