package config

import (
	"net"
	"os"
	"strconv"
	"strings"

	"github.com/Uyanide/Api_Collection/internal/logger"
	"github.com/Uyanide/Api_Collection/internal/models"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Port int

	LocalIP      string
	LocalCIDRs   []*net.IPNet
	LocalCIDRStr []string

	FileMap map[string]models.FileObject
}

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

	// Parse file map
	fileMap := make(map[string]models.FileObject)
	fileMapEnv := os.Getenv("FILE_MAP")
	if fileMapEnv != "" {
		mappings := strings.Split(fileMapEnv, ",")
		for _, mapping := range mappings {
			parts := strings.SplitN(mapping, ":", 3)
			if len(parts) != 3 {
				log.WithField("mapping", mapping).Warn("Invalid FILE_MAP entry, skipping")
				continue
			}
			urlPath := strings.TrimSpace(parts[0])
			filePath := strings.TrimSpace(parts[1])
			fileName := strings.TrimSpace(parts[2])
			if urlPath == "" || filePath == "" || fileName == "" {
				log.WithField("mapping", mapping).Warn("Empty URL path, file path or file name in FILE_MAP entry, skipping")
				continue
			}
			fileMap[urlPath] = models.FileObject{
				Path: filePath,
				Name: fileName,
			}

			log.WithFields(logrus.Fields{
				"url_path":  urlPath,
				"file_path": filePath,
				"file_name": fileName,
			}).Info("Parsed FILE_MAP entry")
		}
	}

	config := &Config{
		Port:         port,
		LocalIP:      localIP,
		LocalCIDRs:   cidrs,
		LocalCIDRStr: localCIDRs,
		FileMap:      fileMap,
	}

	log.Info("Configuration initialized successfully")

	return config
}

func (c *Config) IsLocalIP(ip net.IP) bool {
	for _, cidr := range c.LocalCIDRs {
		if cidr.Contains(ip) {
			return true
		}
	}
	return false
}
