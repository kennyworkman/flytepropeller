// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import (
	context "context"

	admin "github.com/flyteorg/flyteidl/gen/pb-go/flyteidl/admin"

	core "github.com/flyteorg/flyteidl/gen/pb-go/flyteidl/core"

	mock "github.com/stretchr/testify/mock"
)

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

type Client_RecoverNodeExecution struct {
	*mock.Call
}

func (_m Client_RecoverNodeExecution) Return(_a0 *admin.NodeExecution, _a1 error) *Client_RecoverNodeExecution {
	return &Client_RecoverNodeExecution{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *Client) OnRecoverNodeExecution(ctx context.Context, execID *core.WorkflowExecutionIdentifier, id *core.NodeExecutionIdentifier) *Client_RecoverNodeExecution {
	c := _m.On("RecoverNodeExecution", ctx, execID, id)
	return &Client_RecoverNodeExecution{Call: c}
}

func (_m *Client) OnRecoverNodeExecutionMatch(matchers ...interface{}) *Client_RecoverNodeExecution {
	c := _m.On("RecoverNodeExecution", matchers...)
	return &Client_RecoverNodeExecution{Call: c}
}

// RecoverNodeExecution provides a mock function with given fields: ctx, execID, id
func (_m *Client) RecoverNodeExecution(ctx context.Context, execID *core.WorkflowExecutionIdentifier, id *core.NodeExecutionIdentifier) (*admin.NodeExecution, error) {
	ret := _m.Called(ctx, execID, id)

	var r0 *admin.NodeExecution
	if rf, ok := ret.Get(0).(func(context.Context, *core.WorkflowExecutionIdentifier, *core.NodeExecutionIdentifier) *admin.NodeExecution); ok {
		r0 = rf(ctx, execID, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.NodeExecution)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *core.WorkflowExecutionIdentifier, *core.NodeExecutionIdentifier) error); ok {
		r1 = rf(ctx, execID, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type Client_RecoverNodeExecutionData struct {
	*mock.Call
}

func (_m Client_RecoverNodeExecutionData) Return(_a0 *admin.NodeExecutionGetDataResponse, _a1 error) *Client_RecoverNodeExecutionData {
	return &Client_RecoverNodeExecutionData{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *Client) OnRecoverNodeExecutionData(ctx context.Context, execID *core.WorkflowExecutionIdentifier, id *core.NodeExecutionIdentifier) *Client_RecoverNodeExecutionData {
	c := _m.On("RecoverNodeExecutionData", ctx, execID, id)
	return &Client_RecoverNodeExecutionData{Call: c}
}

func (_m *Client) OnRecoverNodeExecutionDataMatch(matchers ...interface{}) *Client_RecoverNodeExecutionData {
	c := _m.On("RecoverNodeExecutionData", matchers...)
	return &Client_RecoverNodeExecutionData{Call: c}
}

// RecoverNodeExecutionData provides a mock function with given fields: ctx, execID, id
func (_m *Client) RecoverNodeExecutionData(ctx context.Context, execID *core.WorkflowExecutionIdentifier, id *core.NodeExecutionIdentifier) (*admin.NodeExecutionGetDataResponse, error) {
	ret := _m.Called(ctx, execID, id)

	var r0 *admin.NodeExecutionGetDataResponse
	if rf, ok := ret.Get(0).(func(context.Context, *core.WorkflowExecutionIdentifier, *core.NodeExecutionIdentifier) *admin.NodeExecutionGetDataResponse); ok {
		r0 = rf(ctx, execID, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.NodeExecutionGetDataResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *core.WorkflowExecutionIdentifier, *core.NodeExecutionIdentifier) error); ok {
		r1 = rf(ctx, execID, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
