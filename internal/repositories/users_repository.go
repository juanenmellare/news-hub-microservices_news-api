package repositories

import (
	"gorm.io/gorm/clause"
	"news-hub-microservices_news-api/internal/databases"
	"news-hub-microservices_news-api/internal/models"
)

type NewsRepository interface {
	Create(news *models.News)
	CreateBulk(newsList []models.News)
}

type newsRepository struct {
	relationalDatabase databases.RelationalDatabase
}

func (u newsRepository) Create(user *models.News) {
	result := u.relationalDatabase.Get().Create(user)
	if result.Error != nil {
		panic(result.Error)
	}
}

func (u newsRepository) CreateBulk(users []models.News) {
	result := u.relationalDatabase.Get().Clauses(clause.OnConflict{DoNothing: true}).Create(users)
	if result.Error != nil {
		panic(result.Error)
	}
}

func NewNewsRepository(relationalDatabase databases.RelationalDatabase) NewsRepository {
	return &newsRepository{
		relationalDatabase,
	}
}
