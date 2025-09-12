package ip_service

import (
	"github.com/Uyanide/Api_Collection/internal/db"
	"github.com/Uyanide/Api_Collection/internal/logger"
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

func increaseCounter() {
	dbInst := db.GetDB()
	log := logger.GetLogger()
	if _, err := db.IncrementInt(dbInst, IPRequestsKey, 0, 1); err != nil {
		log.WithError(err).Error("Failed to record ip request in database")
	}
}
