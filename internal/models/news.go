package models

import (
	"github.com/gofrs/uuid"
	"time"
)

type Channel string

const (
	ChannelInfobae Channel = "infobae"
)

type News struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	Title       string    `json:"title" gorm:"column:title"`
	ImageUrl    string    `json:"imageUrl" gorm:"column:image_url"`
	Channel     Channel   `json:"channel" gorm:"column:channel"`
	Url         string    `json:"url" gorm:"column:url"`
	PublishedAt time.Time `json:"publishedAt" gorm:"column:published_at"`
}

func NewNews(title, imageUrl, url string, channel Channel, publishedAt time.Time) *News {
	return &News{
		Title:       title,
		ImageUrl:    imageUrl,
		Channel:     channel,
		Url:         url,
		PublishedAt: publishedAt,
	}
}
