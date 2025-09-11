package services

import (
	file_service "github.com/Uyanide/Api_Collection/internal/services/file"
	ip_service "github.com/Uyanide/Api_Collection/internal/services/ip"
	"github.com/gin-gonic/gin"
)

type GeneralService interface {
	Init(*gin.Engine)
}

// NewServices creates and initializes all services
func NewServices(e *gin.Engine) {
	services := []GeneralService{
		&file_service.FileService{},
		&ip_service.IPService{},
	}
	for _, service := range services {
		service.Init(e)
	}
}
