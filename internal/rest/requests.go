package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"news-hub-microservices_news-api/internal/errors"
	"strconv"
)

func queryParamToInt(context *gin.Context, param, defaultValue string) int {
	value := context.DefaultQuery(param, defaultValue)
	intValue, err := strconv.Atoi(value)
	if err != nil {
		panic(errors.NewIntQueryParamError(param))
	}
	return intValue
}

type Request interface {
	MarshallAndValidate(context *gin.Context)
}

type ListRequest struct {
	Limit  int
	Offset int
	UserId *uuid.UUID
}

func (r *ListRequest) MarshallAndValidate(context *gin.Context) {
	limit := queryParamToInt(context, "limit", "25")
	minLimit := 1
	maxLimit := 100
	if limit < minLimit || limit > maxLimit {
		panic(errors.NewOutOfRangeIntParamError("limit", minLimit, maxLimit))
	}
	r.Limit = limit

	offset := queryParamToInt(context, "offset", "0")
	r.Offset = offset

	xUserIdHeader := context.Request.Header.Get("X-User-Id")
	userId, err := uuid.FromString(xUserIdHeader)
	if err == nil {
		r.UserId = &userId
	} else {
		r.UserId = nil
	}
}

type ReadRequest struct {
	NewId  uuid.UUID
	UserId uuid.UUID
}

func (r *ReadRequest) MarshallAndValidate(context *gin.Context) {
	xUserIdHeader := context.Request.Header.Get("X-User-Id")
	userId, err := uuid.FromString(xUserIdHeader)
	if err != nil {
		panic(errors.NewBadRequestApiError("missing or invalid token"))
	}
	r.UserId = userId

	newsIdParam := context.Param("id")
	newsId, err := uuid.FromString(newsIdParam)
	if err != nil {
		panic(errors.NewBadRequestApiError("invalid newsId"))
	}
	r.NewId = newsId
}
