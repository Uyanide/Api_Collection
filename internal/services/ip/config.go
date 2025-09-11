package ip_service

import (
	"net"
	"os"
	"strings"

	"github.com/Uyanide/Api_Collection/internal/logger"
	"github.com/sirupsen/logrus"
)

func (s *IPService) loadConfig() {
	log := logger.GetLogger()

	// Parse local IP
	s.localIP = os.Getenv("LOCAL_IP")
	if s.localIP == "" {
		s.localIP = "127.0.0.1"
		log.WithField("local_ip", s.localIP).Warn("No LOCAL_IP environment variable set, using default")
	}
	log.WithField("local_ip", s.localIP).Info("Parsed local IP")

	// Parse CIDR strings
	localCIDRs := strings.Split(os.Getenv("LOCAL_CIDRS"), ",")
	s.localCIDRs = make([]*net.IPNet, len(localCIDRs))
	for i, cidrStr := range localCIDRs {
		_, cidr, err := net.ParseCIDR(cidrStr)
		if err != nil {
			log.WithFields(logrus.Fields{
				"cidr":  cidrStr,
				"error": err.Error(),
			}).Fatal("Invalid CIDR configuration")
		}
		log.WithField("cidr", cidrStr).Info("Parsed local CIDR")
		s.localCIDRs[i] = cidr
	}
}
