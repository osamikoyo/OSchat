package database

import (
	"log/slog"
	"os"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func FindChat(u User) ([]Chat, error){
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	var Chats []Chat
	db, err := gorm.Open(sqlite.Open("storage/chats.db"))
	if err != nil {
		return []Chat{}, err
	}
	if err := db.Where("seconduser = ? OR firstuser = ?", u.Username, u.Username).Find(&Chats).Error; err != nil {
        logger.Error(err.Error())
    }
	return Chats, nil
}
func FindMessages(fu User, su User) ([]Message, error){
	var Message []Message

	db, err := gorm.Open(sqlite.Open("storage/chats.db"))
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	if err != nil {
		logger.Error(err.Error())
		return Message, err
	}
	if err = db.Where("(value = ? OR value2 = ?) AND (value = ? OR value2 = ?)", fu.Username, fu.Username, su.Username, su.Username).Find(&Message).Error
	err != nil{
		return Message, err
	}
	return Message, err
	
}
var Msg []Message
func FindMessageMore(fu User, su User, lastmessagedate string) ([]Message, error){
	var Message []Message

	loger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	db, err := gorm.Open(sqlite.Open("storage"))
	if err != nil {
		loger.Error(err.Error())
		return Message, err
	}

	if err = db.Where("(value = ? OR value2 = ?) AND (value = ? OR value2 = ?)", fu.Username, fu.Username, su.Username, su.Username).Find(&Message).Error
	err != nil{
		return Message, err
	}
	


	return Message, nil
} 
type ChatDB struct{
	FirstUser string
	SecondUser string
	Data []Message
}
type Message struct{
	FirstUser string
	Message string
	Date string
	SendedUser string
}
type Chat struct{
	Firstuser string `json:"firstuser"`
	Seconduser string `json:"seconduser"`
}