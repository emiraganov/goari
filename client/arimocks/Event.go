package arimocks

import ari "github.com/CyCoreSystems/ari"
import mock "github.com/stretchr/testify/mock"

// Event is an autogenerated mock type for the Event type
type Event struct {
	mock.Mock
}

// GetApplication provides a mock function with given fields:
func (_m *Event) GetApplication() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetDialog provides a mock function with given fields:
func (_m *Event) GetDialog() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetNode provides a mock function with given fields:
func (_m *Event) GetNode() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetType provides a mock function with given fields:
func (_m *Event) GetType() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Key provides a mock function with given fields: kind, id
func (_m *Event) Key(kind string, id string) *ari.Key {
	ret := _m.Called(kind, id)

	var r0 *ari.Key
	if rf, ok := ret.Get(0).(func(string, string) *ari.Key); ok {
		r0 = rf(kind, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ari.Key)
		}
	}

	return r0
}

// Keys provides a mock function with given fields:
func (_m *Event) Keys() ari.Keys {
	ret := _m.Called()

	var r0 ari.Keys
	if rf, ok := ret.Get(0).(func() ari.Keys); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ari.Keys)
		}
	}

	return r0
}

// SetDialog provides a mock function with given fields: _a0
func (_m *Event) SetDialog(_a0 string) {
	_m.Called(_a0)
}
