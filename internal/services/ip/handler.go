package ip_service

import (
	"net"
	"net/http"

	"github.com/Uyanide/Api_Collection/internal/logger"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (s *IPService) getIPHandler(c *gin.Context) {
	log := logger.GetLogger()

	// Log the incoming request
	clientIP := c.ClientIP()
	r := c.Request
	log.WithFields(logrus.Fields{
		"method":     r.Method,
		"path":       r.URL.Path,
		"client_ip":  clientIP,
		"user_agent": r.UserAgent(),
	}).Debug("Incoming request")

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
	response, err := s.getIP(clientIP)
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
	log.WithFields(logrus.Fields{
		"client_ip":   clientIP,
		"response_ip": response.IP,
		"status":      "success",
	}).Info("Request processed successfully")
}

func (s *IPService) getIP(clientIP string) (*ipResponse, error) {
	log := logger.GetLogger()

	log.WithField("client_ip", clientIP).Debug("Processing IP request")

	ip := net.ParseIP(clientIP)
	if ip == nil {
		log.WithField("client_ip", clientIP).Warn("Failed to parse client IP, returning as-is")
		// If parsing fails, return the original client IP
		return &ipResponse{IP: clientIP}, nil
	}

	// Check if client IP belongs to local CIDRs
	if s.isLocalIP(ip) {
		log.WithFields(logrus.Fields{
			"client_ip": clientIP,
			"return_ip": s.localIP,
			"reason":    "client_ip_in_local_cidr",
		}).Debug("Returning local IP for client in local CIDR")
		return &ipResponse{IP: s.localIP}, nil
	}

	log.WithFields(logrus.Fields{
		"client_ip": clientIP,
		"return_ip": clientIP,
		"reason":    "external_ip",
	}).Debug("Returning client IP for external client")

	return &ipResponse{IP: clientIP}, nil
}

func (s *IPService) isLocalIP(ip net.IP) bool {
	for _, cidr := range s.localCIDRs {
		if cidr.Contains(ip) {
			return true
		}
	}
	return false
}
