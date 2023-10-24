// Code generated by mockery v2.16.0. DO NOT EDIT.

package datacoord

import (
	context "context"

	datapb "github.com/milvus-io/milvus/internal/proto/datapb"
	mock "github.com/stretchr/testify/mock"
)

// NMockHandler is an autogenerated mock type for the Handler type
type NMockHandler struct {
	mock.Mock
}

type NMockHandler_Expecter struct {
	mock *mock.Mock
}

func (_m *NMockHandler) EXPECT() *NMockHandler_Expecter {
	return &NMockHandler_Expecter{mock: &_m.Mock}
}

// CheckShouldDropChannel provides a mock function with given fields: ch
func (_m *NMockHandler) CheckShouldDropChannel(ch string) bool {
	ret := _m.Called(ch)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(ch)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// NMockHandler_CheckShouldDropChannel_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CheckShouldDropChannel'
type NMockHandler_CheckShouldDropChannel_Call struct {
	*mock.Call
}

// CheckShouldDropChannel is a helper method to define mock.On call
//   - ch string
func (_e *NMockHandler_Expecter) CheckShouldDropChannel(ch interface{}) *NMockHandler_CheckShouldDropChannel_Call {
	return &NMockHandler_CheckShouldDropChannel_Call{Call: _e.mock.On("CheckShouldDropChannel", ch)}
}

func (_c *NMockHandler_CheckShouldDropChannel_Call) Run(run func(ch string)) *NMockHandler_CheckShouldDropChannel_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *NMockHandler_CheckShouldDropChannel_Call) Return(_a0 bool) *NMockHandler_CheckShouldDropChannel_Call {
	_c.Call.Return(_a0)
	return _c
}

// FinishDropChannel provides a mock function with given fields: ch
func (_m *NMockHandler) FinishDropChannel(ch string) error {
	ret := _m.Called(ch)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(ch)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NMockHandler_FinishDropChannel_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FinishDropChannel'
type NMockHandler_FinishDropChannel_Call struct {
	*mock.Call
}

// FinishDropChannel is a helper method to define mock.On call
//   - ch string
func (_e *NMockHandler_Expecter) FinishDropChannel(ch interface{}) *NMockHandler_FinishDropChannel_Call {
	return &NMockHandler_FinishDropChannel_Call{Call: _e.mock.On("FinishDropChannel", ch)}
}

func (_c *NMockHandler_FinishDropChannel_Call) Run(run func(ch string)) *NMockHandler_FinishDropChannel_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *NMockHandler_FinishDropChannel_Call) Return(_a0 error) *NMockHandler_FinishDropChannel_Call {
	_c.Call.Return(_a0)
	return _c
}

// GetCollection provides a mock function with given fields: ctx, collectionID
func (_m *NMockHandler) GetCollection(ctx context.Context, collectionID int64) (*collectionInfo, error) {
	ret := _m.Called(ctx, collectionID)

	var r0 *collectionInfo
	if rf, ok := ret.Get(0).(func(context.Context, int64) *collectionInfo); ok {
		r0 = rf(ctx, collectionID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*collectionInfo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, collectionID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NMockHandler_GetCollection_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetCollection'
type NMockHandler_GetCollection_Call struct {
	*mock.Call
}

// GetCollection is a helper method to define mock.On call
//   - ctx context.Context
//   - collectionID int64
func (_e *NMockHandler_Expecter) GetCollection(ctx interface{}, collectionID interface{}) *NMockHandler_GetCollection_Call {
	return &NMockHandler_GetCollection_Call{Call: _e.mock.On("GetCollection", ctx, collectionID)}
}

func (_c *NMockHandler_GetCollection_Call) Run(run func(ctx context.Context, collectionID int64)) *NMockHandler_GetCollection_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64))
	})
	return _c
}

func (_c *NMockHandler_GetCollection_Call) Return(_a0 *collectionInfo, _a1 error) *NMockHandler_GetCollection_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetDataVChanPositions provides a mock function with given fields: channel, partitionID
func (_m *NMockHandler) GetDataVChanPositions(ch *channel, partitionID int64) *datapb.VchannelInfo {
	ret := _m.Called(ch, partitionID)

	var r0 *datapb.VchannelInfo
	if rf, ok := ret.Get(0).(func(*channel, int64) *datapb.VchannelInfo); ok {
		r0 = rf(ch, partitionID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*datapb.VchannelInfo)
		}
	}

	return r0
}

// NMockHandler_GetDataVChanPositions_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetDataVChanPositions'
type NMockHandler_GetDataVChanPositions_Call struct {
	*mock.Call
}

// GetDataVChanPositions is a helper method to define mock.On call
//   - channel *channel
//   - partitionID int64
func (_e *NMockHandler_Expecter) GetDataVChanPositions(channel interface{}, partitionID interface{}) *NMockHandler_GetDataVChanPositions_Call {
	return &NMockHandler_GetDataVChanPositions_Call{Call: _e.mock.On("GetDataVChanPositions", channel, partitionID)}
}

func (_c *NMockHandler_GetDataVChanPositions_Call) Run(run func(channel *channel, partitionID int64)) *NMockHandler_GetDataVChanPositions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*channel), args[1].(int64))
	})
	return _c
}

func (_c *NMockHandler_GetDataVChanPositions_Call) Return(_a0 *datapb.VchannelInfo) *NMockHandler_GetDataVChanPositions_Call {
	_c.Call.Return(_a0)
	return _c
}

// GetQueryVChanPositions provides a mock function with given fields: channel, partitionID
func (_m *NMockHandler) GetQueryVChanPositions(ch *channel, partitionID ...int64) (*datapb.VchannelInfo, error) {
	_va := make([]interface{}, len(partitionID))
	for _i := range partitionID {
		_va[_i] = partitionID[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ch)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *datapb.VchannelInfo
	if rf, ok := ret.Get(0).(func(*channel, ...int64) *datapb.VchannelInfo); ok {
		r0 = rf(ch, partitionID...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*datapb.VchannelInfo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*channel, ...int64) error); ok {
		r1 = rf(ch, partitionID...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NMockHandler_GetQueryVChanPositions_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetQueryVChanPositions'
type NMockHandler_GetQueryVChanPositions_Call struct {
	*mock.Call
}

// GetQueryVChanPositions is a helper method to define mock.On call
//   - channel *channel
//   - partitionID ...int64
func (_e *NMockHandler_Expecter) GetQueryVChanPositions(channel interface{}, partitionID ...interface{}) *NMockHandler_GetQueryVChanPositions_Call {
	return &NMockHandler_GetQueryVChanPositions_Call{Call: _e.mock.On("GetQueryVChanPositions",
		append([]interface{}{channel}, partitionID...)...)}
}

func (_c *NMockHandler_GetQueryVChanPositions_Call) Run(run func(channel *channel, partitionID ...int64)) *NMockHandler_GetQueryVChanPositions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]int64, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(int64)
			}
		}
		run(args[0].(*channel), variadicArgs...)
	})
	return _c
}

func (_c *NMockHandler_GetQueryVChanPositions_Call) Return(_a0 *datapb.VchannelInfo, _a1 error) *NMockHandler_GetQueryVChanPositions_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

type mockConstructorTestingTNewNMockHandler interface {
	mock.TestingT
	Cleanup(func())
}

// NewNMockHandler creates a new instance of NMockHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewNMockHandler(t mockConstructorTestingTNewNMockHandler) *NMockHandler {
	mock := &NMockHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
