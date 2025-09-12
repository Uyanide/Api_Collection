package ip_service

import (
	"github.com/Uyanide/Api_Collection/internal/db"
)

type StatsIPResponse struct {
	TotalRequests int64 `json:"total_requests"`
}

func ConstructStatsIP() (*StatsIPResponse, error) {
	dbInst := db.GetDB()

	value, err := db.GetOrCreateInt(dbInst, IPRequestsKey, 0)
	if err != nil {
		return nil, err
	}
	return &StatsIPResponse{
		TotalRequests: value,
	}, nil
}
