// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Credential is an autogenerated mock type for the Credential type
type Credential struct {
	mock.Mock
}

type Credential_Expecter struct {
	mock *mock.Mock
}

func (_m *Credential) EXPECT() *Credential_Expecter {
	return &Credential_Expecter{mock: &_m.Mock}
}

// GetAccessKeyId provides a mock function with given fields:
func (_m *Credential) GetAccessKeyId() (*string, error) {
	ret := _m.Called()

	var r0 *string
	if rf, ok := ret.Get(0).(func() *string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Credential_GetAccessKeyId_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAccessKeyId'
type Credential_GetAccessKeyId_Call struct {
	*mock.Call
}

// GetAccessKeyId is a helper method to define mock.On call
func (_e *Credential_Expecter) GetAccessKeyId() *Credential_GetAccessKeyId_Call {
	return &Credential_GetAccessKeyId_Call{Call: _e.mock.On("GetAccessKeyId")}
}

func (_c *Credential_GetAccessKeyId_Call) Run(run func()) *Credential_GetAccessKeyId_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Credential_GetAccessKeyId_Call) Return(_a0 *string, _a1 error) *Credential_GetAccessKeyId_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetAccessKeySecret provides a mock function with given fields:
func (_m *Credential) GetAccessKeySecret() (*string, error) {
	ret := _m.Called()

	var r0 *string
	if rf, ok := ret.Get(0).(func() *string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Credential_GetAccessKeySecret_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAccessKeySecret'
type Credential_GetAccessKeySecret_Call struct {
	*mock.Call
}

// GetAccessKeySecret is a helper method to define mock.On call
func (_e *Credential_Expecter) GetAccessKeySecret() *Credential_GetAccessKeySecret_Call {
	return &Credential_GetAccessKeySecret_Call{Call: _e.mock.On("GetAccessKeySecret")}
}

func (_c *Credential_GetAccessKeySecret_Call) Run(run func()) *Credential_GetAccessKeySecret_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Credential_GetAccessKeySecret_Call) Return(_a0 *string, _a1 error) *Credential_GetAccessKeySecret_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetBearerToken provides a mock function with given fields:
func (_m *Credential) GetBearerToken() *string {
	ret := _m.Called()

	var r0 *string
	if rf, ok := ret.Get(0).(func() *string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*string)
		}
	}

	return r0
}

// Credential_GetBearerToken_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetBearerToken'
type Credential_GetBearerToken_Call struct {
	*mock.Call
}

// GetBearerToken is a helper method to define mock.On call
func (_e *Credential_Expecter) GetBearerToken() *Credential_GetBearerToken_Call {
	return &Credential_GetBearerToken_Call{Call: _e.mock.On("GetBearerToken")}
}

func (_c *Credential_GetBearerToken_Call) Run(run func()) *Credential_GetBearerToken_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Credential_GetBearerToken_Call) Return(_a0 *string) *Credential_GetBearerToken_Call {
	_c.Call.Return(_a0)
	return _c
}

// GetSecurityToken provides a mock function with given fields:
func (_m *Credential) GetSecurityToken() (*string, error) {
	ret := _m.Called()

	var r0 *string
	if rf, ok := ret.Get(0).(func() *string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Credential_GetSecurityToken_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSecurityToken'
type Credential_GetSecurityToken_Call struct {
	*mock.Call
}

// GetSecurityToken is a helper method to define mock.On call
func (_e *Credential_Expecter) GetSecurityToken() *Credential_GetSecurityToken_Call {
	return &Credential_GetSecurityToken_Call{Call: _e.mock.On("GetSecurityToken")}
}

func (_c *Credential_GetSecurityToken_Call) Run(run func()) *Credential_GetSecurityToken_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Credential_GetSecurityToken_Call) Return(_a0 *string, _a1 error) *Credential_GetSecurityToken_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetType provides a mock function with given fields:
func (_m *Credential) GetType() *string {
	ret := _m.Called()

	var r0 *string
	if rf, ok := ret.Get(0).(func() *string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*string)
		}
	}

	return r0
}

// Credential_GetType_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetType'
type Credential_GetType_Call struct {
	*mock.Call
}

// GetType is a helper method to define mock.On call
func (_e *Credential_Expecter) GetType() *Credential_GetType_Call {
	return &Credential_GetType_Call{Call: _e.mock.On("GetType")}
}

func (_c *Credential_GetType_Call) Run(run func()) *Credential_GetType_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Credential_GetType_Call) Return(_a0 *string) *Credential_GetType_Call {
	_c.Call.Return(_a0)
	return _c
}

type mockConstructorTestingTNewCredential interface {
	mock.TestingT
	Cleanup(func())
}

// NewCredential creates a new instance of Credential. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCredential(t mockConstructorTestingTNewCredential) *Credential {
	mock := &Credential{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
