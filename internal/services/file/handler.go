package file_service

import (
	"fmt"
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
	// if filepath == "" || filepath == "/" {
	// 	c.AbortWithStatusJSON(400, gin.H{"error": "File path is required"})
	// 	log.WithField("urlPath", urlPath).Warn("File path is required")
	// 	return
	// }

	fullPath := dirObj.Path + filepath

	if _, err := os.Stat(fullPath); err != nil {
		c.AbortWithStatusJSON(404, gin.H{"error": "File not found"})
		log.WithFields(logrus.Fields{
			"urlPath":  urlPath,
			"filepath": filepath,
		}).Warn("File not found in directory")
		return
	}

	// If is a directory
	if stat, _ := os.Stat(fullPath); stat.IsDir() {
		files, err := os.ReadDir(fullPath)
		if err != nil {
			c.AbortWithStatusJSON(500, gin.H{"error": "Failed to read directory"})
			return
		}
		html := fmt.Sprintf("<html><body><h2>index of %s%s/ </h2><ul>", urlPath, filepath)
		dirEntries := []string{}
		fileEntries := []string{}
		for _, f := range files {
			name := f.Name()
			link := fmt.Sprintf("%s%s/%s", urlPath, filepath, name)
			if f.IsDir() {
				dirEntries = append(dirEntries, fmt.Sprintf(`<li><a href="%s/">%s/</a></li>`, link, name))
			} else {
				fileEntries = append(fileEntries, fmt.Sprintf(`<li><a href="%s">%s</a></li>`, link, name))
			}
		}
		// Directories first
		for _, entry := range dirEntries {
			html += entry
		}
		for _, entry := range fileEntries {
			html += entry
		}
		html += "</ul></body></html>"
		c.Data(200, "text/html; charset=utf-8", []byte(html))
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
