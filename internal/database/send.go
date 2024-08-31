package database

import (
	"log/slog"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)
const newLenth int = 70


func GetMessageEver(fu User) ([]Message, error){
	var Message []Message
	

	db, err := gorm.Open(sqlite.Open("storage/chats.db"))
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	if err != nil {
		logger.Error(err.Error())
		return Message, err
	}
	if err = db.Where("firstuser = ? OR seconduser = ?", fu.Username, fu.Username).Find(&Message).Error
	err != nil{
		return Message, err
	}
	msg := Message[:newLenth]
	return msg, err
	
}