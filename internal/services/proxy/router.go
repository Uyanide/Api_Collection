package proxy_service

import (
	"github.com/Uyanide/Api_Collection/internal/middleware"
	"github.com/gin-gonic/gin"
)

func (s *ProxyService) setupRoutes(r *gin.Engine) {
	group := r.Group("/proxy", middleware.CORSMiddleware())
	group.Any("", s.handleProxy)
}
