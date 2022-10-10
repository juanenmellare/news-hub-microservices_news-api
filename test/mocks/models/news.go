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
	uuidMock, _ := uuid.FromString("800d249f-a7f7-4129-a8a6-14d0cf9667e5")
	publishedAt := time.Date(2022, 10, 2, 1, 1, 1, 1, time.Local)
	return &NewsBuilder{
		user: models.News{
			ID:          uuidMock,
			ImageUrl:    "foo-image-url",
			Channel:     "foo-channel",
			Title:       "foo-title",
			Url:         "foo-url",
			PublishedAt: publishedAt,
		},
	}
}
