package file_service

import (
	"os"
	"strings"

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

	increaseCounter(FileDownloadsKeyPrefix + urlPath)
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

func (s *FileService) serveDirFile(c *gin.Context, urlPath string) {
	log := logger.GetLogger()

	log.WithField("handler", "DownloadDirFile").Debug("Handling directory file download request")

	dirObj, exists := s.dirMap[urlPath]
	if !exists || dirObj.Path == "" {
		c.AbortWithStatusJSON(404, gin.H{"error": "Directory not found"})
		log.WithField("urlPath", urlPath).Warn("Directory not found")
		return
	}

	filepath := c.Param("filepath")
	if filepath == "" || filepath == "/" {
		c.AbortWithStatusJSON(400, gin.H{"error": "File path is required"})
		log.WithField("urlPath", urlPath).Warn("File path is required")
		return
	}

	fullPath := dirObj.Path + filepath
	if _, err := os.Stat(fullPath); err != nil {
		c.AbortWithStatusJSON(404, gin.H{"error": "File not found"})
		log.WithFields(logrus.Fields{
			"urlPath":  urlPath,
			"filepath": filepath,
		}).Warn("File not found in directory")
		return
	}

	splitted := strings.Split(filepath, "/")
	fileName := splitted[len(splitted)-1]
	c.FileAttachment(fullPath, fileName)
	log.WithFields(logrus.Fields{
		"url_path": urlPath,
		"filepath": filepath,
	}).Info("Request processed successfully")

	increaseCounter(FileDownloadsKeyPrefix + urlPath + filepath)
}
