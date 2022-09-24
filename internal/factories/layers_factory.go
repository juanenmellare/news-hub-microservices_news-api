package factories

import (
	"net/http"
	"news-hub-microservices_news-api/configs"
	"news-hub-microservices_news-api/internal/clients"
	"news-hub-microservices_news-api/internal/controllers"
	"news-hub-microservices_news-api/internal/databases"
	"news-hub-microservices_news-api/internal/repositories"
	"news-hub-microservices_news-api/internal/services"
)

func buildHealthChecksController() controllers.HealthChecksController {
	return controllers.NewHealthChecksController()
}

func buildNewsService(relationalDatabase databases.RelationalDatabase, config configs.Config) services.NewsService {
	newsRepository := repositories.NewNewsRepository(relationalDatabase)

	restClient := clients.NewRestClient(config.GetNewProxyApiBaseUrl(), &http.Client{})
	newsProxyApiClient := clients.NewNewsProxyApiClient(restClient, config.GetNewProxyApiUsername(), config.GetNewProxyApiPassword())

	newsService := services.NewNewsService(newsRepository, newsProxyApiClient)

	return newsService
}

func buildNewsController(newsService services.NewsService) controllers.NewsController {
	newsController := controllers.NewNewsController(newsService)

	return newsController
}

type LayersFactory interface {
	GetHealthChecksController() controllers.HealthChecksController
	GetNewsService() services.NewsService
	GetNewsController() controllers.NewsController
}

type layersFactory struct {
	healthChecksController controllers.HealthChecksController
	newsController         controllers.NewsController
	newsService            services.NewsService
}

func NewLayersFactory(relationalDatabase databases.RelationalDatabase, config configs.Config) LayersFactory {
	newsService := buildNewsService(relationalDatabase, config)
	return &layersFactory{
		healthChecksController: buildHealthChecksController(),
		newsService:            newsService,
		newsController:         buildNewsController(newsService),
	}
}

func (c layersFactory) GetHealthChecksController() controllers.HealthChecksController {
	return c.healthChecksController
}

func (c layersFactory) GetNewsService() services.NewsService {
	return c.newsService
}

func (c layersFactory) GetNewsController() controllers.NewsController {
	return c.newsController
}
