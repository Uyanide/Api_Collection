package config

import (
	"os"
	"strconv"

	"github.com/Uyanide/Api_Collection/internal/logger"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Port int
}

func NewConfig() *Config {
	log := logger.GetLogger()
	log.Info("Initializing configuration")

	if err := godotenv.Load(); err != nil {
		log.Warn("No .env file found")
	}

	// Parse port
	portString := os.Getenv("PORT")
	if portString == "" {
		portString = "10087"
		log.WithField("port", portString).Warn("No PORT environment variable set, using default")
	}
	port, err := strconv.Atoi(portString)
	if err != nil || port <= 0 || port > 65535 {
		log.WithFields(logrus.Fields{
			"port":  portString,
			"error": err.Error(),
		}).Fatal("Invalid port configuration")
	}

	config := &Config{
		Port: port,
	}

	log.Info("Configuration initialized successfully")

	return config
}
