// Code generated by mockery v2.46.3. DO NOT EDIT.

package mocks

import (
	p "mockery-issue/p"

	mock "github.com/stretchr/testify/mock"
)

// I is an autogenerated mock type for the I type
type I struct {
	mock.Mock
}

// F provides a mock function with given fields: _a0
func (_m *I) F(_a0 p.T) {
	_m.Called(_a0)
}

// NewI creates a new instance of I. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewI(t interface {
	mock.TestingT
	Cleanup(func())
}) *I {
	mock := &I{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
