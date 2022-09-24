package services

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"news-hub-microservices_news-api/internal/clients"
	mocksClients "news-hub-microservices_news-api/test/mocks/clients"
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

func Test_newsService_FetchNews(t *testing.T) {
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

	newsService.FetchNews()
}

func Test_newsService_FetchNews_panic_GetChannelLatestNews_recover(t *testing.T) {
	assertShouldNotPanic(t)

	newsRepository := &mocksRepositories.NewsRepository{}
	newsRepository.On("CreateBulk", mock.Anything).Once()

	newsProxyApiClient := &mocksClients.NewsProxyApiClient{}
	newsProxyApiClient.On("GetChannelsNames").Return(clients.GetChannelsNamesResponse{Channels: []string{"Infobae", "TN"}})
	newsProxyApiClient.On("GetChannelLatestNews", mock.Anything).Panic("foo").Once()

	newsService := NewNewsService(newsRepository, newsProxyApiClient)

	newsService.FetchNews()
}
