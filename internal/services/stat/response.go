package stats_service

import (
	file_service "github.com/Uyanide/Api_Collection/internal/services/file"
	ip_service "github.com/Uyanide/Api_Collection/internal/services/ip"
	proxy_service "github.com/Uyanide/Api_Collection/internal/services/proxy"
)

type StatsGeneralResponse struct {
	TotalRequests   int64                            `json:"total_requests"`
	IPRequests      ip_service.StatsIPResponse       `json:"ip_requests"`
	FileDownloads   file_service.StatsFileResponse   `json:"file_downloads"`
	ProxiedRequests proxy_service.StatsProxyResponse `json:"proxied_requests"`
}
