package handler

import (
	"log/slog"
	"net/http"
	"os"
	"oschat/internal/database"
	"oschat/internal/servies"
	"time"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var jwtSecret = servies.GeterateJWTkey()
func Home(c echo.Context) error{
	return c.File("static/index.html")
}
func Register(c echo.Context) error {
	db, err := gorm.Open(sqlite.Open("storage/main.db"))
	if err != nil {
		return err
	}
	loger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	var User database.User
	err = c.Bind(&User)
	if err != nil {
		loger.Error(err.Error())
		return err
	}

	res := db.Create(&User)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
func Login(c echo.Context) error {
	db, errd := gorm.Open(sqlite.Open("storage/main.db"))
	if errd != nil {
		return errd
	}

	loger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	
	var us database.User

	erre := c.Bind(&us)

	var dbPassword string
	var u database.User
	if errs := db.Where("password = ?", us.Password).First(&u).Error; errs != nil {
		return errs
	}

	if erre != nil || dbPassword != us.Password {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid credentials")
	}

	claims := jwt.MapClaims{}
	claims["email"] = us.Email
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		loger.Error(err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, "Could not generate token")

	}

	return c.JSON(http.StatusOK, map[string]string{"token": tokenString})
}

func GetChats(c echo.Context) error {
	loger := slog.New(slog.NewJSONHandler(os.Stdout, nil))


	var User database.User
	var Chats []database.Chat
	err := c.Bind(&User)
	if err != nil {
		loger.Error(err.Error())
		return nil
	}
	Chats, err = database.FindChat(User)
	if err != nil {
		loger.Error(err.Error())
	}
	return c.JSON(http.StatusOK, Chats)
}
