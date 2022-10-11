package rest

import (
	"github.com/gofrs/uuid"
	"math"
	"news-hub-microservices_news-api/internal/models"
	"time"
)

type newsResponse struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	ImageUrl    string    `json:"imageUrl"`
	Channel     string    `json:"channel" `
	Url         string    `json:"url"`
	PublishedAt time.Time `json:"publishedAt"`
	HasBeenRead *bool     `json:"hasBeenRead,omitempty"`
}

type ListResponse struct {
	NewsList []newsResponse `json:"newsList"`
	Offset   int            `json:"offset"`
	Limit    int            `json:"limit"`
	Pages    int64          `json:"pages"`
	Total    int64          `json:"total"`
}

func parseNewList(newsList *[]models.News, userId *uuid.UUID) []newsResponse {
	var newsListValue []models.News
	if newsList != nil {
		newsListValue = *newsList
	}
	hasUserId := userId != nil
	newsListResponse := make([]newsResponse, len(newsListValue))
	for index, news := range newsListValue {
		var hasBeenRead *bool
		if hasUserId {
			hasNewsBeenRead := len(news.NewsReaders) == 1 && news.NewsReaders[0].UserId.String() == userId.String()
			hasBeenRead = &hasNewsBeenRead
		}

		newsListResponse[index] = newsResponse{
			ID:          news.ID.String(),
			Title:       news.Title,
			ImageUrl:    news.ImageUrl,
			Channel:     news.Channel,
			Url:         news.Url,
			PublishedAt: news.PublishedAt,
			HasBeenRead: hasBeenRead,
		}
	}

	return newsListResponse
}

func parseTotal(total *int64) int64 {
	var totalValue int64
	if total != nil {
		totalValue = *total
	}
	return totalValue
}

func NewListResponse(newsList *[]models.News, offset, limit int, total *int64, userId *uuid.UUID) *ListResponse {
	newsListResponse := parseNewList(newsList, userId)
	totalValue := parseTotal(total)
	pages := int64(math.Ceil(float64(*total) / float64(limit)))

	return &ListResponse{
		newsListResponse,
		offset,
		limit,
		pages,
		totalValue,
	}
}
