package domain_test

import (
	"fmt"
	"testing"

	"github.com/christian-gama/go-booking-api/internal/guest/domain"
	"github.com/stretchr/testify/suite"
)

type GuestTestSuite struct {
	suite.Suite

	guestId uint32
	credits float32
	roomIds []uint8
}

func (s *GuestTestSuite) SetupTest() {
	s.guestId = 1
	s.credits = 0.0
	s.roomIds = []uint8{1, 2, 3}
}

func (s *GuestTestSuite) TestNewGuest() {
	type args struct {
		id      uint32
		credits float32
		roomIds []uint8
	}

	tests := []struct {
		name string
		args args
		err  error
	}{
		{
			name: "should create a new guest",
			args: args{
				id:      s.guestId,
				credits: s.credits,
				roomIds: s.roomIds,
			},
			err: nil,
		},
		{
			name: "should return an error when guest id is zero",
			args: args{
				id:      0,
				credits: s.credits,
				roomIds: s.roomIds,
			},
			err: fmt.Errorf("guest id must be greater than zero"),
		},
		{
			name: "should return an error when guest credit is negative",
			args: args{
				id:      s.guestId,
				credits: -1,
				roomIds: s.roomIds,
			},
			err: fmt.Errorf("guest credit cannot be negative"),
		},
		{
			name: "should return an error when guest room id length is greater than 12",
			args: args{
				id:      s.guestId,
				credits: s.credits,
				roomIds: make([]uint8, 13),
			},
			err: fmt.Errorf("guest cannot have more than %d rooms reserved at the same time", domain.MaxRooms),
		},
	}

	for _, tt := range tests {
		_, err := domain.NewGuest(tt.args.id, tt.args.credits, tt.args.roomIds)
		if tt.err != nil {
			s.EqualError(err, tt.err.Error(), tt.name)
		} else {
			s.Nil(err, tt.name)
		}
	}
}

func TestGuestTestSuite(t *testing.T) {
	suite.Run(t, new(GuestTestSuite))
}
