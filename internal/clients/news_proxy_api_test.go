package clients

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewNewsProxyApiClient(t *testing.T) {
	newsProxyApiClient := NewNewsProxyApiClient(&mockRestClient{}, "", "")

	assert.Implements(t, (*NewsProxyApiClient)(nil), newsProxyApiClient)
}

func TestNewsProxyApiClient_GetChannelLatestNews(t *testing.T) {
	username := "admin"
	password := "password"
	defaultChannel := "channel"

	restClientMock := &mockRestClient{}
	var responseBody GetChannelLatestNewsResponse

	newsProxyApiClient := NewNewsProxyApiClient(restClientMock, username, password)

	response := newsProxyApiClient.GetChannelLatestNews(defaultChannel)

	assert.Equal(t, responseBody, response)
}

func TestNewsProxyApiClient_GetChannelLatestNews_error(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			assert.Equal(t, errors.New("foo-error"), r)
		} else {
			t.Errorf("did not panic")
		}
	}()
	username := "admin"
	password := "password"
	defaultChannel := "channel"

	restClientMock := &mockRestClientError{}

	newsProxyApiClient := NewNewsProxyApiClient(restClientMock, username, password)

	_ = newsProxyApiClient.GetChannelLatestNews(defaultChannel)
}

func Test_newsProxyApiClient_GetChannelsNames(t *testing.T) {
	username := "admin"
	password := "password"

	restClientMock := &mockRestClient{}
	var responseBody GetChannelsNamesResponse

	newsProxyApiClient := NewNewsProxyApiClient(restClientMock, username, password)

	response := newsProxyApiClient.GetChannelsNames()

	assert.Equal(t, responseBody, response)
}

func Test_newsProxyApiClient_GetChannelsNames_error(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			assert.Equal(t, errors.New("foo-error"), r)
		} else {
			t.Errorf("did not panic")
		}
	}()
	username := "admin"
	password := "password"

	restClientMock := &mockRestClientError{}

	newsProxyApiClient := NewNewsProxyApiClient(restClientMock, username, password)

	_ = newsProxyApiClient.GetChannelsNames()
}
