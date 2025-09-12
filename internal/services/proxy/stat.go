package proxy_service

import (
	"github.com/Uyanide/Api_Collection/internal/db"
)

type StatsProxyResponse struct {
	TotalRequests int64 `json:"total_requests"`
	Successful    int64 `json:"successful"`
	GET           int64 `json:"get"`
	POST          int64 `json:"post"`
	PUT           int64 `json:"put"`
	DELETE        int64 `json:"delete"`
}

func ConstructStatsProxy() (*StatsProxyResponse, error) {
	dbInst := db.GetDB()

	vGet, err := db.GetOrCreateInt(dbInst, ProxiedRequestsGETKey, 0)
	if err != nil {
		return nil, err
	}
	vPost, err := db.GetOrCreateInt(dbInst, ProxiedRequestsPOSTKey, 0)
	if err != nil {
		return nil, err
	}
	vPut, err := db.GetOrCreateInt(dbInst, ProxiedRequestsPUTKey, 0)
	if err != nil {
		return nil, err
	}
	vDelete, err := db.GetOrCreateInt(dbInst, ProxiedRequestsDELETEKey, 0)
	if err != nil {
		return nil, err
	}
	vSuccessful, err := db.GetOrCreateInt(dbInst, ProxiedRequestsSuccessfulKey, 0)
	if err != nil {
		return nil, err
	}
	return &StatsProxyResponse{
		TotalRequests: vGet + vPost + vPut + vDelete,
		Successful:    vSuccessful,
		GET:           vGet,
		POST:          vPost,
		PUT:           vPut,
		DELETE:        vDelete,
	}, nil
}
