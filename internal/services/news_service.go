package services

import (
	"news-hub-microservices_news-api/internal/clients"
	"news-hub-microservices_news-api/internal/models"
	"news-hub-microservices_news-api/internal/repositories"
	"news-hub-microservices_news-api/internal/utils"
	"sync"
)

type NewsService interface {
	Fetch()
	List(offset, limit int) *[]models.News
	GetTotal() *int64
}

type newsService struct {
	newsRepository  repositories.NewsRepository
	newsProxyClient clients.NewsProxyApiClient
}

func (n newsService) getNewsFromChannel(channel string) []models.News {
	channelNewList := n.newsProxyClient.GetChannelLatestNews(channel).NewsList

	newsList := make([]models.News, len(channelNewList))
	for index, news := range channelNewList {
		newsList[index] = models.News{
			Title:       news.Title,
			ImageUrl:    news.ImageUrl,
			Channel:     news.Channel,
			Url:         news.Url,
			PublishedAt: news.PublishedAt,
		}
	}

	return newsList
}

func (n newsService) queueNewsFromChannel(channel string, queue chan []models.News) {
	defer utils.RecoverGoRoutineWithHandler("Fetch.queueNewsFromChannel", func() {
		queue <- make([]models.News, 0)
	})
	queue <- n.getNewsFromChannel(channel)
}

func (n newsService) Fetch() {
	getChannelsNamesResponse := n.newsProxyClient.GetChannelsNames()
	channels := getChannelsNamesResponse.Channels
	channelsLen := len(channels)
	queue := make(chan []models.News, channelsLen)

	var newsList []models.News
	var wg sync.WaitGroup
	wg.Add(channelsLen)
	for _, channel := range channels {
		go n.queueNewsFromChannel(channel, queue)
	}
	go func() {
		defer utils.RecoverGoRoutine("Fetch.appendNewList")
		for newsListToAppend := range queue {
			newsList = append(newsList, newsListToAppend...)
			wg.Done()
		}
	}()
	wg.Wait()

	if len(newsList) > 0 {
		n.newsRepository.CreateBulk(newsList)
	}
}

func (n newsService) List(offset, limit int) *[]models.News {
	return n.newsRepository.List(offset, limit)
}

func (n newsService) GetTotal() *int64 {
	return n.newsRepository.GetTotal()
}

func NewNewsService(newsRepository repositories.NewsRepository, newsProxyClient clients.NewsProxyApiClient) NewsService {
	return &newsService{
		newsRepository,
		newsProxyClient,
	}
}
