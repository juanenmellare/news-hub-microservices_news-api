package rest

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"news-hub-microservices_news-api/internal/errors"
	"testing"
)

func TestListRequest_MarshallAndValidate_default_param_values(t *testing.T) {
	writer := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(writer)

	context.Request, _ = http.NewRequest(http.MethodGet, "/v1", nil)

	var request ListRequest
	request.MarshallAndValidate(context)

	assert.Equal(t, 0, request.Offset)
	assert.Equal(t, 25, request.Limit)
	assert.Nil(t, request.UserId)
}

func TestListRequest_MarshallAndValidate_set_userId(t *testing.T) {
	writer := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(writer)

	context.Request, _ = http.NewRequest(http.MethodGet, "/v1", nil)
	userId, _ := uuid.NewV4()
	context.Request.Header.Add("X-User-ID", userId.String())

	var request ListRequest
	request.MarshallAndValidate(context)

	assert.Equal(t, userId.String(), request.UserId.String())
}

func TestListRequest_MarshallAndValidate_set_param_values(t *testing.T) {
	writer := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(writer)

	expectedOffset := 45
	expectedLimit := 1

	url := fmt.Sprintf("/v1?limit=%d&offset=%d", expectedLimit, expectedOffset)
	context.Request, _ = http.NewRequest(http.MethodGet, url, nil)

	var request ListRequest
	request.MarshallAndValidate(context)

	assert.Equal(t, expectedOffset, request.Offset)
	assert.Equal(t, expectedLimit, request.Limit)
}

func TestListRequest_MarshallAndValidate_error_param_not_int(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			assert.Equal(t, &errors.ApiError{Code: 400, Status: "Bad Request", Message: "the param 'offset' should be an int value"}, r)
		} else {
			t.Errorf("did not panic")
		}
	}()

	writer := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(writer)

	expectedOffset := "a"
	expectedLimit := 100

	url := fmt.Sprintf("/v1?limit=%d&offset=%s", expectedLimit, expectedOffset)
	context.Request, _ = http.NewRequest(http.MethodGet, url, nil)

	var request ListRequest
	request.MarshallAndValidate(context)
}

func TestListRequest_MarshallAndValidate_error_limit_off(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			assert.Equal(t, &errors.ApiError{Code: 400, Status: "Bad Request", Message: "the param 'limit' should be between 1 and 100"}, r)
		} else {
			t.Errorf("did not panic")
		}
	}()

	writer := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(writer)

	expectedOffset := 0
	expectedLimit := 101

	url := fmt.Sprintf("/v1?limit=%d&offset=%d", expectedLimit, expectedOffset)
	context.Request, _ = http.NewRequest(http.MethodGet, url, nil)

	var request ListRequest
	request.MarshallAndValidate(context)
}

func TestReadRequest_MarshallAndValidate_default_param_values(t *testing.T) {
	writer := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(writer)

	context.Request, _ = http.NewRequest(http.MethodPut, "/v1", nil)

	userId, _ := uuid.NewV4()
	context.Request.Header.Add("X-User-Id", userId.String())

	newsId, _ := uuid.NewV4()
	context.AddParam("id", newsId.String())

	var request ReadRequest
	request.MarshallAndValidate(context)

	assert.Equal(t, newsId.String(), request.NewId.String())
	assert.Equal(t, userId.String(), request.UserId.String())
}

func TestReadRequest_MarshallAndValidate_panic_userId(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			assert.Equal(t, &errors.ApiError{Code: 400, Status: "Bad Request", Message: "missing or invalid token"}, r)
		} else {
			t.Errorf("did not panic")
		}
	}()

	writer := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(writer)

	context.Request, _ = http.NewRequest(http.MethodPut, "/v1", nil)

	var request ReadRequest
	request.MarshallAndValidate(context)
}

func TestReadRequest_MarshallAndValidate_panic_newsId(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			assert.Equal(t, &errors.ApiError{Code: 400, Status: "Bad Request", Message: "invalid newsId"}, r)
		} else {
			t.Errorf("did not panic")
		}
	}()

	writer := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(writer)

	context.Request, _ = http.NewRequest(http.MethodPut, "/v1", nil)

	userId, _ := uuid.NewV4()
	context.Request.Header.Add("X-User-Id", userId.String())

	var request ReadRequest
	request.MarshallAndValidate(context)
}
