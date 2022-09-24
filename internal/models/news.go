package models

import (
	"github.com/gofrs/uuid"
	"time"
)

type News struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	Title       string    `json:"title" gorm:"column:title;unique;not null"`
	ImageUrl    string    `json:"imageUrl" gorm:"column:image_url"`
	Channel     string    `json:"channel" gorm:"column:channel"`
	Url         string    `json:"url" gorm:"column:url;unique;not null"`
	PublishedAt time.Time `json:"publishedAt" gorm:"column:published_at"`
}

func NewNews(title, imageUrl, url, channel string, publishedAt time.Time) *News {
	return &News{
		Title:       title,
		ImageUrl:    imageUrl,
		Channel:     channel,
		Url:         url,
		PublishedAt: publishedAt,
	}
}
