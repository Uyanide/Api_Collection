package proxy_service

import (
	"os"

	"github.com/Uyanide/Api_Collection/internal/logger"
)

func (s *ProxyService) loadConfig() {
	log := logger.GetLogger()

	s.autoCorrectScheme = os.Getenv("AUTO_CORRECT_SCHEME") == "1"
	log.WithField("auto_correct_scheme", s.autoCorrectScheme).Info("Proxy service configuration loaded")
}
