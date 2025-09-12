package proxy_service

import (
	"io"
	"net/http"
	"strings"

	"github.com/Uyanide/Api_Collection/internal/logger"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (s *ProxyService) handleProxy(c *gin.Context) {
	log := logger.GetLogger()

	// Record request anyway
	increaseCounter(ProxiedRequestsKeyPrefix + c.Request.Method)

	targetURL := c.Query("url")
	if targetURL == "" {
		log.Warn("No URL provided in the request")
		c.JSON(http.StatusBadRequest, gin.H{"error": "No URL provided"})
		return
	}

	// Ensure URl is URL
	if !strings.HasPrefix(targetURL, "http://") && !strings.HasPrefix(targetURL, "https://") {
		log.WithField("target_url", targetURL).Warn("Invalid URL scheme")
		if !s.autoCorrectScheme {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid scheme"})
			return
		}
		targetURL = "http://" + targetURL
	}

	// Create a new HTTP request
	req, err := http.NewRequest(c.Request.Method, targetURL, c.Request.Body)
	if err != nil {
		log.WithError(err).Error("Failed to create request")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request"})
		return
	}

	// Copy original request headers
	for key, values := range c.Request.Header {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.WithError(err).Error("Failed to send request")
		c.JSON(http.StatusBadGateway, gin.H{"error": "Failed to send request"})
		return
	}
	defer resp.Body.Close()

	// Copy response headers
	for key, values := range resp.Header {
		for _, value := range values {
			c.Header(key, value)
		}
	}

	// Set status and return body
	c.Status(resp.StatusCode)
	_, err = io.Copy(c.Writer, resp.Body)
	if err != nil {
		log.WithError(err).Error("Failed to copy response body")
	}

	log.WithFields(logrus.Fields{
		"target_url": targetURL,
		"status":     resp.StatusCode,
	}).Info("Proxied request successfully")

	// Record successful request
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		increaseCounter(ProxiedRequestsSuccessfulKey)
	}
}
