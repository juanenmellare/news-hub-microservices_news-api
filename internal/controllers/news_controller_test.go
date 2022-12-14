package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"news-hub-microservices_news-api/internal/models"
	"news-hub-microservices_news-api/internal/services"
	modelsMocks "news-hub-microservices_news-api/test/mocks/models"
	servicesMocks "news-hub-microservices_news-api/test/mocks/services"
	"testing"
)

func Test_NewNewsController(t *testing.T) {
	var newsService services.NewsService

	assert.Implements(t, (*NewsController)(nil), NewNewsController(newsService))
}

func Test_newsController_List(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			str := fmt.Sprintf("the test should not panic: %v", r)
			t.Errorf(str)
		}
	}()

	newsMock := modelsMocks.NewNewsBuilder().Build()

	newsList := []models.News{newsMock}
	total := int64(len(newsList))
	var id *uuid.UUID

	var newsServices servicesMocks.NewsService
	newsServices.On("List", 0, 20, id).Return(&newsList)
	newsServices.On("GetTotal").Return(&total)

	controller := NewNewsController(&newsServices)

	writer := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(writer)

	context.Request, _ = http.NewRequest(http.MethodGet, "/", nil)

	controller.List(context)

	assert.Equal(t, http.StatusOK, writer.Code)
	assert.Equal(t, "{\"newsList\":[{\"id\":\"800d249f-a7f7-4129-a8a6-14d0cf9667e5\",\"title\":\"foo-title\",\"imageUrl\":\"foo-image-url\",\"channel\":\"foo-channel\",\"url\":\"foo-url\",\"publishedAt\":\"2022-10-02T01:01:01.000000001-03:00\"}],\"offset\":0,\"limit\":20,\"pages\":1,\"total\":1}", writer.Body.String())
}

func Test_newsController_Read(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			str := fmt.Sprintf("the test should not panic: %v", r)
			t.Errorf(str)
		}
	}()

	userId, _ := uuid.NewV4()
	newsId, _ := uuid.NewV4()

	var newsServices servicesMocks.NewsService
	newsServices.On("AddReader", newsId, userId).Return()

	controller := NewNewsController(&newsServices)

	writer := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(writer)

	context.Request, _ = http.NewRequest(http.MethodPut, "/", nil)
	context.Request.Header.Add("X-User-Id", userId.String())
	context.AddParam("id", newsId.String())

	controller.Read(context)

	assert.Equal(t, http.StatusOK, writer.Code)
}
