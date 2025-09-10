package middleware

import (
	"net/http"
	"strings"

	"github.com/Uyanide/Api_Collection/internal/logger"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// StripTrailingSlash removes trailing slashes from request URLs
func StripTrailingSlash(next *gin.Engine) *gin.Engine {
	log := logger.GetLogger()

	next.Use(func(c *gin.Context) {
		if c.Request.URL.Path != "/" && strings.HasSuffix(c.Request.URL.Path, "/") {
			// Create a new URL without the trailing slash
			newPath := strings.TrimSuffix(c.Request.URL.Path, "/")

			log.WithFields(logrus.Fields{
				"original_path": c.Request.URL.Path,
				"new_path":      newPath,
				"client_ip":     c.ClientIP(),
			}).Debug("Redirecting to remove trailing slash")

			// Redirect to the path without trailing slash
			c.Redirect(http.StatusMovedPermanently, newPath)
			c.Abort()
			return
		}

		c.Next()
	})

	return next
}
