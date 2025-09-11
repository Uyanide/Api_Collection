package file_service

import (
	"github.com/gin-gonic/gin"
)

type FileService struct {
	fileMap map[string]FileObject
}

func (s *FileService) Init(e *gin.Engine) {
	s.loadConfig()
	s.setupRoutes(e)
}
