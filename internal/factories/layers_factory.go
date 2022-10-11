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

func buildUserApiClient(config configs.Config) clients.UsersApiClient {
	restClient := clients.NewRestClient(config.GetUsersApiBaseUrl(), &http.Client{})
	userApiClient := clients.NewUsersApiClient(restClient, config.GetUsersApiUsername(), config.GetUsersApiPassword())
	return userApiClient
}

func buildNewsService(relationalDatabase databases.RelationalDatabase, config configs.Config) services.NewsService {
	newsRepository := repositories.NewNewsRepository(relationalDatabase)

	restClient := clients.NewRestClient(config.GetNewsProxyApiBaseUrl(), &http.Client{})
	newsProxyApiClient := clients.NewNewsProxyApiClient(restClient, config.GetNewsProxyApiUsername(), config.GetNewsProxyApiPassword())

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
	GetUsersApiClient() clients.UsersApiClient
}

type layersFactory struct {
	healthChecksController controllers.HealthChecksController
	newsController         controllers.NewsController
	newsService            services.NewsService
	userApiClient          clients.UsersApiClient
}

func NewLayersFactory(relationalDatabase databases.RelationalDatabase, config configs.Config) LayersFactory {
	newsService := buildNewsService(relationalDatabase, config)
	return &layersFactory{
		healthChecksController: buildHealthChecksController(),
		newsService:            newsService,
		newsController:         buildNewsController(newsService),
		userApiClient:          buildUserApiClient(config),
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

func (c layersFactory) GetUsersApiClient() clients.UsersApiClient {
	return c.userApiClient
}
