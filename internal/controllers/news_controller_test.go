package controllers

import (
	"github.com/stretchr/testify/assert"
	"news-hub-microservices_news-api/internal/services"
	"testing"
)

func Test_NewNewsController(t *testing.T) {
	var newsService services.NewsService

	assert.Implements(t, (*NewsController)(nil), NewNewsController(newsService))
}
