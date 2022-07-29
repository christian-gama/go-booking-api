package service

import (
	"github.com/christian-gama/go-booking-api/internal/room/app/repo"
	"github.com/christian-gama/go-booking-api/internal/room/domain/entity"
	"github.com/christian-gama/go-booking-api/internal/shared/app/exception"
	"github.com/christian-gama/go-booking-api/internal/shared/app/uuid"
)

type CreateRoomService interface {
	Handle(input *CreateRoomInput) (*entity.Room, *exception.Error)
}

type CreateRoomInput struct {
	Name        string
	Description string
	BedCount    uint8
	Price       float32
}

type createRoom struct {
	repo      repo.Room
	exception exception.Exception
	uuid      uuid.UUID
}

// Handle receives an input and creates a room. It will return an error if something
// goes wrong with room creation or if the room repo return an error.
func (c *createRoom) Handle(input *CreateRoomInput) (*entity.Room, *exception.Error) {
	uuid := c.uuid.Generate()
	room, err := entity.NewRoom(uuid, input.Name, input.Description, input.BedCount, input.Price, false)
	if err != nil {
		return nil, c.exception.BadRequest(err.Error())
	}

	room, exception := c.repo.SaveRoom(room)
	if exception != nil {
		return nil, exception
	}

	return room, nil
}

func NewCreateRoom(repo repo.Room, exception exception.Exception, uuid uuid.UUID) CreateRoomService {
	return &createRoom{
		repo,
		exception,
		uuid,
	}
}
