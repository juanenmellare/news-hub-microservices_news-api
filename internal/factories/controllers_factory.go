package factories

import (
	"news-hub-microservices_news-api/internal/controllers"
	"news-hub-microservices_news-api/internal/databases"
	"news-hub-microservices_news-api/internal/repositories"
	"news-hub-microservices_news-api/internal/services"
)

func buildHealthChecksController() controllers.HealthChecksController {
	return controllers.NewHealthChecksController()
}

func buildNewsController(relationalDatabase databases.RelationalDatabase) controllers.NewsController {
	newsRepository := repositories.NewNewsRepository(relationalDatabase)
	newsService := services.NewNewsService(newsRepository)
	newsController := controllers.NewNewsController(newsService)

	return newsController
}

type ControllersFactory interface {
	GetHealthChecksController() controllers.HealthChecksController
	GetNewsController() controllers.NewsController
}

type controllersFactoryImpl struct {
	healthChecksController controllers.HealthChecksController
	newsController         controllers.NewsController
}

func NewControllersFactory(relationalDatabase databases.RelationalDatabase) ControllersFactory {
	return &controllersFactoryImpl{
		healthChecksController: buildHealthChecksController(),
		newsController:         buildNewsController(relationalDatabase),
	}
}

func (c controllersFactoryImpl) GetHealthChecksController() controllers.HealthChecksController {
	return c.healthChecksController
}

func (c controllersFactoryImpl) GetNewsController() controllers.NewsController {
	return c.newsController
}
