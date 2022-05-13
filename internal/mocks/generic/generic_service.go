// Code generated by mockery v2.12.2. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	testing "testing"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// GenericCall provides a mock function with given fields: ctx, method, request
func (_m *Service) GenericCall(ctx context.Context, method string, request interface{}) (interface{}, error) {
	ret := _m.Called(ctx, method, request)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(context.Context, string, interface{}) interface{}); ok {
		r0 = rf(ctx, method, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, interface{}) error); ok {
		r1 = rf(ctx, method, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewService creates a new instance of Service. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewService(t testing.TB) *Service {
	mock := &Service{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
