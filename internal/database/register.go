package database

type User struct{
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	Email 	 string `json:"email" form:"email"`
}
type ChatDB struct{
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
type Chat struct{
	LastMassage string
	WithUser string
}