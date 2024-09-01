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
	var u1 database.User
	var us database.User

	erre := c.Bind(&us)


	if erru := db.Where("email = ?", us.Email).First(&u1).Error; erru != nil {
		return erru
	}

	var u2 database.User
	if errs := db.Where("password = ? OR email = ?", us.Password, us.Email).First(&u2).Error; errs != nil {
		return errs
	}
	loger.Info(u2.Email, u2.Password, nil)
	if erre != nil || u1.Password != us.Password {
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

