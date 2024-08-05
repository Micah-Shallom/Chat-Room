package room

import (
	"github.com/hngprojects/hng_boilerplate_golang_web/internal/models"
	"github.com/hngprojects/hng_boilerplate_golang_web/pkg/repository/storage/postgresql"
	"github.com/hngprojects/hng_boilerplate_golang_web/utility"
	"gorm.io/gorm"
)

func GetRooms(db *gorm.DB) ([]models.Room,error){
	var rooms []models.Room


	err := postgresql.SelectAllFromDb(db, "", &rooms, "")
	if err != nil {
		
	}

}

func CreateRoom(req models.CreateRoomRequest,db *gorm.DB) (models.Room,error ){
	
	room := models.Room{
		ID: utility.GenerateUUID(),
		Name: req.Name,
		Description: req.Description,
	}

	err := room.CreateRoom(db); 
	if err != nil {
		return room, err
	}
	return room, nil
}

func GetRoom(db *gorm.DB) {

}

func GetRoomMsg(db *gorm.DB) {

}

func JoinRoom(db *gorm.DB) {

}

func LeaveRoom(db *gorm.DB) {

}
