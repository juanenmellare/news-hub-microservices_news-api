package services

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	mocks "news-hub-microservices_news-api/test/mocks/repositories"
	"testing"
)

var bCryptCost = bcrypt.DefaultCost

func Test_NewNewsService(t *testing.T) {
	var newsRepository mocks.NewsRepository

	assert.Implements(t, (*NewsService)(nil), NewNewsService(&newsRepository))
}

func assertShouldNotPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			str := fmt.Sprintf("the test should not panic: %v", r)
			t.Errorf(str)
		}
	}()
}
