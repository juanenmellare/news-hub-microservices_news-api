package services

import (
	"news-hub-microservices_news-api/internal/repositories"
)

type NewsService interface {
}

type newsServiceImpl struct {
	newsRepository repositories.NewsRepository
}

func NewNewsService(newsRepository repositories.NewsRepository) NewsService {
	return &newsServiceImpl{
		newsRepository,
	}
}
