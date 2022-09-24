package models

import (
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestUser(t *testing.T) {
	user := &News{}

	var emptyString string
	assert.Equal(t, uuid.UUID{}, user.ID)
	assert.Equal(t, emptyString, user.Title)
	assert.Equal(t, emptyString, user.ImageUrl)
	assert.Equal(t, emptyString, user.Channel)
	assert.Equal(t, emptyString, user.Url)
}

func TestNewUser(t *testing.T) {
	imageUrl := "foo-image-url"
	channel := "foo-channel"
	title := "foo-title"
	url := "foo-url"
	publishedAt := time.Now()

	news := NewNews(title, imageUrl, url, channel, publishedAt)

	assert.Equal(t, title, news.Title)
	assert.Equal(t, imageUrl, news.ImageUrl)
	assert.Equal(t, channel, news.Channel)
	assert.Equal(t, url, news.Url)
	assert.Equal(t, publishedAt, news.PublishedAt)
}
