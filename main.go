package main

import (
	"github.com/Uyanide/Api_Collection/internal/app"
	"github.com/Uyanide/Api_Collection/internal/logger"
)

func main() {
	logger.InitLogger()
	log := logger.GetLogger()

	log.Info("Starting server")

	application := app.NewApp()

	cleanup, err := application.Start()
	defer cleanup()
	if err != nil {
		log.WithError(err).Fatal("Failed to start server")
	}
}
