package domain

import (
	"time"
)

type Tweet struct {
	ID        int64
	AuthorID  string
	Content   string
	CreatedAt time.Time
	IsEdited  bool
	UpdatedAt time.Time
	IsDeleted bool
}
