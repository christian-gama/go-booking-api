package service

import (
	"github.com/christian-gama/go-booking-api/internal/room/app/dto"
	"github.com/christian-gama/go-booking-api/internal/room/app/repo"
	"github.com/christian-gama/go-booking-api/internal/room/domain/entity"
	"github.com/christian-gama/go-booking-api/internal/shared/app/uuid"
	"github.com/christian-gama/go-booking-api/internal/shared/domain/error"
)

// CreatRoomService is the interface that defines the CreateRoomService.
type CreateRoomService interface {
	Handle(input *dto.CreateRoom) (*entity.Room, []*error.Error)
}

// createRoom is a concrete implementation of the CreateRoomService.
type createRoom struct {
	repo repo.Room
	uuid uuid.UUID
}

// Handle receives an input and creates a room. It will return an error if something
// goes wrong with room creation or if the room repo return an error.
func (c *createRoom) Handle(input *dto.CreateRoom) (*entity.Room, []*error.Error) {
	uuid := c.uuid.Generate()
	room, err := entity.NewRoom(uuid, input.Name, input.Description, input.BedCount, input.Price, false)
	if err != nil {
		return nil, err
	}

	room, err = c.repo.SaveRoom(room)
	if err != nil {
		return nil, err
	}

	return room, nil
}

// NewCreateRoom creates a new CreateRoomService.
func NewCreateRoom(repo repo.Room, uuid uuid.UUID) CreateRoomService {
	return &createRoom{
		repo,
		uuid,
	}
}
