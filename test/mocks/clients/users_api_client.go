// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	clients "news-hub-microservices_news-api/internal/clients"

	mock "github.com/stretchr/testify/mock"
)

// UsersApiClient is an autogenerated mock type for the UsersApiClient type
type UsersApiClient struct {
	mock.Mock
}

// Get provides a mock function with given fields: token
func (_m *UsersApiClient) Get(token string) clients.GetResponse {
	ret := _m.Called(token)

	var r0 clients.GetResponse
	if rf, ok := ret.Get(0).(func(string) clients.GetResponse); ok {
		r0 = rf(token)
	} else {
		r0 = ret.Get(0).(clients.GetResponse)
	}

	return r0
}

type mockConstructorTestingTNewUsersApiClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewUsersApiClient creates a new instance of UsersApiClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUsersApiClient(t mockConstructorTestingTNewUsersApiClient) *UsersApiClient {
	mock := &UsersApiClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}