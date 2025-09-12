package file_service

import (
	"strings"

	"github.com/Uyanide/Api_Collection/internal/db"
	"github.com/Uyanide/Api_Collection/internal/logger"
)

type StatsFileResponse struct {
	TotalDownloads int64             `json:"total_downloads"`
	MostDownloaded string            `json:"most_downloaded"`
	Files          []StatsFileDetail `json:"files"`
}

type StatsFileDetail struct {
	UrlPath        string `json:"url_path"`
	DownloadsCount int64  `json:"downloads_count"`
}

func ConstructStatsFile() (*StatsFileResponse, error) {
	dbInst := db.GetDB()

	result := StatsFileResponse{}
	result.Files = []StatsFileDetail{}

	for _, key := range FileDownloadsKeys {
		value, err := db.GetOrCreateInt(dbInst, key, 0)
		if err != nil {
			continue // skip error
		}
		urlPath := key[len(FileDownloadsKeyPrefix):]
		result.Files = append(result.Files, StatsFileDetail{
			UrlPath:        urlPath,
			DownloadsCount: value,
		})
	}

	// No files
	if len(result.Files) == 0 {
		result.MostDownloaded = "N/A"
		return &result, nil
	}

	// Find most downloaded & calculate sum
	maxList := []string{}
	maxCount := int64(-1)
	for _, file := range result.Files {
		if file.DownloadsCount > maxCount {
			maxList = []string{file.UrlPath}
			maxCount = file.DownloadsCount
		} else if file.DownloadsCount == maxCount {
			maxList = append(maxList, file.UrlPath)
		}
		result.TotalDownloads += file.DownloadsCount
	}
	result.MostDownloaded = strings.Join(maxList, ", ")

	return &result, nil
}

func increaseCounter(key string) {
	dbInst := db.GetDB()
	log := logger.GetLogger()

	if _, err := db.IncrementInt(dbInst, key, 0, 1); err != nil {
		log.WithError(err).Error("Failed to record file download in database")
	}
}
