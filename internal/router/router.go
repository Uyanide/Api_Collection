package router

import (
	"github.com/Uyanide/Api_Collection/internal/handlers"
	"github.com/Uyanide/Api_Collection/internal/middleware"
	"github.com/gin-gonic/gin"
)

// Router manages HTTP routes
type Router struct {
	ipHandler *handlers.IPHandler
}

// NewRouter creates a new router instance
func NewRouter(ipHandler *handlers.IPHandler) *Router {
	return &Router{
		ipHandler: ipHandler,
	}
}

// SetupRoutes configures all routes with middleware
func (r *Router) SetupRoutes() *gin.Engine {
	router := gin.Default()

	// IP routes
	router.GET("/ip", r.ipHandler.GetIP)

	striped := middleware.StripTrailingSlash(router)

	return striped
}
