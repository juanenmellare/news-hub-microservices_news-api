package clients

import (
	"fmt"
	restUtils "github.com/juanenmellare/gorequestbuilder"
	"net/http"
)

type UsersApiClient interface {
	Get(token string) GetResponse
}

type usersApiClient struct {
	restClient RestClient
	username   string
	password   string
}

type GetResponse struct {
	Id        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

func (c usersApiClient) Get(authorizationValue string) GetResponse {
	path := "/v1"

	request := &http.Request{Header: map[string][]string{}}
	request.SetBasicAuth(c.username, c.password)
	basicAuth := request.Header.Get("Authorization")
	authValue := fmt.Sprintf("%s, %s", basicAuth, authorizationValue)

	requestBuilder := restUtils.NewRequestBuilder().SetMethodGet().
		AddHeader("Authorization", authValue).SetPath(path)

	var responseBody GetResponse
	_, err := c.restClient.Call(requestBuilder, &responseBody)
	if err != nil {
		panic(err)
	}

	return responseBody
}

func NewUsersApiClient(restClient RestClient, username, password string) UsersApiClient {
	return &usersApiClient{
		restClient: restClient,
		username:   username,
		password:   password,
	}
}
