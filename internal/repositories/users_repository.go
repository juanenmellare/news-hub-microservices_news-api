package repositories

import (
	"news-hub-microservices_news-api/internal/databases"
	"news-hub-microservices_news-api/internal/models"
)

type NewsRepository interface {
	Create(user *models.News)
}

type newsRepositoryImpl struct {
	relationalDatabase databases.RelationalDatabase
}

func (u newsRepositoryImpl) Create(user *models.News) {
	result := u.relationalDatabase.Get().Create(user)
	if result.Error != nil {
		panic(result.Error)
	}
}

func NewNewsRepository(relationalDatabase databases.RelationalDatabase) NewsRepository {
	return &newsRepositoryImpl{
		relationalDatabase,
	}
}
