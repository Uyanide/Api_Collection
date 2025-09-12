package file_service

import (
	"github.com/gin-gonic/gin"
)

var (
	FileDownloadsKeyPrefix = "file_downloads_"
	FileDownloadsKeys      = []string{}
)

type FileService struct {
	fileMap map[string]fileObject
}

func (s *FileService) Init(e *gin.Engine) {
	s.loadConfig()
	s.setupRoutes(e)
}
