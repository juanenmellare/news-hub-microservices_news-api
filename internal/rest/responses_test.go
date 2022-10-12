package rest

import (
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
	"math"
	"news-hub-microservices_news-api/internal/models"
	mocksModels "news-hub-microservices_news-api/test/mocks/models"
	"testing"
)

func TestNewListResponse(t *testing.T) {
	userId, _ := uuid.NewV4()
	newId2, _ := uuid.NewV4()

	news := mocksModels.NewNewsBuilder().AddReader(userId).Build()
	news2 := mocksModels.NewNewsBuilder().SetId(newId2).Build()

	newsList := []models.News{news, news2}
	offsetExpected := 0
	limitExpected := 20
	totalExpected := int64(85)
	pagesExpected := int64(math.Ceil(float64(totalExpected) / float64(limitExpected)))

	newsListResponse := NewListResponse(&newsList, offsetExpected, limitExpected, &totalExpected, &userId)

	assert.Equal(t, len(newsList), len(newsListResponse.NewsList))
	assert.Equal(t, offsetExpected, newsListResponse.Offset)
	assert.Equal(t, limitExpected, newsListResponse.Limit)
	assert.Equal(t, totalExpected, newsListResponse.Total)
	assert.Equal(t, pagesExpected, newsListResponse.Pages)

	assert.True(t, *newsListResponse.NewsList[0].HasBeenRead)
	assert.False(t, *newsListResponse.NewsList[1].HasBeenRead)
}
