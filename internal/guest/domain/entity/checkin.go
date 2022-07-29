package entity

import (
	"fmt"
	"time"
)

const (
	// WaitTimeToCheckin is the time that a guest must checkin in advance.
	WaitTimeToCheckin = 1 * time.Hour

	// WaitTimeToCheckout is the time that a guest must wait before checking out after checkin.
	WaitTimeToCheckout = 3 * time.Hour
)

// Checkin represents a reservation of a room for a guest.
type Checkin struct {
	id           uint32
	roomId       uint32
	guest        *Guest
	checkinDate  time.Time
	checkoutDate time.Time
}

// Id returns the checkin id.
func (c *Checkin) Id() uint32 {
	return c.id
}

// RoomId returns the room id that the guest is checking in to.
func (c *Checkin) RoomId() uint32 {
	return c.roomId
}

// Guest returns the guest that is checking in.
func (c *Checkin) Guest() *Guest {
	return c.guest
}

// CheckinDate returns the checkin date.
func (c *Checkin) CheckinDate() time.Time {
	return c.checkinDate
}

// CheckoutDate returns the checkout date.
func (c *Checkin) CheckoutDate() time.Time {
	return c.checkoutDate
}

// NewCheckin creates a new checkin. It will return an error if does not pass the validation.
func NewCheckin(
	id uint32,
	guest *Guest,
	roomId uint32,
	checkinDate time.Time,
	checkoutDate time.Time,
) (*Checkin, error) {
	if id == 0 {
		return nil, fmt.Errorf("checkin id must be greater than zero")
	}

	if roomId == 0 {
		return nil, fmt.Errorf("room id must be greater than zero")
	}

	if guest == nil {
		return nil, fmt.Errorf("guest must not be nil")
	}

	if checkoutDate.Before(checkinDate) {
		return nil, fmt.Errorf("checkin cannot be made after checkout")
	}

	if time.Until(checkinDate) < WaitTimeToCheckin {
		return nil, fmt.Errorf("checkin must be made at least %.0f hour from now", WaitTimeToCheckin.Hours())
	}

	if time.Until(checkoutDate) < WaitTimeToCheckout {
		return nil, fmt.Errorf("checkout must be made at least %.0f hour after checkin", WaitTimeToCheckout.Hours())
	}

	return &Checkin{
		id:           id,
		roomId:       roomId,
		guest:        guest,
		checkinDate:  checkinDate,
		checkoutDate: checkoutDate,
	}, nil
}