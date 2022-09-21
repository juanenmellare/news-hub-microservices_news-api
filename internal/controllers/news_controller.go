package controllers

import (
	"news-hub-microservices_news-api/internal/services"
)

type NewsController interface {
}

type newsControllerImpl struct {
	NewsService services.NewsService
}

func NewNewsController(newsService services.NewsService) NewsController {
	return &newsControllerImpl{
		newsService,
	}
}
