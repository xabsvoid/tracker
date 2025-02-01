// Code generated by mockery v2.51.0. DO NOT EDIT.

package mock

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	model "github.com/xabsvoid/tracker/internal/app/geo/model"
)

// MockLocation is an autogenerated mock type for the Location type
type MockLocation struct {
	mock.Mock
}

type MockLocation_Expecter struct {
	mock *mock.Mock
}

func (_m *MockLocation) EXPECT() *MockLocation_Expecter {
	return &MockLocation_Expecter{mock: &_m.Mock}
}

// GetByUUID provides a mock function with given fields: ctx, uuid
func (_m *MockLocation) GetByUUID(ctx context.Context, uuid model.UUID) (model.Location, error) {
	ret := _m.Called(ctx, uuid)

	if len(ret) == 0 {
		panic("no return value specified for GetByUUID")
	}

	var r0 model.Location
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, model.UUID) (model.Location, error)); ok {
		return rf(ctx, uuid)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.UUID) model.Location); ok {
		r0 = rf(ctx, uuid)
	} else {
		r0 = ret.Get(0).(model.Location)
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.UUID) error); ok {
		r1 = rf(ctx, uuid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockLocation_GetByUUID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByUUID'
type MockLocation_GetByUUID_Call struct {
	*mock.Call
}

// GetByUUID is a helper method to define mock.On call
//   - ctx context.Context
//   - uuid model.UUID
func (_e *MockLocation_Expecter) GetByUUID(ctx interface{}, uuid interface{}) *MockLocation_GetByUUID_Call {
	return &MockLocation_GetByUUID_Call{Call: _e.mock.On("GetByUUID", ctx, uuid)}
}

func (_c *MockLocation_GetByUUID_Call) Run(run func(ctx context.Context, uuid model.UUID)) *MockLocation_GetByUUID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(model.UUID))
	})
	return _c
}

func (_c *MockLocation_GetByUUID_Call) Return(_a0 model.Location, _a1 error) *MockLocation_GetByUUID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockLocation_GetByUUID_Call) RunAndReturn(run func(context.Context, model.UUID) (model.Location, error)) *MockLocation_GetByUUID_Call {
	_c.Call.Return(run)
	return _c
}

// Set provides a mock function with given fields: ctx, location
func (_m *MockLocation) Set(ctx context.Context, location model.Location) error {
	ret := _m.Called(ctx, location)

	if len(ret) == 0 {
		panic("no return value specified for Set")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.Location) error); ok {
		r0 = rf(ctx, location)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockLocation_Set_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Set'
type MockLocation_Set_Call struct {
	*mock.Call
}

// Set is a helper method to define mock.On call
//   - ctx context.Context
//   - location model.Location
func (_e *MockLocation_Expecter) Set(ctx interface{}, location interface{}) *MockLocation_Set_Call {
	return &MockLocation_Set_Call{Call: _e.mock.On("Set", ctx, location)}
}

func (_c *MockLocation_Set_Call) Run(run func(ctx context.Context, location model.Location)) *MockLocation_Set_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(model.Location))
	})
	return _c
}

func (_c *MockLocation_Set_Call) Return(_a0 error) *MockLocation_Set_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockLocation_Set_Call) RunAndReturn(run func(context.Context, model.Location) error) *MockLocation_Set_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockLocation creates a new instance of MockLocation. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockLocation(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockLocation {
	mock := &MockLocation{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
