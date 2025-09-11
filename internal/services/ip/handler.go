package ip_service

import (
	"net"
	"net/http"
	"time"

	"github.com/Uyanide/Api_Collection/internal/logger"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (s *IPService) GetIPHandler(c *gin.Context) {
	startTime := time.Now()
	log := logger.GetLogger()

	// Log the incoming request
	clientIP := c.ClientIP()
	r := c.Request
	log.WithFields(logrus.Fields{
		"method":     r.Method,
		"path":       r.URL.Path,
		"client_ip":  clientIP,
		"user_agent": r.UserAgent(),
	}).Info("Incoming request")

	// Deny non-GET methods
	if r.Method != http.MethodGet {
		log.WithFields(logrus.Fields{
			"method":    r.Method,
			"client_ip": clientIP,
		}).Warn("Method not allowed")
		c.AbortWithStatusJSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
		return
	}

	// Get IP response from service
	response, err := s.GetIP(clientIP)
	if err != nil {
		log.WithFields(logrus.Fields{
			"client_ip": clientIP,
			"error":     err.Error(),
		}).Error("Internal server error in service")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	// Send response
	c.JSON(http.StatusOK, response)

	// Log successful response
	duration := time.Since(startTime)
	log.WithFields(logrus.Fields{
		"client_ip":   clientIP,
		"response_ip": response.IP,
		"duration_ms": duration.Milliseconds(),
		"status":      "success",
	}).Info("Request processed successfully")
}

func (s *IPService) GetIP(clientIP string) (*IPResponse, error) {
	log := logger.GetLogger()

	log.WithField("client_ip", clientIP).Debug("Processing IP request")

	ip := net.ParseIP(clientIP)
	if ip == nil {
		log.WithField("client_ip", clientIP).Warn("Failed to parse client IP, returning as-is")
		// If parsing fails, return the original client IP
		return &IPResponse{IP: clientIP}, nil
	}

	// Check if client IP belongs to local CIDRs
	if s.IsLocalIP(ip) {
		log.WithFields(logrus.Fields{
			"client_ip": clientIP,
			"return_ip": s.localIP,
			"reason":    "client_ip_in_local_cidr",
		}).Info("Returning local IP for client in local CIDR")
		return &IPResponse{IP: s.localIP}, nil
	}

	log.WithFields(logrus.Fields{
		"client_ip": clientIP,
		"return_ip": clientIP,
		"reason":    "external_ip",
	}).Info("Returning client IP for external client")

	return &IPResponse{IP: clientIP}, nil
}

func (s *IPService) IsLocalIP(ip net.IP) bool {
	for _, cidr := range s.localCIDRs {
		if cidr.Contains(ip) {
			return true
		}
	}
	return false
}
