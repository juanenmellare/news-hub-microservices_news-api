package repositories

import (
	"gorm.io/gorm/clause"
	"news-hub-microservices_news-api/internal/databases"
	"news-hub-microservices_news-api/internal/models"
)

type NewsRepository interface {
	CreateBulk(newsList []models.News)
	List(offset, limit int) *[]models.News
	GetTotal() *int64
}

type newsRepository struct {
	relationalDatabase databases.RelationalDatabase
}

func (u newsRepository) CreateBulk(users []models.News) {
	result := u.relationalDatabase.Get().Clauses(clause.OnConflict{DoNothing: true}).Create(users)
	if result.Error != nil {
		panic(result.Error)
	}
}

func (u newsRepository) List(offset, limit int) *[]models.News {
	var newsList []models.News
	tx := u.relationalDatabase.Get().Offset(offset).Limit(limit).Find(&newsList)
	if err := tx.Error; err != nil {
		panic(err)
	}

	return &newsList
}

func (u newsRepository) GetTotal() *int64 {
	var count int64
	tx := u.relationalDatabase.Get().Find(&models.News{}).Count(&count)
	if err := tx.Error; err != nil {
		panic(err)
	}

	return &count
}

func NewNewsRepository(relationalDatabase databases.RelationalDatabase) NewsRepository {
	return &newsRepository{
		relationalDatabase,
	}
}
