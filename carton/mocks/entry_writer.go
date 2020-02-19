// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// EntryWriter is an autogenerated mock type for the EntryWriter type
type EntryWriter struct {
	mock.Mock
}

// Write provides a mock function with given fields: source, destination
func (_m *EntryWriter) Write(source string, destination string) error {
	ret := _m.Called(source, destination)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(source, destination)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
