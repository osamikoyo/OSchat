package app

import (
	"log/slog"
	"os"
	"oschat/internal/server"
)

func App(){
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	s := server.New()
	err := s.Run()
	if err != nil {
		logger.Error(err.Error())
	}
}