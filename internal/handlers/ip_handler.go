package handlers

import (
	"net/http"
	"time"

	"github.com/Uyanide/Api_Collection/internal/logger"
	"github.com/Uyanide/Api_Collection/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// IPHandler handles IP-related HTTP requests
type IPHandler struct {
	ipService *service.IPService
	logger    *logrus.Logger
}

// NewIPHandler creates a new IPHandler instance
func NewIPHandler(ipService *service.IPService) *IPHandler {
	return &IPHandler{
		ipService: ipService,
		logger:    logger.GetLogger(),
	}
}

// GetIP handles GET /ip requests
func (h *IPHandler) GetIP(c *gin.Context) {
	startTime := time.Now()

	// Log the incoming request
	clientIP := c.ClientIP()
	r := c.Request
	h.logger.WithFields(logrus.Fields{
		"method":     r.Method,
		"path":       r.URL.Path,
		"client_ip":  clientIP,
		"user_agent": r.UserAgent(),
	}).Info("Incoming request")

	// Deny non-GET methods
	if r.Method != http.MethodGet {
		h.logger.WithFields(logrus.Fields{
			"method":    r.Method,
			"client_ip": clientIP,
		}).Warn("Method not allowed")
		c.AbortWithStatusJSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
		return
	}

	// Get IP response from service
	response, err := h.ipService.GetIP(clientIP)
	if err != nil {
		h.logger.WithFields(logrus.Fields{
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
	h.logger.WithFields(logrus.Fields{
		"client_ip":   clientIP,
		"response_ip": response.IP,
		"duration_ms": duration.Milliseconds(),
		"status":      "success",
	}).Info("Request processed successfully")
}
