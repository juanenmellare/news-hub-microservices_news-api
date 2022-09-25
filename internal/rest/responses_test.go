package rest

import (
	"github.com/stretchr/testify/assert"
	"math"
	"news-hub-microservices_news-api/internal/models"
	mocksModels "news-hub-microservices_news-api/test/mocks/models"
	"testing"
)

func TestNewListResponse(t *testing.T) {
	newsList := []models.News{mocksModels.NewNewsBuilder().Build()}
	offsetExpected := 0
	limitExpected := 25
	totalExpected := int64(85)
	pagesExpected := int64(math.Ceil(float64(totalExpected) / float64(limitExpected)))

	newsListResponse := NewListResponse(&newsList, offsetExpected, limitExpected, &totalExpected)

	assert.Equal(t, len(newsList), len(newsListResponse.NewsList))
	assert.Equal(t, offsetExpected, newsListResponse.Offset)
	assert.Equal(t, limitExpected, newsListResponse.Limit)
	assert.Equal(t, totalExpected, newsListResponse.Total)
	assert.Equal(t, pagesExpected, newsListResponse.Pages)
}
