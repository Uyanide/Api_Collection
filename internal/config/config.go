package config

import (
	"net"
	"os"
	"strconv"
	"strings"

	"github.com/Uyanide/Api_Collection/internal/logger"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

// Config holds the application configuration
type Config struct {
	Port         int
	LocalIP      string
	LocalCIDRs   []*net.IPNet
	LocalCIDRStr []string
}

// NewConfig creates a new configuration with default values
func NewConfig() *Config {

	log := logger.GetLogger()
	log.Info("Initializing configuration")

	if err := godotenv.Load(); err != nil {
		log.Warn("No .env file found")
	}

	// Parse CIDR strings
	localCIDRs := strings.Split(os.Getenv("LOCAL_CIDRS"), ",")
	cidrs := make([]*net.IPNet, len(localCIDRs))
	for i, cidrStr := range localCIDRs {
		_, cidr, err := net.ParseCIDR(cidrStr)
		if err != nil {
			log.WithFields(logrus.Fields{
				"cidr":  cidrStr,
				"error": err.Error(),
			}).Fatal("Invalid CIDR configuration")
		}
		log.WithField("cidr", cidrStr).Info("Parsed local CIDR")
		cidrs[i] = cidr
	}

	// Parse port
	portString := os.Getenv("PORT")
	if portString == "" {
		portString = "10087"
		log.WithField("port", portString).Warn("No PORT environment variable set, using default")
	}
	port, err := strconv.Atoi(portString)
	if err != nil || port <= 0 || port > 65535 {
		log.WithFields(logrus.Fields{
			"port":  portString,
			"error": err.Error(),
		}).Fatal("Invalid port configuration")
	}

	// Parse local IP
	localIP := os.Getenv("LOCAL_IP")
	if localIP == "" {
		localIP = "127.0.0.1"
		log.WithField("local_ip", localIP).Warn("No LOCAL_IP environment variable set, using default")
	}
	log.WithField("local_ip", localIP).Info("Parsed local IP")

	config := &Config{
		Port:         port,
		LocalIP:      localIP,
		LocalCIDRs:   cidrs,
		LocalCIDRStr: localCIDRs,
	}

	log.WithFields(logrus.Fields{
		"port":        config.Port,
		"local_ip":    config.LocalIP,
		"local_cidrs": config.LocalCIDRStr,
	}).Info("Configuration initialized successfully")

	return config
}

// IsLocalIP checks if the given IP belongs to any of the local CIDRs
func (c *Config) IsLocalIP(ip net.IP) bool {
	for _, cidr := range c.LocalCIDRs {
		if cidr.Contains(ip) {
			return true
		}
	}
	return false
}
