package proxy_service

import "github.com/gin-gonic/gin"

type ProxyService struct {
	autoCorrectScheme bool
}

func (s *ProxyService) Init(e *gin.Engine) {
	s.loadConfig()
	s.setupRoutes(e)
}
