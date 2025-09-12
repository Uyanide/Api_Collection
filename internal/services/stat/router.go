package stats_service

import "github.com/gin-gonic/gin"

func (s *StatsService) setupRoutes(engine *gin.Engine) {
	statsGroup := engine.Group("/stats")
	statsGroup.GET("", s.getGeneralStatsHandler)
}
