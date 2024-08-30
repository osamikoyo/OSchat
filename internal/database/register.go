package database

type User struct{
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	Email 	 string `json:"email" form:"email"`
}
