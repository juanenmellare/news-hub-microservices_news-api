package rest

import (
	"github.com/gin-gonic/gin"
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
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

func (r *ListRequest) MarshallAndValidate(context *gin.Context) {
	offset := queryParamToInt(context, "offset", "0")
	limit := queryParamToInt(context, "limit", "25")

	minLimit := 1
	maxLimit := 100
	if limit < minLimit || limit > maxLimit {
		panic(errors.NewOutOfRangeIntParamError("limit", minLimit, maxLimit))
	}

	r.Limit = limit
	r.Offset = offset
}
