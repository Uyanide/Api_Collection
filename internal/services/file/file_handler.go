package file_service

import (
	"github.com/Uyanide/Api_Collection/internal/logger"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type FileSingleHandler struct {
	fileService *FileSingleService
}

func NewFileSingleHandler(fileService *FileSingleService) *FileSingleHandler {
	return &FileSingleHandler{
		fileService: fileService,
	}
}

func (h *FileSingleHandler) ServeFile(c *gin.Context, urlPath string) {
	log := logger.GetLogger()

	log.WithField("handler", "DownloadFile").Info("Handling file download request")

	filePath, exists := h.fileService.GetFilePath(urlPath)
	if !exists || filePath == "" {
		c.AbortWithStatusJSON(404, gin.H{"error": "File not found"})
		log.WithField("urlPath", urlPath).Warn("File not found")
		return
	}
	fileName, _ := h.fileService.GetFileName(urlPath) // already checked existence

	c.FileAttachment(filePath, fileName)
	log.WithFields(logrus.Fields{
		"url_path":  urlPath,
		"file_path": filePath,
		"file_name": fileName,
	}).Info("Request processed successfully")
}
