package router

import (
	"github.com/Uyanide/Api_Collection/internal/config"
	"github.com/Uyanide/Api_Collection/internal/middleware"
	file_service "github.com/Uyanide/Api_Collection/internal/services/file"
	ip_service "github.com/Uyanide/Api_Collection/internal/services/ip"
	"github.com/gin-gonic/gin"
)

// Router manages HTTP routes
type Router struct {
	config            *config.Config
	ipHandler         *ip_service.IPHandler
	fileSingleHandler *file_service.FileSingleHandler
}

// NewRouter creates a new router instance
func NewRouter(config *config.Config, ipHandler *ip_service.IPHandler, fileSingleHandler *file_service.FileSingleHandler) *Router {
	return &Router{
		config:            config,
		ipHandler:         ipHandler,
		fileSingleHandler: fileSingleHandler,
	}
}

// SetupRoutes configures all routes with middleware
func (r *Router) SetupRoutes() *gin.Engine {
	router := gin.Default()

	// IP routes
	router.GET("/ip", r.ipHandler.GetIP)

	// File routes
	r.setupFileRoutes(router, r.config.FileMap)

	striped := middleware.StripTrailingSlash(router)

	return striped
}

func (r *Router) setupFileRoutes(router *gin.Engine, fileMap map[string]config.FileEntry) {
	for urlPath := range fileMap {
		path := urlPath
		router.GET(path, func(c *gin.Context) { r.fileSingleHandler.ServeFile(c, path) })
	}
}
