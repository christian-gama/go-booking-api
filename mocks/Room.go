// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	entity "github.com/christian-gama/go-booking-api/internal/room/domain/entity"
	exception "github.com/christian-gama/go-booking-api/internal/shared/app/exception"

	mock "github.com/stretchr/testify/mock"
)

// Room is an autogenerated mock type for the Room type
type Room struct {
	mock.Mock
}

// SaveRoom provides a mock function with given fields: room
func (_m *Room) SaveRoom(room *entity.Room) (*entity.Room, *exception.Error) {
	ret := _m.Called(room)

	var r0 *entity.Room
	if rf, ok := ret.Get(0).(func(*entity.Room) *entity.Room); ok {
		r0 = rf(room)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Room)
		}
	}

	var r1 *exception.Error
	if rf, ok := ret.Get(1).(func(*entity.Room) *exception.Error); ok {
		r1 = rf(room)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*exception.Error)
		}
	}

	return r0, r1
}

type mockConstructorTestingTNewRoom interface {
	mock.TestingT
	Cleanup(func())
}

// NewRoom creates a new instance of Room. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRoom(t mockConstructorTestingTNewRoom) *Room {
	mock := &Room{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
