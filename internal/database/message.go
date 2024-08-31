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
func FindMessageMore(fu User, su User, lastmessage Message) ([]Message, error){
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
	
	Msg = TrimSlice(Message, lastmessage.Date)

	return Message, nil
} 
func TrimSlice(slice []Message, threshold string) []Message {
	var trimmed []Message
	started := false

	for _, item := range slice {
		if item.Date== threshold {
			started = true
		}
		if started {
			trimmed = append(trimmed, item)
		}
	}

	return trimmed
}
func AddMessage(msg Message) error{
	db, err := gorm.Open(sqlite.Open("storage/chats.db"), &gorm.Config{})
	loger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	

	if err != nil {
		loger.Error(err.Error())
		return err
	}

	res := db.Create(&msg)
	if res.Error != nil {
		loger.Error(res.Error.Error())
	}
	return err
}
type ChatDB struct{
	FirstUser string  `json:"firstuser"`
	SecondUser string `json:"seconduser"`
	Data []Message    `json:"message"`
}
type Message struct{
	FirstUser string  `json:"firstuser"`
	Message string    `json:"datamsg"`
	Date string	      `json:"date"`
	SecondUser string `json:"seconduser"`
}
type Chat struct{
	Firstuser string  `json:"firstuser"`
	Seconduser string `json:"seconduser"`
}