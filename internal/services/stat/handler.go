package stats_service

import (
	"github.com/Uyanide/Api_Collection/internal/logger"
	file_service "github.com/Uyanide/Api_Collection/internal/services/file"
	ip_service "github.com/Uyanide/Api_Collection/internal/services/ip"
	proxy_service "github.com/Uyanide/Api_Collection/internal/services/proxy"
	"github.com/gin-gonic/gin"
)

func (s *StatsService) getGeneralStatsHandler(c *gin.Context) {
	log := logger.GetLogger()

	statIP, err := ip_service.ConstructStatsIP()
	if err != nil {
		log.WithError(err).Error("Failed to get IP stats")
		c.JSON(500, gin.H{"error": "Failed to get IP stats"})
		return
	}

	statProxy, err := proxy_service.ConstructStatsProxy()
	if err != nil {
		log.WithError(err).Error("Failed to get Proxy stats")
		c.JSON(500, gin.H{"error": "Failed to get Proxy stats"})
		return
	}

	statFile, err := file_service.ConstructStatsFile()
	if err != nil {
		log.WithError(err).Error("Failed to get File stats")
		c.JSON(500, gin.H{"error": "Failed to get File stats"})
		return
	}

	c.JSON(200, StatsGeneralResponse{
		TotalRequests:   statIP.TotalRequests + statFile.TotalDownloads + statProxy.TotalRequests,
		IPRequests:      *statIP,
		FileDownloads:   *statFile,
		ProxiedRequests: *statProxy,
	})
}
