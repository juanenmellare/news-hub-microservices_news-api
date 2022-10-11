package clients

import (
	"encoding/json"
	"errors"
	"fmt"
	restUtils "github.com/juanenmellare/gorequestbuilder"
	"io"
	"io/ioutil"
	"net/http"
)

type mockRestClient struct{}

func (r mockRestClient) Call(_ restUtils.RequestBuilder, _ interface{}) (*http.Response, error) {
	return &http.Response{}, nil
}

type mockRestClientError struct{}

func (r mockRestClientError) Call(_ restUtils.RequestBuilder, _ interface{}) (*http.Response, error) {
	return &http.Response{}, errors.New("foo-error")
}

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type RestClient interface {
	Call(requestBuilder restUtils.RequestBuilder, responseObject interface{}) (*http.Response, error)
}

type restClient struct {
	baseURL string
	client  HttpClient
}

func NewRestClient(baseURL string, client HttpClient) RestClient {
	return &restClient{
		baseURL: baseURL,
		client:  client,
	}
}

func (r restClient) Call(requestBuilder restUtils.RequestBuilder, responseBody interface{}) (
	*http.Response, error) {
	request, err := requestBuilder.SetBaseURL(r.baseURL).Build()
	if err != nil {
		return nil, errors.New("Build request error: " + err.Error())
	}

	response, err := r.client.Do(request)
	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			fmt.Printf("Close body response error: " + err.Error())
		}
	}(response.Body)

	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	if response.StatusCode < 300 {
		err = json.Unmarshal(bodyBytes, &responseBody)
		if err != nil {
			return nil, err
		}
	} else {
		err = errors.New(string(bodyBytes))
		return response, err
	}

	return response, nil
}
