package ip_service

import (
	"net"

	"github.com/gin-gonic/gin"
)

var (
	IPRequestsKey = "ip_requests"
)

type IPService struct {
	localIP    string
	localCIDRs []*net.IPNet
}

func (s *IPService) Init(e *gin.Engine) {
	s.loadConfig()
	s.setupRoutes(e)
}
