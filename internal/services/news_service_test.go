package services

import (
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"news-hub-microservices_news-api/internal/clients"
	"news-hub-microservices_news-api/internal/models"
	mocksClients "news-hub-microservices_news-api/test/mocks/clients"
	mocksModels "news-hub-microservices_news-api/test/mocks/models"
	mocksRepositories "news-hub-microservices_news-api/test/mocks/repositories"
	"testing"
	"time"
)

func assertShouldNotPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			str := fmt.Sprintf("the test should not panic: %v", r)
			t.Errorf(str)
		}
	}()
}

func Test_NewNewsService(t *testing.T) {
	var newsRepository mocksRepositories.NewsRepository
	var newsProxyApiNews mocksClients.NewsProxyApiClient

	assert.Implements(t, (*NewsService)(nil), NewNewsService(&newsRepository, &newsProxyApiNews))
}

func Test_newsService_Fetch(t *testing.T) {
	assertShouldNotPanic(t)

	newsRepository := &mocksRepositories.NewsRepository{}
	newsRepository.On("CreateBulk", mock.Anything).Once()

	newsProxyApiClient := &mocksClients.NewsProxyApiClient{}
	newsProxyApiClient.On("GetChannelsNames").Return(clients.GetChannelsNamesResponse{Channels: []string{"Infobae", "TN"}})
	newsProxyApiClient.On("GetChannelLatestNews", mock.Anything).
		Return(clients.GetChannelLatestNewsResponse{
			NewsList: []clients.NewsProxyApiNews{{Title: "foo-title", Url: "foo-url", ImageUrl: "foo-image-url", Channel: "Infobae", PublishedAt: time.Now()}},
		}).Once()
	newsProxyApiClient.On("GetChannelLatestNews", mock.Anything).
		Return(clients.GetChannelLatestNewsResponse{
			NewsList: []clients.NewsProxyApiNews{{Title: "foo-title", Url: "foo-url", ImageUrl: "foo-image-url", Channel: "TN", PublishedAt: time.Now()}},
		}).Once()

	newsService := NewNewsService(newsRepository, newsProxyApiClient)

	newsService.Fetch()
}

func Test_newsService_Fetch_panic_GetChannelLatestNews_recover(t *testing.T) {
	assertShouldNotPanic(t)

	newsRepository := &mocksRepositories.NewsRepository{}
	newsRepository.On("CreateBulk", mock.Anything).Once()

	newsProxyApiClient := &mocksClients.NewsProxyApiClient{}
	newsProxyApiClient.On("GetChannelsNames").Return(clients.GetChannelsNamesResponse{Channels: []string{"Infobae", "TN"}})
	newsProxyApiClient.On("GetChannelLatestNews", mock.Anything).Panic("foo").Once()

	newsService := NewNewsService(newsRepository, newsProxyApiClient)

	newsService.Fetch()
}

func Test_newsService_List(t *testing.T) {
	offset := 0
	limit := 20
	var userId *uuid.UUID

	newsMock := mocksModels.NewNewsBuilder().Build()
	newsListExpected := &[]models.News{newsMock}
	newsRepository := &mocksRepositories.NewsRepository{}
	newsRepository.On("FindAll", offset, limit, userId).Return(newsListExpected)

	newsService := NewNewsService(newsRepository, nil)
	newsList := newsService.List(offset, limit, userId)

	assert.Equal(t, newsListExpected, newsList)
}

func Test_newsService_Count(t *testing.T) {
	totalExpected := int64(10)

	newsRepository := &mocksRepositories.NewsRepository{}
	newsRepository.On("GetTotal").Return(&totalExpected)

	newsService := NewNewsService(newsRepository, nil)
	total := newsService.GetTotal()

	assert.Equal(t, &totalExpected, total)
}

func Test_newsService_AddReader(t *testing.T) {
	newsId, _ := uuid.NewV4()
	userId, _ := uuid.NewV4()

	newsReader := models.NewsReader{NewsId: newsId, UserId: userId}

	newsRepository := &mocksRepositories.NewsRepository{}
	newsRepository.On("AddNewsReader", &newsReader).Return()

	newsService := NewNewsService(newsRepository, nil)
	newsService.AddReader(newsId, userId)
}
