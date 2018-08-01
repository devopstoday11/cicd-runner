// Code generated by mockery v1.0.0. DO NOT EDIT.

// This comment works around https://github.com/vektra/mockery/issues/155

package mocks

import mock "github.com/stretchr/testify/mock"

// Logger is an autogenerated mock type for the Logger type
type Logger struct {
	mock.Mock
}

// Error provides a mock function with given fields: v
func (_m *Logger) Error(v ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, v...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(...interface{}) error); ok {
		r0 = rf(v...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Errorf provides a mock function with given fields: format, a
func (_m *Logger) Errorf(format string, a ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, a...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, ...interface{}) error); ok {
		r0 = rf(format, a...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Info provides a mock function with given fields: v
func (_m *Logger) Info(v ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, v...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(...interface{}) error); ok {
		r0 = rf(v...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Infof provides a mock function with given fields: format, a
func (_m *Logger) Infof(format string, a ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, a...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, ...interface{}) error); ok {
		r0 = rf(format, a...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Warning provides a mock function with given fields: v
func (_m *Logger) Warning(v ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, v...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(...interface{}) error); ok {
		r0 = rf(v...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Warningf provides a mock function with given fields: format, a
func (_m *Logger) Warningf(format string, a ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, a...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, ...interface{}) error); ok {
		r0 = rf(format, a...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
