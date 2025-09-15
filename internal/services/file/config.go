package file_service

import (
	"os"
	"strings"

	"github.com/Uyanide/Api_Collection/internal/db"
	"github.com/sirupsen/logrus"
)

func (s *FileService) loadConfig() {

	// Parse file map
	s.loadFileConfig()

	// Parse directory map
	s.loadDirConfig()

	// Load keys from database
	s.loadKeysFromDB()
}

func (s *FileService) loadFileConfig() {
	log := logrus.New()
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

func (s *FileService) loadDirConfig() {
	log := logrus.New()
	s.dirMap = make(map[string]dirObject)
	dirMapEnv := os.Getenv("DIR_MAP")
	if dirMapEnv != "" {
		mappings := strings.Split(dirMapEnv, ",")
		for _, mapping := range mappings {
			parts := strings.SplitN(mapping, ":", 2)
			if len(parts) != 2 {
				log.WithField("mapping", mapping).Warn("Invalid DIR_MAP entry, skipping")
				continue
			}
			urlPath := strings.TrimSpace(parts[0])
			dirPath := strings.TrimSpace(parts[1])
			if urlPath == "" || dirPath == "" {
				log.WithField("mapping", mapping).Warn("Empty URL path or directory path in DIR_MAP entry, skipping")
				continue
			}
			s.dirMap[urlPath] = dirObject{
				Path: dirPath,
			}

			log.WithFields(logrus.Fields{
				"url_path": urlPath,
				"dir_path": dirPath,
			}).Info("Parsed DIR_MAP entry")
		}
	}
}

func (s *FileService) loadKeysFromDB() {
	dbInst := db.GetDB()

	keys, err := dbInst.Keys()
	if err != nil {
		logrus.WithError(err).Error("Failed to load keys from database")
		return
	}
	for _, key := range keys {
		if strings.HasPrefix(key, FileDownloadsKeyPrefix) {
			FileDownloadsKeys[key] = struct{}{}
		}
	}
	logrus.WithField("count", len(FileDownloadsKeys)).Info("Loaded file download keys from database")
}
