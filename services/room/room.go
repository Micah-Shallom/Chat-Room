package room

import (
	"net/http"

	"gorm.io/gorm"

	"github.com/hngprojects/hng_boilerplate_golang_web/internal/models"
	"github.com/hngprojects/hng_boilerplate_golang_web/pkg/repository/storage/postgresql"
	"github.com/hngprojects/hng_boilerplate_golang_web/utility"
)

func GetRooms(db *gorm.DB) ([]models.Room, error) {
	var rooms []models.Room

	err := postgresql.SelectAllFromDb(db, "", &rooms, "")
	if err != nil {

	}

}

func CreateRoom(req models.CreateRoomRequest, db *gorm.DB) (models.Room, error) {

	room := models.Room{
		ID:          utility.GenerateUUID(),
		Name:        req.Name,
		Description: req.Description,
	}

	err := room.CreateRoom(db)
	if err != nil {
		return room, err
	}
	return room, nil
}

func GetRoom(db *gorm.DB) {

}

func GetRoomMsg(roomId string, db *gorm.DB) ([]models.Message, int, error) {
	var message models.Message

	resp, err := message.GetMessagesByRoomID(db, roomId)

	if err != nil {
		return []models.Message{}, http.StatusInternalServerError, err
	}

	return resp, http.StatusOK, nil

}

func JoinRoom(db *gorm.DB) {

}

func LeaveRoom(db *gorm.DB) {

}

func AddRoomMsg(req models.CreateMessageRequest, db *gorm.DB) (int, error) {

	var message models.Message

	message = models.Message{
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
