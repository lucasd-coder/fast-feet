// Code generated by mockery v2.20.2. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Execute_internal_test is an autogenerated mock type for the Execute type
type Execute_internal_test struct {
	mock.Mock
}

// Execute provides a mock function with given fields: dto
func (_m *Execute_internal_test) Execute(dto string) {
	_m.Called(dto)
}

type mockConstructorTestingTNewExecute_internal_test interface {
	mock.TestingT
	Cleanup(func())
}

// NewExecute_internal_test creates a new instance of Execute_internal_test. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewExecute_internal_test(t mockConstructorTestingTNewExecute_internal_test) *Execute_internal_test {
	mock := &Execute_internal_test{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
