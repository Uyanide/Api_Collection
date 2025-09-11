package file_service

import (
	"os"

	"github.com/Uyanide/Api_Collection/internal/logger"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (s *FileService) serveFile(c *gin.Context, urlPath string) {
	log := logger.GetLogger()

	log.WithField("handler", "DownloadFile").Debug("Handling file download request")

	filePath, exists := s.getFilePath(urlPath)
	if !exists || filePath == "" {
		c.AbortWithStatusJSON(404, gin.H{"error": "File not found"})
		log.WithField("urlPath", urlPath).Warn("File not found")
		return
	}
	fileName, _ := s.getFileName(urlPath) // already checked existence

	c.FileAttachment(filePath, fileName)
	log.WithFields(logrus.Fields{
		"url_path":  urlPath,
		"file_path": filePath,
		"file_name": fileName,
	}).Info("Request processed successfully")
}

func (s *FileService) getFilePath(key string) (string, bool) {
	obj, exists := s.fileMap[key]
	if !exists {
		return "", false
	}
	if _, err := os.Stat(obj.Path); err != nil {
		return "", false
	}

	return obj.Path, true
}

func (s *FileService) getFileName(key string) (string, bool) {
	obj, exists := s.fileMap[key]
	if !exists {
		return "", false
	}
	return obj.Name, true
}
