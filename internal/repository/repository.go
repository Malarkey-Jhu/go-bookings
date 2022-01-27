package repository

import (
	"time"

	"github.com/Malarkey-Jhu/go-bookings/internal/models"
)

type DataBaseRepo interface {
	AllUsers() bool

	InsertReservation(res models.Reservation) (int, error)
	InsertRoomRestriction(r models.RoomRestriction) error
	SearchAvailablityByDatesByRoomID(start, end time.Time, roomID int) (bool, error)
	SearchAvailablityForAllRooms(start, end time.Time) ([]models.Room, error)
	GetRoomByID(id int) (models.Room, error)
}
