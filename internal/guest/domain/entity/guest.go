package entity

import "fmt"

const (
	// The amount of rooms that a guest can have reserved at the same time, but not necessarily in the same date.
	MaxRooms = 12
)

// Guest represents a guest in the hotel. He is able to make checkins and checkouts of rooms and are allowed to
// have a limited amount of credits, which can be used to pay for rooms. Credits are earned when a guest asks
// for a refund.
type Guest struct {
	id      uint32
	credits float32
	roomIds []uint8
}

// Id returns the guest id.
func (g *Guest) Id() uint32 {
	return g.id
}

// Credits returns the guest credits. It will never return a negative value.
func (g *Guest) Credits() float32 {
	return g.credits
}

// Checkin adds a room id to the guest's room ids. It will return an error if does not pass the validation.
func (g *Guest) RoomIds() []uint8 {
	return g.roomIds
}

// NewGuest creates a new guest. It will return an error if does not pass the validation.
func NewGuest(
	id uint32,
	credits float32,
	roomIds []uint8,
) (*Guest, error) {
	if id == 0 {
		return nil, fmt.Errorf("guest id must be greater than zero")
	}

	if credits < 0 {
		return nil, fmt.Errorf("guest credit cannot be negative")
	}

	if len(roomIds) > MaxRooms {
		return nil, fmt.Errorf("guest cannot have more than %d rooms reserved at the same time", MaxRooms)
	}

	return &Guest{
		id:      id,
		credits: credits,
		roomIds: roomIds,
	}, nil
}