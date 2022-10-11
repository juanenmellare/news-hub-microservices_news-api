package api

import (
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"news-hub-microservices_news-api/configs"
	"news-hub-microservices_news-api/internal/clients"
	"news-hub-microservices_news-api/internal/errors"
	"news-hub-microservices_news-api/internal/factories"
	"strings"
)

func HandlePanicRecoveryMiddleware(context *gin.Context, i interface{}) {
	var apiError *errors.ApiError
	switch err := i.(type) {
	case *errors.ApiError:
		apiError = err
	case *errors.AlreadyExistModelError:
		apiError = errors.NewBadRequestApiError(err.Message)
	case *errors.InvalidEmailOrPasswordError:
		apiError = errors.NewBadRequestApiError(err.Message)
	case error:
		apiError = errors.NewInternalServerApiError(fmt.Sprintf("unexpected error: %v", err))
	default:
		apiError = errors.NewInternalServerApiError(fmt.Sprintf("unhandled error: %v", err))
	}
	context.JSON(apiError.Code, apiError)
}

func BasicAuthenticationMiddleware(c *gin.Context, candidateUsername, candidatePassword string) {
	authorizationHeaderKey := "Authorization"
	statusCode := http.StatusUnauthorized
	errResponse := errors.NewUnauthorizedApiError("invalid credentials")

	authorizationValue := c.Request.Header.Get(authorizationHeaderKey)
	if authorizationValue == "" {
		c.JSON(statusCode, errResponse)
		c.Abort()
		return
	}

	var authorizationValueBasicEncoded string
	authorizationValueParts := strings.Split(authorizationValue, ",")
	var authorizationValueBasicIndex int
	for index, value := range authorizationValueParts {
		valueTrimmed := strings.TrimSpace(authorizationValueParts[index])
		authorizationValueParts[index] = valueTrimmed
		if strings.Contains(value, "Basic ") {
			authorizationValueBasicEncoded = valueTrimmed
			authorizationValueBasicIndex = index
		}
	}
	if authorizationValueBasicEncoded == "" {
		c.JSON(statusCode, errResponse)
		c.Abort()
		return
	}

	authorizationValueBasic, err := base64.StdEncoding.DecodeString(strings.Split(
		strings.TrimSpace(authorizationValueBasicEncoded), " ")[1])
	if err != nil {
		c.JSON(statusCode, errResponse)
		c.Abort()
		return
	}

	authorizationValueBasicParts := strings.Split(strings.TrimSpace(string(authorizationValueBasic)), ":")
	if len(authorizationValueBasicParts) != 2 ||
		authorizationValueBasicParts[0] != candidateUsername ||
		authorizationValueBasicParts[1] != candidatePassword {
		c.JSON(statusCode, errResponse)
		c.Abort()
		return
	}

	authorizationValuePartsUpdated := append(authorizationValueParts[:authorizationValueBasicIndex],
		authorizationValueParts[authorizationValueBasicIndex+1:]...)

	c.Request.BasicAuth()
	c.Request.Header.Set(authorizationHeaderKey, strings.Join(authorizationValuePartsUpdated, ", "))
}

func UserAuthorizationMiddleware(c *gin.Context, userApiClient clients.UsersApiClient) {
	userIdHeaderKey := "X-User-Id"
	authorizationValue := c.Request.Header.Get("Authorization")
	if authorizationValue != "" {
		userData := userApiClient.Get(authorizationValue)
		c.Request.Header.Set(userIdHeaderKey, userData.Id)
	} else {
		c.Request.Header.Set(userIdHeaderKey, "")
	}
}

func NewRouter(layers factories.LayersFactory, config configs.Config) *gin.Engine {
	router := gin.Default()
	router.Use(gin.CustomRecovery(HandlePanicRecoveryMiddleware))

	router.GET("/ping", layers.GetHealthChecksController().Ping)

	newsController := layers.GetNewsController()
	v1 := router.Group("/v1")
	v1.Use(func(c *gin.Context) {
		BasicAuthenticationMiddleware(c, config.GetBasicAuthUsername(), config.GetBasicAuthPassword())
		UserAuthorizationMiddleware(c, layers.GetUsersApiClient())
	})
	{
		v1.GET("", newsController.List)
		v1.PUT("/:id", newsController.Read)
	}
	return router
}
