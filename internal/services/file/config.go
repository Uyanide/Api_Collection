package file_service

import (
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

func (s *FileService) loadConfig() {
	log := logrus.New()

	// Parse file map
	s.fileMap = make(map[string]fileObject)
	fileMapEnv := os.Getenv("FILE_MAP")
	if fileMapEnv != "" {
		mappings := strings.Split(fileMapEnv, ",")
		for _, mapping := range mappings {
			parts := strings.SplitN(mapping, ":", 3)
			if len(parts) != 3 {
				log.WithField("mapping", mapping).Warn("Invalid FILE_MAP entry, skipping")
				continue
			}
			urlPath := strings.TrimSpace(parts[0])
			filePath := strings.TrimSpace(parts[1])
			fileName := strings.TrimSpace(parts[2])
			if urlPath == "" || filePath == "" || fileName == "" {
				log.WithField("mapping", mapping).Warn("Empty URL path, file path or file name in FILE_MAP entry, skipping")
				continue
			}
			s.fileMap[urlPath] = fileObject{
				Path: filePath,
				Name: fileName,
			}

			log.WithFields(logrus.Fields{
				"url_path":  urlPath,
				"file_path": filePath,
				"file_name": fileName,
			}).Info("Parsed FILE_MAP entry")
		}
	}
}
