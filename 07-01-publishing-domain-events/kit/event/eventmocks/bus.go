// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package eventmocks

import (
	context "context"

	event "github.com/diegodhdev/hexagonal-go-api/07-01-publishing-domain-events/kit/event"
	mock "github.com/stretchr/testify/mock"
)

// Bus is an autogenerated mock type for the Bus type
type Bus struct {
	mock.Mock
}

// Publish provides a mock function with given fields: _a0, _a1
func (_m *Bus) Publish(_a0 context.Context, _a1 []event.Event) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []event.Event) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
