package clients

import (
	"errors"
	restUtils "github.com/juanenmellare/gorequestbuilder"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

type MockResponseBody struct {
	Message string `json:"message"`
}

type ClientMock struct{}

func (c *ClientMock) Do(_ *http.Request) (*http.Response, error) {
	return &http.Response{
		StatusCode: http.StatusOK,
		Body:       ioutil.NopCloser(strings.NewReader(`{ "message": "foo-message" }`)),
	}, nil
}

type ClientMockBadRequest struct{}

func (c *ClientMockBadRequest) Do(_ *http.Request) (*http.Response, error) {
	return &http.Response{
		StatusCode: http.StatusBadRequest,
		Body:       ioutil.NopCloser(strings.NewReader(`{ "message": "foo-message" }`)),
	}, nil
}

type ClientMockDoError struct{}

func (c *ClientMockDoError) Do(_ *http.Request) (*http.Response, error) {
	return nil, errors.New("foo-error")
}

type ClientMockWithBrokenJson struct{}

func (c *ClientMockWithBrokenJson) Do(_ *http.Request) (*http.Response, error) {
	return &http.Response{
		StatusCode: http.StatusOK,
		Body:       ioutil.NopCloser(strings.NewReader("0")),
	}, nil
}

func TestNewRestClient(t *testing.T) {
	assert.Implements(t, (*RestClient)(nil), NewRestClient("foo-base-url", &ClientMock{}))
}

func Test_restClient_Call(t *testing.T) {
	restClient := NewRestClient("foo", &ClientMock{})
	requestBuilder := restUtils.NewRequestBuilder().SetMethodGet()
	var mockResponseBody MockResponseBody
	response, err := restClient.Call(requestBuilder, &mockResponseBody)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, "foo-message", mockResponseBody.Message)
}

func Test_restClient_Call_no_response_body(t *testing.T) {
	restClient := NewRestClient("foo", &ClientMock{})
	requestBuilder := restUtils.NewRequestBuilder().SetMethodGet()

	response, err := restClient.Call(requestBuilder, nil)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)
}

func Test_restClient_Call_err_gorequestbuilder(t *testing.T) {
	restClient := NewRestClient("", &ClientMock{})
	requestBuilder := restUtils.NewRequestBuilder().SetMethodGet()

	response, err := restClient.Call(requestBuilder, nil)

	assert.Equal(t, "Build request error: base URL is not defined", err.Error())
	assert.Nil(t, response)
}

func Test_restClient_Call_err_ClientMockWithBrokenJson(t *testing.T) {
	restClient := NewRestClient("foo", &ClientMockWithBrokenJson{})
	requestBuilder := restUtils.NewRequestBuilder().SetMethodGet()
	var mockResponseBody MockResponseBody
	response, err := restClient.Call(requestBuilder, &mockResponseBody)

	assert.Equal(t, "json: cannot unmarshal number into Go value of type clients.MockResponseBody", err.Error())
	assert.Nil(t, response)
}

func Test_restClient_Call_err_ClientMockBadRequest(t *testing.T) {
	restClient := NewRestClient("foo", &ClientMockBadRequest{})
	requestBuilder := restUtils.NewRequestBuilder().SetMethodGet()
	var mockResponseBody MockResponseBody
	response, err := restClient.Call(requestBuilder, &mockResponseBody)

	assert.Equal(t, "{ \"message\": \"foo-message\" }", err.Error())
	assert.Equal(t, http.StatusBadRequest, response.StatusCode)
}

func Test_restClient_Call_err_ClientMockDoError(t *testing.T) {
	restClient := NewRestClient("foo", &ClientMockDoError{})
	requestBuilder := restUtils.NewRequestBuilder().SetMethodGet()
	var mockResponseBody MockResponseBody
	response, err := restClient.Call(requestBuilder, &mockResponseBody)

	assert.Equal(t, "foo-error", err.Error())
	assert.Nil(t, response)
}
