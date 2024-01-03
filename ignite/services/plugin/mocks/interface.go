// Code generated by mockery v2.36.1. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	plugin "github.com/ignite/cli/v28/ignite/services/plugin"

	v1 "github.com/ignite/cli/v28/ignite/services/plugin/grpc/v1"
)

// PluginInterface is an autogenerated mock type for the Interface type
type PluginInterface struct {
	mock.Mock
}

type PluginInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *PluginInterface) EXPECT() *PluginInterface_Expecter {
	return &PluginInterface_Expecter{mock: &_m.Mock}
}

// Execute provides a mock function with given fields: _a0, _a1, _a2
func (_m *PluginInterface) Execute(_a0 context.Context, _a1 *v1.ExecutedCommand, _a2 plugin.ClientAPI) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1.ExecutedCommand, plugin.ClientAPI) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PluginInterface_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type PluginInterface_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *v1.ExecutedCommand
//   - _a2 plugin.ClientAPI
func (_e *PluginInterface_Expecter) Execute(_a0 interface{}, _a1 interface{}, _a2 interface{}) *PluginInterface_Execute_Call {
	return &PluginInterface_Execute_Call{Call: _e.mock.On("Execute", _a0, _a1, _a2)}
}

func (_c *PluginInterface_Execute_Call) Run(run func(_a0 context.Context, _a1 *v1.ExecutedCommand, _a2 plugin.ClientAPI)) *PluginInterface_Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*v1.ExecutedCommand), args[2].(plugin.ClientAPI))
	})
	return _c
}

func (_c *PluginInterface_Execute_Call) Return(_a0 error) *PluginInterface_Execute_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *PluginInterface_Execute_Call) RunAndReturn(run func(context.Context, *v1.ExecutedCommand, plugin.ClientAPI) error) *PluginInterface_Execute_Call {
	_c.Call.Return(run)
	return _c
}

// ExecuteHookCleanUp provides a mock function with given fields: _a0, _a1, _a2
func (_m *PluginInterface) ExecuteHookCleanUp(_a0 context.Context, _a1 *v1.ExecutedHook, _a2 plugin.ClientAPI) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1.ExecutedHook, plugin.ClientAPI) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PluginInterface_ExecuteHookCleanUp_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ExecuteHookCleanUp'
type PluginInterface_ExecuteHookCleanUp_Call struct {
	*mock.Call
}

// ExecuteHookCleanUp is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *v1.ExecutedHook
//   - _a2 plugin.ClientAPI
func (_e *PluginInterface_Expecter) ExecuteHookCleanUp(_a0 interface{}, _a1 interface{}, _a2 interface{}) *PluginInterface_ExecuteHookCleanUp_Call {
	return &PluginInterface_ExecuteHookCleanUp_Call{Call: _e.mock.On("ExecuteHookCleanUp", _a0, _a1, _a2)}
}

func (_c *PluginInterface_ExecuteHookCleanUp_Call) Run(run func(_a0 context.Context, _a1 *v1.ExecutedHook, _a2 plugin.ClientAPI)) *PluginInterface_ExecuteHookCleanUp_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*v1.ExecutedHook), args[2].(plugin.ClientAPI))
	})
	return _c
}

func (_c *PluginInterface_ExecuteHookCleanUp_Call) Return(_a0 error) *PluginInterface_ExecuteHookCleanUp_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *PluginInterface_ExecuteHookCleanUp_Call) RunAndReturn(run func(context.Context, *v1.ExecutedHook, plugin.ClientAPI) error) *PluginInterface_ExecuteHookCleanUp_Call {
	_c.Call.Return(run)
	return _c
}

// ExecuteHookPost provides a mock function with given fields: _a0, _a1, _a2
func (_m *PluginInterface) ExecuteHookPost(_a0 context.Context, _a1 *v1.ExecutedHook, _a2 plugin.ClientAPI) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1.ExecutedHook, plugin.ClientAPI) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PluginInterface_ExecuteHookPost_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ExecuteHookPost'
type PluginInterface_ExecuteHookPost_Call struct {
	*mock.Call
}

