package database

type User struct{
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	Email 	 string `json:"email" form:"email"`
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
