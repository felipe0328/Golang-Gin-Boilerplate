// Code generated by mockery v2.18.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	usersmodels "gin-microservice/models/usersModels"
)

// IUsersController is an autogenerated mock type for the IUsersController type
type IUsersController struct {
	mock.Mock
}

// CreateUser provides a mock function with given fields: _a0
func (_m *IUsersController) CreateUser(_a0 usersmodels.User) (usersmodels.UserObject, error) {
	ret := _m.Called(_a0)

	var r0 usersmodels.UserObject
	if rf, ok := ret.Get(0).(func(usersmodels.User) usersmodels.UserObject); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(usersmodels.UserObject)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(usersmodels.User) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// VerifyLogin provides a mock function with given fields: _a0
func (_m *IUsersController) VerifyLogin(_a0 usersmodels.UserLogin) (string, error) {
	ret := _m.Called(_a0)

	var r0 string
	if rf, ok := ret.Get(0).(func(usersmodels.UserLogin) string); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(usersmodels.UserLogin) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewIUsersController interface {
	mock.TestingT
	Cleanup(func())
}

// NewIUsersController creates a new instance of IUsersController. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIUsersController(t mockConstructorTestingTNewIUsersController) *IUsersController {
	mock := &IUsersController{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
