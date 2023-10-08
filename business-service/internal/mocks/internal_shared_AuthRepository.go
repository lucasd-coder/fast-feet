// Code generated by mockery v2.20.2. DO NOT EDIT.

package mocks

import (
	context "context"

	shared "github.com/lucasd-coder/fast-feet/business-service/internal/shared"
	mock "github.com/stretchr/testify/mock"
)

// AuthRepository_internal_shared is an autogenerated mock type for the AuthRepository type
type AuthRepository_internal_shared struct {
	mock.Mock
}

// FindByEmail provides a mock function with given fields: ctx, email
func (_m *AuthRepository_internal_shared) FindByEmail(ctx context.Context, email string) (*shared.GetUserResponse, error) {
	ret := _m.Called(ctx, email)

	var r0 *shared.GetUserResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*shared.GetUserResponse, error)); ok {
		return rf(ctx, email)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *shared.GetUserResponse); ok {
		r0 = rf(ctx, email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*shared.GetUserResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindRolesByID provides a mock function with given fields: ctx, ID
func (_m *AuthRepository_internal_shared) FindRolesByID(ctx context.Context, ID string) (*shared.GetRolesResponse, error) {
	ret := _m.Called(ctx, ID)

	var r0 *shared.GetRolesResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*shared.GetRolesResponse, error)); ok {
		return rf(ctx, ID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *shared.GetRolesResponse); ok {
		r0 = rf(ctx, ID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*shared.GetRolesResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsActiveUser provides a mock function with given fields: ctx, ID
func (_m *AuthRepository_internal_shared) IsActiveUser(ctx context.Context, ID string) (*shared.IsActiveUser, error) {
	ret := _m.Called(ctx, ID)

	var r0 *shared.IsActiveUser
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*shared.IsActiveUser, error)); ok {
		return rf(ctx, ID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *shared.IsActiveUser); ok {
		r0 = rf(ctx, ID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*shared.IsActiveUser)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Register provides a mock function with given fields: ctx, pld
func (_m *AuthRepository_internal_shared) Register(ctx context.Context, pld *shared.Register) (*shared.RegisterUserResponse, error) {
	ret := _m.Called(ctx, pld)

	var r0 *shared.RegisterUserResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *shared.Register) (*shared.RegisterUserResponse, error)); ok {
		return rf(ctx, pld)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *shared.Register) *shared.RegisterUserResponse); ok {
		r0 = rf(ctx, pld)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*shared.RegisterUserResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *shared.Register) error); ok {
		r1 = rf(ctx, pld)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewAuthRepository_internal_shared interface {
	mock.TestingT
	Cleanup(func())
}

// NewAuthRepository_internal_shared creates a new instance of AuthRepository_internal_shared. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAuthRepository_internal_shared(t mockConstructorTestingTNewAuthRepository_internal_shared) *AuthRepository_internal_shared {
	mock := &AuthRepository_internal_shared{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}