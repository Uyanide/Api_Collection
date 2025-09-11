package ip_service

import "github.com/gin-gonic/gin"

func (s *IPService) setupRoutes(r *gin.Engine) {
	r.GET("/ip", s.getIPHandler)
}
