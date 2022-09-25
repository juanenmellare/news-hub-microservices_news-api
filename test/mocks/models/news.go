package models

import (
	"github.com/gofrs/uuid"
	"news-hub-microservices_news-api/internal/models"
	"time"
)

type NewsBuilder struct {
	user models.News
}

func (u NewsBuilder) Build() models.News {
	return u.user
}

func NewNewsBuilder() *NewsBuilder {
	uuidMock, _ := uuid.NewV4()

	return &NewsBuilder{
		user: models.News{
			ID:          uuidMock,
			ImageUrl:    "foo-image-url",
			Channel:     "foo-channel",
			Title:       "foo-title",
			Url:         "foo-url",
			PublishedAt: time.Now(),
		},
	}
}
