package services

import (
	file_service "github.com/Uyanide/Api_Collection/internal/services/file"
	ip_service "github.com/Uyanide/Api_Collection/internal/services/ip"
	proxy_service "github.com/Uyanide/Api_Collection/internal/services/proxy"
	stats_service "github.com/Uyanide/Api_Collection/internal/services/stat"
	"github.com/gin-gonic/gin"
)

type GeneralService interface {
	Init(*gin.Engine)
}

var Services map[string]GeneralService

// NewServices creates and initializes all services
func NewServices(e *gin.Engine) {
	Services = map[string]GeneralService{
		"file":  &file_service.FileService{},
		"ip":    &ip_service.IPService{},
		"proxy": &proxy_service.ProxyService{},
		"stats": &stats_service.StatsService{},
	}
	for _, service := range Services {
		service.Init(e)
	}
}
