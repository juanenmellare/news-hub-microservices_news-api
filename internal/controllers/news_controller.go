package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"news-hub-microservices_news-api/internal/rest"
	"news-hub-microservices_news-api/internal/services"
)

type NewsController interface {
	List(context *gin.Context)
	Read(context *gin.Context)
}

type newsController struct {
	NewsService services.NewsService
}

func (n newsController) List(context *gin.Context) {
	var request rest.ListRequest
	request.MarshallAndValidate(context)

	newsList := n.NewsService.List(request.Offset, request.Limit, request.UserId)
	total := n.NewsService.GetTotal()

	response := rest.NewListResponse(newsList, request.Offset, request.Limit, total, request.UserId)

	context.JSON(http.StatusOK, &response)
}

func (n newsController) Read(context *gin.Context) {
	var request rest.ReadRequest
	request.MarshallAndValidate(context)

	n.NewsService.AddReader(request.NewId, request.UserId)

	context.Done()
}

func NewNewsController(newsService services.NewsService) NewsController {
	return &newsController{
		newsService,
	}
}
