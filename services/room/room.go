package room

import (
	"errors"
	"net/http"

	"gorm.io/gorm"

	"github.com/hngprojects/hng_boilerplate_golang_web/internal/models"
	"github.com/hngprojects/hng_boilerplate_golang_web/utility"
)

func GetRooms(db *gorm.DB) ([]models.Room, error) {
	var room models.Room

	rooms, err := room.GetRooms(db)
	if err != nil {
		return rooms, err
	}
	return rooms, nil
}

func CreateRoom(req models.CreateRoomRequest, db *gorm.DB, user_id string) (models.Room, error) {

	room := models.Room{
		ID:          utility.GenerateUUID(),
		Name:        req.Name,
		Description: req.Description,
	}

	err := room.CreateRoom(db)
	if err != nil {
		return room, err
	}

	err = room.AddUserToRoom(db, room.ID, user_id)
	if err != nil {
		return room, err
	}

	return room, nil
}

func GetRoom(db *gorm.DB, roomID string) (models.Room, error) {
	var room models.Room

	fetchedRoom, err := room.GetRoomByID(db, roomID)
	if err != nil {
		return fetchedRoom, err
	}
	return fetchedRoom, nil
}

func GetRoomMsg(roomId string, db *gorm.DB) ([]models.Message, int, error) {
	var message models.Message

	resp, err := message.GetMessagesByRoomID(db, roomId)

	if err != nil {
		return []models.Message{}, http.StatusInternalServerError, err
	}

	return resp, http.StatusOK, nil

}

func JoinRoom(db *gorm.DB, room_id, user_id string) error {
	var room models.Room

	_, err := GetRoom(db, room_id)
	if err != nil {
		return errors.New("room does not exist")
	}

	err = room.AddUserToRoom(db, room_id, user_id)
	if err != nil {
		return err
	}
	return nil
}

func LeaveRoom(db *gorm.DB, room_id, user_id string) error {
	var room models.Room

	_, err := GetRoom(db, room_id)
	if err != nil {
		return errors.New("room does not exist")
	}

	err = room.RemoveUserFromRoom(db, room_id, user_id)
	if err != nil {
		return err
	}
	return nil

}

func AddRoomMsg(req models.CreateMessageRequest, db *gorm.DB) (int, error) {
	message := models.Message{
		Content: req.Content,
		RoomID:  req.RoomId,
		UserID:  req.UserId,
	}

	err := message.CreateMessage(db)

	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusCreated, nil
}
