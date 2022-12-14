// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// Token is an autogenerated mock type for the Token type
type Token struct {
	mock.Mock
}

// ToString provides a mock function with given fields: userTokenSecretKey
func (_m *Token) ToString(userTokenSecretKey string) string {
	ret := _m.Called(userTokenSecretKey)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(userTokenSecretKey)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Verify provides a mock function with given fields: userTokenSecretKey, r
func (_m *Token) Verify(userTokenSecretKey string, r *http.Request) {
	_m.Called(userTokenSecretKey, r)
}

type mockConstructorTestingTNewToken interface {
	mock.TestingT
	Cleanup(func())
}

// NewToken creates a new instance of Token. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewToken(t mockConstructorTestingTNewToken) *Token {
	mock := &Token{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
