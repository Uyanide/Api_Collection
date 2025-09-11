package file_service

import (
	"os"

	"github.com/Uyanide/Api_Collection/internal/config"
)

type FileSingleService struct {
	config *config.Config
}

func NewFileSingleService(cfg *config.Config) *FileSingleService {
	return &FileSingleService{
		config: cfg,
	}
}

func (s *FileSingleService) GetFilePath(key string) (string, bool) {
	obj, exists := s.config.FileMap[key]
	if !exists {
		return "", false
	}
	if _, err := os.Stat(obj.Path); err != nil {
		return "", false
	}

	return obj.Path, true
}

func (s *FileSingleService) GetFileName(key string) (string, bool) {
	obj, exists := s.config.FileMap[key]
	if !exists {
		return "", false
	}
	return obj.Name, true
}
