// Code generated by github.com/efritz/go-mockgen 0.1.0; DO NOT EDIT.
// This file was generated by robots at
// 2019-06-04T17:47:47-05:00
// using the command
// $ go-mockgen -f github.com/go-nacelle/service

package mocks

import service "github.com/go-nacelle/service"

// MockPostInject is a mock impelementation of the PostInject interface
// (from the package github.com/go-nacelle/service) used for unit testing.
type MockPostInject struct {
	// PostInjectFunc is an instance of a mock function object controlling
	// the behavior of the method PostInject.
	PostInjectFunc *PostInjectPostInjectFunc
}

// NewMockPostInject creates a new mock of the PostInject interface. All
// methods return zero values for all results, unless overwritten.
func NewMockPostInject() *MockPostInject {
	return &MockPostInject{
		PostInjectFunc: &PostInjectPostInjectFunc{
			defaultHook: func() error {
				return nil
			},
		},
	}
}

// NewMockPostInjectFrom creates a new mock of the MockPostInject interface.
// All methods delegate to the given implementation, unless overwritten.
func NewMockPostInjectFrom(i service.PostInject) *MockPostInject {
	return &MockPostInject{
		PostInjectFunc: &PostInjectPostInjectFunc{
			defaultHook: i.PostInject,
		},
	}
}

// PostInjectPostInjectFunc describes the behavior when the PostInject
// method of the parent MockPostInject instance is invoked.
type PostInjectPostInjectFunc struct {
	defaultHook func() error
	hooks       []func() error
	history     []PostInjectPostInjectFuncCall
}

// PostInject delegates to the next hook function in the queue and stores
// the parameter and result values of this invocation.
func (m *MockPostInject) PostInject() error {
	r0 := m.PostInjectFunc.nextHook()()
	m.PostInjectFunc.history = append(m.PostInjectFunc.history, PostInjectPostInjectFuncCall{r0})
	return r0
}

// SetDefaultHook sets function that is called when the PostInject method of
// the parent MockPostInject instance is invoked and the hook queue is
// empty.
func (f *PostInjectPostInjectFunc) SetDefaultHook(hook func() error) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// PostInject method of the parent MockPostInject instance inovkes the hook
// at the front of the queue and discards it. After the queue is empty, the
// default hook function is invoked for any future action.
func (f *PostInjectPostInjectFunc) PushHook(hook func() error) {
	f.hooks = append(f.hooks, hook)
}

// SetDefaultReturn calls SetDefaultDefaultHook with a function that returns
// the given values.
func (f *PostInjectPostInjectFunc) SetDefaultReturn(r0 error) {
	f.SetDefaultHook(func() error {
		return r0
	})
}

// PushReturn calls PushDefaultHook with a function that returns the given
// values.
func (f *PostInjectPostInjectFunc) PushReturn(r0 error) {
	f.PushHook(func() error {
		return r0
	})
}

func (f *PostInjectPostInjectFunc) nextHook() func() error {
	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

// History returns a sequence of PostInjectPostInjectFuncCall objects
// describing the invocations of this function.
func (f *PostInjectPostInjectFunc) History() []PostInjectPostInjectFuncCall {
	return f.history
}

// PostInjectPostInjectFuncCall is an object that describes an invocation of
// method PostInject on an instance of MockPostInject.
type PostInjectPostInjectFuncCall struct {
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c PostInjectPostInjectFuncCall) Args() []interface{} {
	return []interface{}{}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c PostInjectPostInjectFuncCall) Results() []interface{} {
	return []interface{}{c.Result0}
}
