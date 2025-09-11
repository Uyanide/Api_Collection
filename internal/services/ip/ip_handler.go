package ip_service

import (
	"net/http"
	"time"

	"github.com/Uyanide/Api_Collection/internal/logger"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type IPHandler struct {
	ipService *IPService
}

func NewIPHandler(ipService *IPService) *IPHandler {
	return &IPHandler{
		ipService: ipService,
	}
}

func (h *IPHandler) GetIP(c *gin.Context) {
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
	response, err := h.ipService.GetIP(clientIP)
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
