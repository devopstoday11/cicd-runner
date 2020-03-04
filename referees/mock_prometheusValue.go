// Code generated by mockery v1.0.0. DO NOT EDIT.

// This comment works around https://github.com/vektra/mockery/issues/155

package referees

import (
	model "github.com/prometheus/common/model"
	mock "github.com/stretchr/testify/mock"
)

// mockPrometheusValue is an autogenerated mock type for the prometheusValue type
type mockPrometheusValue struct {
	mock.Mock
}

// String provides a mock function with given fields:
func (_m *mockPrometheusValue) String() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Type provides a mock function with given fields:
func (_m *mockPrometheusValue) Type() model.ValueType {
	ret := _m.Called()

	var r0 model.ValueType
	if rf, ok := ret.Get(0).(func() model.ValueType); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(model.ValueType)
	}

	return r0
}