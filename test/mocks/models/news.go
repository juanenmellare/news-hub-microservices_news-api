package models

import (
	"github.com/gofrs/uuid"
	"news-hub-microservices_news-api/internal/models"
	"time"
)

type NewsBuilder struct {
	news models.News
}

func (n NewsBuilder) AddReader(userId uuid.UUID) *NewsBuilder {
	n.news.NewsReaders = append(n.news.NewsReaders, models.NewsReader{UserId: userId, NewsId: n.news.ID})
	return &n
}

func (n NewsBuilder) SetId(id uuid.UUID) *NewsBuilder {
	n.news.ID = id
	return &n
}

func (n NewsBuilder) Build() models.News {
	return n.news
}

func NewNewsBuilder() *NewsBuilder {
	uuidMock, _ := uuid.FromString("800d249f-a7f7-4129-a8a6-14d0cf9667e5")
	publishedAt := time.Date(2022, 10, 2, 1, 1, 1, 1, time.Local)

	return &NewsBuilder{
		news: models.News{
			ID:          uuidMock,
			ImageUrl:    "foo-image-url",
			Channel:     "foo-channel",
			Title:       "foo-title",
			Url:         "foo-url",
			PublishedAt: publishedAt,
		},
	}
}
