package service

import (
	"net"

	"github.com/Uyanide/Api_Collection/internal/config"
	"github.com/Uyanide/Api_Collection/internal/logger"
	"github.com/Uyanide/Api_Collection/internal/models"
	"github.com/sirupsen/logrus"
)

// IPService handles IP-related business logic
type IPService struct {
	config *config.Config
	logger *logrus.Logger
}

// NewIPService creates a new IPService instance
func NewIPService(cfg *config.Config) *IPService {
	return &IPService{
		config: cfg,
		logger: logger.GetLogger(),
	}
}

// GetIP determines the IP to return based on client IP
func (s *IPService) GetIP(clientIP string) (*models.IPResponse, error) {
	s.logger.WithField("client_ip", clientIP).Debug("Processing IP request")

	ip := net.ParseIP(clientIP)
	if ip == nil {
		s.logger.WithField("client_ip", clientIP).Warn("Failed to parse client IP, returning as-is")
		// If parsing fails, return the original client IP
		return &models.IPResponse{IP: clientIP}, nil
	}

	// Check if client IP belongs to local CIDRs
	if s.config.IsLocalIP(ip) {
		s.logger.WithFields(logrus.Fields{
			"client_ip": clientIP,
			"return_ip": s.config.LocalIP,
			"reason":    "client_ip_in_local_cidr",
		}).Info("Returning local IP for client in local CIDR")
		return &models.IPResponse{IP: s.config.LocalIP}, nil
	}

	s.logger.WithFields(logrus.Fields{
		"client_ip": clientIP,
		"return_ip": clientIP,
		"reason":    "external_ip",
	}).Info("Returning client IP for external client")

	return &models.IPResponse{IP: clientIP}, nil
}
