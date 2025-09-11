package app

import (
	"net/http"
	"strconv"

	"github.com/Uyanide/Api_Collection/internal/config"
	"github.com/Uyanide/Api_Collection/internal/logger"
	"github.com/Uyanide/Api_Collection/internal/middleware"
	"github.com/Uyanide/Api_Collection/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// App represents the main application
type App struct {
	config *config.Config
	engine *gin.Engine
	logger *logrus.Logger
}

// NewApp creates a new application instance
func NewApp() *App {
	cfg := config.NewConfig()
	log := logger.GetLogger()

	log.Info("Initializing application components")

	// Initialize Gin engine
	engine := gin.Default()
	engine.Use(middleware.StripTrailingSlash())

	// Initialize services
	services.NewServices(engine)

	log.Info("Application components initialized successfully")

	return &App{
		config: cfg,
		engine: engine,
		logger: log,
	}
}

// Start starts the application server
func (a *App) Start() error {
	a.logger.WithFields(logrus.Fields{
		"port": a.config.Port,
	}).Info("Starting server")

	addr := ":" + strconv.Itoa(a.config.Port)

	if err := a.engine.Run(addr); err != nil && err != http.ErrServerClosed {
		a.logger.WithError(err).Error("Server failed to start or crashed")
		return err
	}

	return nil
}
