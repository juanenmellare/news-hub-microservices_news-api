package rest

import (
	"github.com/gin-gonic/gin"
	"news-hub-microservices_news-api/internal/errors"
)

func validateStringField(fieldName string, field *string, notValidFields *[]string) {
	if field == nil || *field == "" {
		*notValidFields = append(*notValidFields, fieldName)
	}
}

func validateNotValidFieldsSlice(notValidFields []string) {
	if len(notValidFields) > 0 {
		panic(errors.NewRequestFieldsShouldNotBeEmptyError(notValidFields))
	}
}

func marshallRequestBody(context *gin.Context, i interface{}) {
	if err := context.BindJSON(&i); err != nil {
		panic(errors.NewBadRequestApiError(err.Error()))
	}
}

type Request interface {
	MarshallAndValidate(context *gin.Context)
}
