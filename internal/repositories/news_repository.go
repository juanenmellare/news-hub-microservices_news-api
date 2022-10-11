package repositories

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm/clause"
	"news-hub-microservices_news-api/internal/databases"
	"news-hub-microservices_news-api/internal/models"
)

type NewsRepository interface {
	CreateBulk(newsList []models.News)
	FindAll(offset, limit int, userId *uuid.UUID) *[]models.News
	GetTotal() *int64
	FindById(id string) *models.News
	AddNewsReader(reader *models.NewsReader)
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

func (u newsRepository) FindAll(offset, limit int, userId *uuid.UUID) *[]models.News {
	var newsList []models.News
	var userIdValue string
	if userId != nil {
		userIdValue = userId.String()
	}
	tx := u.relationalDatabase.Get().
		Order("published_at DESC").
		Offset(offset).
		Limit(limit).
		Preload("NewsReaders", "\"news_readers\".\"user_id\" = ?", userIdValue).
		Find(&newsList)
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

func (u newsRepository) FindById(id string) *models.News {
	var news models.News
	result := u.relationalDatabase.Get().First(&news, "id = ?", id)
	if result.Error != nil {
		return nil
	}
	return &news
}

func (u newsRepository) AddNewsReader(reader *models.NewsReader) {
	result := u.relationalDatabase.Get().Clauses(clause.OnConflict{DoNothing: true}).Create(reader)
	if result.Error != nil {
		panic(result.Error)
	}
}

func NewNewsRepository(relationalDatabase databases.RelationalDatabase) NewsRepository {
	return &newsRepository{
		relationalDatabase,
	}
}
