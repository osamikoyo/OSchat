package server

import (
	"oschat/internal/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct{
	*echo.Echo
}
func New() Server{
	e := echo.New()
	return Server{e}
}
func (s Server) Run() error {
	s.Use(middleware.Logger())

	s.GET("/", handler.Home)
	err := s.Start(":2020")
	if err != nil {
		return err
	}
	return  nil
}