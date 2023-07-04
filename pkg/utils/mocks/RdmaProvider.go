// Code generated by mockery v2.26.1. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// RdmaProvider is an autogenerated mock type for the RdmaProvider type
type RdmaProvider struct {
	mock.Mock
}

// GetRdmaCharDevices provides a mock function with given fields: rdmaDeviceName
func (_m *RdmaProvider) GetRdmaCharDevices(rdmaDeviceName string) []string {
	ret := _m.Called(rdmaDeviceName)

	var r0 []string
	if rf, ok := ret.Get(0).(func(string) []string); ok {
		r0 = rf(rdmaDeviceName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// GetRdmaDevicesForAuxdev provides a mock function with given fields: deviceID
func (_m *RdmaProvider) GetRdmaDevicesForAuxdev(deviceID string) []string {
	ret := _m.Called(deviceID)

	var r0 []string
	if rf, ok := ret.Get(0).(func(string) []string); ok {
		r0 = rf(deviceID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// GetRdmaDevicesForPcidev provides a mock function with given fields: pciAddr
func (_m *RdmaProvider) GetRdmaDevicesForPcidev(pciAddr string) []string {
	ret := _m.Called(pciAddr)

	var r0 []string
	if rf, ok := ret.Get(0).(func(string) []string); ok {
		r0 = rf(pciAddr)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

type mockConstructorTestingTNewRdmaProvider interface {
	mock.TestingT
	Cleanup(func())
}

// NewRdmaProvider creates a new instance of RdmaProvider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRdmaProvider(t mockConstructorTestingTNewRdmaProvider) *RdmaProvider {
	mock := &RdmaProvider{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
