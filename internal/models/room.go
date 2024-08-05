package models

import (
	"errors"
	"time"

	"github.com/hngprojects/hng_boilerplate_golang_web/internal/models"
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

func (r *Room) GetRooms(db *gorm.DB) ([]Room, error) {
	var rooms []Room

	err := postgresql.SelectAllFromDb(db, "", &rooms, "")
	if err != nil {
		return rooms, err
	}
	return rooms, nil
}

func (r *Room) GetRoomMessages(db *gorm.DB, roomID string) ([]models.Message, error) {
	//query the room and get all the messages
	var messages []models.Message

	err := postgresql.SelectAllFromDb(
		db.Where("room_id = ?", roomID),
		"",
		&messages,
		"room_id = ?",
		roomID,
	)
	if err != nil {
		return messages, err
	}

	return messages, nil
}

func (r *Room) AddUserToRoom(db *gorm.DB, roomID, userID string) error {
	//add user to room
	//check if user is already in room
	//if user is not in room, add user to room

	var user models.User
	var room Room

	_, err := user.GetUserByID(db, userID)
	if err != nil {
		return errors.New("user does not exist")
	}

	_, err = room.GetRoomByID(db, roomID)
	if err != nil {
		return errors.New("room does not exist")
	}

	var userRoom models.UserRoom
	err, _ = postgresql.SelectOneFromDb(db, &userRoom, "room_id = ? AND user_id = ?", roomID, userID)
	if err != nil {
		return errors.New("user already in room")
	}

	//if user not in room, add user to room
	userRoom = models.UserRoom{
		RoomID: roomID,
		UserID: userID,
	}

	err = postgresql.CreateOneRecord(db, &userRoom)
	if err != nil {
		return errors.New("could not add user to room")
	}
	return nil
}

func (r *Room) RemoveUserFromRoom(db *gorm.DB, roomID, userID string) error {
	//remove user from room
	//check if user is in room
	//if user is in room, remove user from room
	var userRoom models.UserRoom

	err, _ := postgresql.SelectOneFromDb(db, &userRoom, "room_id = ? AND user_id = ?", roomID, userID)
	if err != nil {
		return errors.New("user not in room")
	}

	err = postgresql.DeleteRecordFromDb(db, &userRoom)
	if err != nil {
		return errors.New("could not remove user from room")
	}
	return nil
}
