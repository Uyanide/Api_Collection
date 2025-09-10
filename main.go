package main

import (
	"github.com/Uyanide/Api_Collection/internal/app"
	"github.com/Uyanide/Api_Collection/internal/logger"
)

func main() {
	logger.InitLogger()
	log := logger.GetLogger()

	log.Info("Starting application")

	application := app.NewApp()

	if err := application.Start(); err != nil {
		log.WithError(err).Fatal("Failed to start server")
	}
}
