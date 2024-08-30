package database

type User struct{
	Username string `json:"username"`
	Password string `json:"password"`
	Email 	 string `json:"email"`
}
type Chat struct{
	FirstUser string
	SecondUser string
	Data []Message
}
type Message struct{
	Date string
	FromUser string
	ToUser string
	Message string
}