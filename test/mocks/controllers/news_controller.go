// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// NewsController is an autogenerated mock type for the NewsController type
type NewsController struct {
	mock.Mock
}

type mockConstructorTestingTNewNewsController interface {
	mock.TestingT
	Cleanup(func())
}

// NewNewsController creates a new instance of NewsController. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewNewsController(t mockConstructorTestingTNewNewsController) *NewsController {
	mock := &NewsController{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}