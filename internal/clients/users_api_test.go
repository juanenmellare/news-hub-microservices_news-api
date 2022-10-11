package clients

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewUsersApiClient(t *testing.T) {
	usersApiClient := NewUsersApiClient(&mockRestClient{}, "", "")

	assert.Implements(t, (*UsersApiClient)(nil), usersApiClient)
}

func Test_UsersApiClient_Get(t *testing.T) {
	token := "Bearer foo"

	restClientMock := &mockRestClient{}
	var responseBody GetResponse

	usersApiClient := NewUsersApiClient(restClientMock, "", "")

	response := usersApiClient.Get(token)

	assert.Equal(t, responseBody, response)
}

func Test_UsersApiClient_Get_error(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			assert.Equal(t, errors.New("foo-error"), r)
		} else {
			t.Errorf("did not panic")
		}
	}()
	token := "Bearer foo"

	restClientMock := &mockRestClientError{}

	usersApiClient := NewUsersApiClient(restClientMock, "", "")

	_ = usersApiClient.Get(token)
}
