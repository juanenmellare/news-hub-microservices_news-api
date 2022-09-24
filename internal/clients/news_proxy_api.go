package clients

import (
	"fmt"
	restUtils "github.com/juanenmellare/gorequestbuilder"
	"time"
)

type NewsProxyApiClient interface {
	GetChannelsNames() GetChannelsNamesResponse
	GetChannelLatestNews(channel string) GetChannelLatestNewsResponse
}

type newsProxyApiClient struct {
	restClient RestClient
	username   string
	password   string
}

type GetChannelsNamesResponse struct {
	Channels []string `json:"channels"`
}

func (c newsProxyApiClient) GetChannelsNames() GetChannelsNamesResponse {
	path := "/v1/channels/"
	requestBuilder := restUtils.NewRequestBuilder().SetMethodGet().SetPath(path)

	var responseBody GetChannelsNamesResponse
	_, err := c.restClient.Call(requestBuilder, &responseBody)
	if err != nil {
		panic(err)
	}

	return responseBody
}

type NewsProxyApiNews struct {
	Title       string    `json:"title"`
	ImageUrl    string    `json:"image_url"`
	Channel     string    `json:"channel"`
	Url         string    `json:"url"`
	PublishedAt time.Time `json:"published_at"`
}

type GetChannelLatestNewsResponse struct {
	NewsList []NewsProxyApiNews `json:"news_list"`
}

func (c newsProxyApiClient) GetChannelLatestNews(channel string) GetChannelLatestNewsResponse {
	path := fmt.Sprintf("/v1/channels/%s/latest", channel)
	requestBuilder := restUtils.NewRequestBuilder().SetMethodGet().
		SetBasicAuthentication(c.username, c.password).SetPath(path)

	var responseBody GetChannelLatestNewsResponse
	_, err := c.restClient.Call(requestBuilder, &responseBody)
	if err != nil {
		panic(err)
	}

	return responseBody
}

func NewNewsProxyApiClient(restClient RestClient, username, password string) NewsProxyApiClient {
	return &newsProxyApiClient{
		restClient: restClient,
		username:   username,
		password:   password,
	}
}