// ExecuteHookPost is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *v1.ExecutedHook
//   - _a2 plugin.ClientAPI
func (_e *PluginInterface_Expecter) ExecuteHookPost(_a0 interface{}, _a1 interface{}, _a2 interface{}) *PluginInterface_ExecuteHookPost_Call {
	return &PluginInterface_ExecuteHookPost_Call{Call: _e.mock.On("ExecuteHookPost", _a0, _a1, _a2)}
}

func (_c *PluginInterface_ExecuteHookPost_Call) Run(run func(_a0 context.Context, _a1 *v1.ExecutedHook, _a2 plugin.ClientAPI)) *PluginInterface_ExecuteHookPost_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*v1.ExecutedHook), args[2].(plugin.ClientAPI))
	})
	return _c
}

func (_c *PluginInterface_ExecuteHookPost_Call) Return(_a0 error) *PluginInterface_ExecuteHookPost_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *PluginInterface_ExecuteHookPost_Call) RunAndReturn(run func(context.Context, *v1.ExecutedHook, plugin.ClientAPI) error) *PluginInterface_ExecuteHookPost_Call {
	_c.Call.Return(run)
	return _c
}

// ExecuteHookPre provides a mock function with given fields: _a0, _a1, _a2
func (_m *PluginInterface) ExecuteHookPre(_a0 context.Context, _a1 *v1.ExecutedHook, _a2 plugin.ClientAPI) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1.ExecutedHook, plugin.ClientAPI) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PluginInterface_ExecuteHookPre_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ExecuteHookPre'
type PluginInterface_ExecuteHookPre_Call struct {
	*mock.Call
}

// ExecuteHookPre is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *v1.ExecutedHook
//   - _a2 plugin.ClientAPI
func (_e *PluginInterface_Expecter) ExecuteHookPre(_a0 interface{}, _a1 interface{}, _a2 interface{}) *PluginInterface_ExecuteHookPre_Call {
	return &PluginInterface_ExecuteHookPre_Call{Call: _e.mock.On("ExecuteHookPre", _a0, _a1, _a2)}
}

func (_c *PluginInterface_ExecuteHookPre_Call) Run(run func(_a0 context.Context, _a1 *v1.ExecutedHook, _a2 plugin.ClientAPI)) *PluginInterface_ExecuteHookPre_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*v1.ExecutedHook), args[2].(plugin.ClientAPI))
	})
	return _c
}

func (_c *PluginInterface_ExecuteHookPre_Call) Return(_a0 error) *PluginInterface_ExecuteHookPre_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *PluginInterface_ExecuteHookPre_Call) RunAndReturn(run func(context.Context, *v1.ExecutedHook, plugin.ClientAPI) error) *PluginInterface_ExecuteHookPre_Call {
	_c.Call.Return(run)
	return _c
}

// Manifest provides a mock function with given fields: _a0
func (_m *PluginInterface) Manifest(_a0 context.Context) (*v1.Manifest, error) {
	ret := _m.Called(_a0)

	var r0 *v1.Manifest
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*v1.Manifest, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *v1.Manifest); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.Manifest)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PluginInterface_Manifest_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Manifest'
type PluginInterface_Manifest_Call struct {
	*mock.Call
}

// Manifest is a helper method to define mock.On call
//   - _a0 context.Context
func (_e *PluginInterface_Expecter) Manifest(_a0 interface{}) *PluginInterface_Manifest_Call {
	return &PluginInterface_Manifest_Call{Call: _e.mock.On("Manifest", _a0)}
}

func (_c *PluginInterface_Manifest_Call) Run(run func(_a0 context.Context)) *PluginInterface_Manifest_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *PluginInterface_Manifest_Call) Return(_a0 *v1.Manifest, _a1 error) *PluginInterface_Manifest_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *PluginInterface_Manifest_Call) RunAndReturn(run func(context.Context) (*v1.Manifest, error)) *PluginInterface_Manifest_Call {
	_c.Call.Return(run)
	return _c
}

// NewPluginInterface creates a new instance of PluginInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPluginInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *PluginInterface {
	mock := &PluginInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
