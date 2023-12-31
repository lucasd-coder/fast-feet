// Code generated by mockery v2.20.2. DO NOT EDIT.

package mocks

import (
	context "context"

	model "github.com/lucasd-coder/fast-feet/user-manger-service/internal/domain/user"
	mock "github.com/stretchr/testify/mock"
)

// UserRepository_internal_domain_user is an autogenerated mock type for the UserRepository type
type UserRepository_internal_domain_user struct {
	mock.Mock
}

// FindByCpf provides a mock function with given fields: ctx, cpf
func (_m *UserRepository_internal_domain_user) FindByCpf(ctx context.Context, cpf string) (*model.User, error) {
	ret := _m.Called(ctx, cpf)

	var r0 *model.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*model.User, error)); ok {
		return rf(ctx, cpf)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *model.User); ok {
		r0 = rf(ctx, cpf)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, cpf)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByEmail provides a mock function with given fields: ctx, email
func (_m *UserRepository_internal_domain_user) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	ret := _m.Called(ctx, email)

	var r0 *model.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*model.User, error)); ok {
		return rf(ctx, email)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *model.User); ok {
		r0 = rf(ctx, email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByUserID provides a mock function with given fields: ctx, userID
func (_m *UserRepository_internal_domain_user) FindByUserID(ctx context.Context, userID string) (*model.User, error) {
	ret := _m.Called(ctx, userID)

	var r0 *model.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*model.User, error)); ok {
		return rf(ctx, userID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *model.User); ok {
		r0 = rf(ctx, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: ctx, user
func (_m *UserRepository_internal_domain_user) Save(ctx context.Context, user *model.User) (*model.User, error) {
	ret := _m.Called(ctx, user)

	var r0 *model.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.User) (*model.User, error)); ok {
		return rf(ctx, user)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *model.User) *model.User); ok {
		r0 = rf(ctx, user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *model.User) error); ok {
		r1 = rf(ctx, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewUserRepository_internal_domain_user interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserRepository_internal_domain_user creates a new instance of UserRepository_internal_domain_user. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserRepository_internal_domain_user(t mockConstructorTestingTNewUserRepository_internal_domain_user) *UserRepository_internal_domain_user {
	mock := &UserRepository_internal_domain_user{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
