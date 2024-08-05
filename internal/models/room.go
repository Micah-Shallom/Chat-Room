package models

import (
	"time"

	"github.com/hngprojects/hng_boilerplate_golang_web/pkg/repository/storage/postgresql"
	"gorm.io/gorm"
)

type Room struct {
	ID          string    `gorm:"type:uuid;primary_key" json:"room_id"`
	Name        string    `gorm:"column:name; type:text; not null" json:"name"`
	Description string    `gorm:"column:description; type:text; not null" json:"description"`
	Users       []User    `gorm:"many2many:user_rooms;" json:"users"`
	CreatedAt   time.Time `gorm:"column:created_at; not null; autoCreateTime" json:"created_at"`
	DeletedAt   time.Time `gorm:"column: deleted_at; not null; autoDeleteTime" json:"deleted_at"`
}

type UserRoom struct {
	RoomID    string    `gorm:"type:uuid;primaryKey;not null" json:"room_id"`
	UserID    string    `gorm:"type:uuid;primaryKey;not null" json:"user_id"`
	CreatedAt time.Time `gorm:"column:created_at;not null;autoCreateTime" json:"created_at"`
	DeletedAt time.Time `gorm:"index" json:"deleted_at"`
}

type CreateRoomRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

func (r *Room) CreateRoom(db *gorm.DB) error {
	err := postgresql.CreateOneRecord(db, r)
	if err != nil {
		return err
	}
	return nil
}

func (r *Room) GetRoomByID(db *gorm.DB, roomID string) (Room, error) {
	var room Room 

	err, nerr := postgresql.SelectOneFromDb(db, &room, "room_id = ?", roomID)
	if err != nil {
		return room, nerr
	}
	return room, nil
}
