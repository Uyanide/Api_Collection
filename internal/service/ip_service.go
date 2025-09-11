package service

import (
	"net"

	"github.com/Uyanide/Api_Collection/internal/config"
	"github.com/Uyanide/Api_Collection/internal/logger"
	"github.com/Uyanide/Api_Collection/internal/models"
	"github.com/sirupsen/logrus"
)

type IPService struct {
	config *config.Config
}

func NewIPService(cfg *config.Config) *IPService {
	return &IPService{
		config: cfg,
	}
}

// GetIP determines the IP to return based on client IP
func (s *IPService) GetIP(clientIP string) (*models.IPResponse, error) {
	log := logger.GetLogger()

	log.WithField("client_ip", clientIP).Debug("Processing IP request")

	ip := net.ParseIP(clientIP)
	if ip == nil {
		log.WithField("client_ip", clientIP).Warn("Failed to parse client IP, returning as-is")
		// If parsing fails, return the original client IP
		return &models.IPResponse{IP: clientIP}, nil
	}

	// Check if client IP belongs to local CIDRs
	if s.config.IsLocalIP(ip) {
		log.WithFields(logrus.Fields{
			"client_ip": clientIP,
			"return_ip": s.config.LocalIP,
			"reason":    "client_ip_in_local_cidr",
		}).Info("Returning local IP for client in local CIDR")
		return &models.IPResponse{IP: s.config.LocalIP}, nil
	}

	log.WithFields(logrus.Fields{
		"client_ip": clientIP,
		"return_ip": clientIP,
		"reason":    "external_ip",
	}).Info("Returning client IP for external client")

	return &models.IPResponse{IP: clientIP}, nil
}
