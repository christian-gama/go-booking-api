// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	entity "github.com/christian-gama/go-booking-api/internal/room/domain/entity"
	mock "github.com/stretchr/testify/mock"
)

// Room is an autogenerated mock type for the Room type
type Room struct {
	mock.Mock
}

// SaveRoom provides a mock function with given fields: room
func (_m *Room) SaveRoom(room *entity.Room) (*entity.Room, error) {
	ret := _m.Called(room)

	var r0 *entity.Room
	if rf, ok := ret.Get(0).(func(*entity.Room) *entity.Room); ok {
		r0 = rf(room)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Room)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*entity.Room) error); ok {
		r1 = rf(room)
	} else {
		r1 = ret.Error(1)
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