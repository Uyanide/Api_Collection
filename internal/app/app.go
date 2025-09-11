package app

import (
	"net/http"
	"strconv"

	"github.com/Uyanide/Api_Collection/internal/config"
	"github.com/Uyanide/Api_Collection/internal/handlers"
	"github.com/Uyanide/Api_Collection/internal/logger"
	"github.com/Uyanide/Api_Collection/internal/router"
	"github.com/Uyanide/Api_Collection/internal/service"
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

	// Initialize services
	ipService := service.NewIPService(cfg)
	fileSingleService := service.NewFileSingleService(cfg)

	// Initialize handlers
	ipHandler := handlers.NewIPHandler(ipService)
	fileSingleHandler := handlers.NewFileSingleHandler(fileSingleService)

	// Initialize router
	r := router.NewRouter(cfg, ipHandler, fileSingleHandler)
	engine := r.SetupRoutes()

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
		"port":        a.config.Port,
		"local_ip":    a.config.LocalIP,
		"local_cidrs": a.config.LocalCIDRStr,
	}).Info("Starting server")

	addr := ":" + strconv.Itoa(a.config.Port)

	if err := a.engine.Run(addr); err != nil && err != http.ErrServerClosed {
		a.logger.WithError(err).Error("Server failed to start or crashed")
		return err
	}

	return nil
}
