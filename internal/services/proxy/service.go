package proxy_service

import "github.com/gin-gonic/gin"

var (
	ProxiedRequestsKeyPrefix     = "proxied_requests_"
	ProxiedRequestsGETKey        = ProxiedRequestsKeyPrefix + "GET"
	ProxiedRequestsPOSTKey       = ProxiedRequestsKeyPrefix + "POST"
	ProxiedRequestsPUTKey        = ProxiedRequestsKeyPrefix + "PUT"
	ProxiedRequestsDELETEKey     = ProxiedRequestsKeyPrefix + "DELETE"
	ProxiedRequestsSuccessfulKey = ProxiedRequestsKeyPrefix + "successful"
)

type ProxyService struct {
	autoCorrectScheme bool
}

func (s *ProxyService) Init(e *gin.Engine) {
	s.loadConfig()
	s.setupRoutes(e)
}
