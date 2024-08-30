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
	db, err := gorm.Open(sqlite.Open("storage/main.db"))
	loger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	email := c.FormValue("email")
	password := c.FormValue("password")

	var dbPassword string
	var u database.User
	if err := db.Where("password = ?", password).First(&u).Error; err != nil {
		return err
	}

	if err != nil || dbPassword != password {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid credentials")
	}

	claims := jwt.MapClaims{}
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		loger.Error(err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, "Could not generate token")

	}

	return c.JSON(http.StatusOK, map[string]string{"token": tokenString})
}

