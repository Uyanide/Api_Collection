package stats_service

import (
	"github.com/gin-gonic/gin"
)

type StatsService struct{}

func (s *StatsService) Init(engine *gin.Engine) {
	s.setupRoutes(engine)
}
