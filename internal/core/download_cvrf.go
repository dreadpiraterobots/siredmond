package core

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

// DownloadCVRF downloads the CVRF updates index, and then fetches individual CVRF documents
// if the local cached copy is older/missing
func (e *Engine) DownloadCVRF(ctx context.Context) error {
	cvrfCacheDir, err := e.getCacheDir("cvrf", cvrfVersion)
	if err != nil {
		return err
	}
	cvrfUpdates, err := e.DownloadCVRFIndex(ctx, cvrfCacheDir)
	if err != nil {
		return err
	}
	e.CVRFIndexStats(cvrfUpdates)
	return nil
}

func (e *Engine) DownloadCVRFIndex(ctx context.Context, cvrfCacheDir string) (*CVRFUpdates, error) {
	e.logger.Debug("Updating CVRF index", "cvrfUpdatesURL", cvrfUpdatesURL, "cacheDir", cvrfCacheDir)
	cvrfIndexUpdatesBody, err := e.downloadAndCacheURL(ctx, cvrfUpdatesURL, cvrfCacheDir, "updates.json")
	if err != nil {
		return nil, err
	}
	defer cvrfIndexUpdatesBody.Close()

	// Parse the JSON response of the updates index into the structs
	var cvrfIndex CVRFIndex
	if err := json.NewDecoder(cvrfIndexUpdatesBody).Decode(&cvrfIndex); err != nil {
		return nil, fmt.Errorf("failed to decode CVRF index updates: %w", err)
	}
	cvrfUpdates := CVRFUpdates(cvrfIndex.Value)
	cvrfUpdates.SortByInitialReleaseDate()

	return &cvrfUpdates, nil
}

// For fun, let's do some basic stats on the CVRF index info in the structs
// We don't strictly need this to be a pointer receiver, since we're not modifying anything
// However, passing pointers avoids copying stuff in memory unnecessarily
// so it's a good habit to get into because structs can get large
// https://go.dev/tour/methods/4
func (e *Engine) CVRFIndexStats(cvrfUpdates *CVRFUpdates) {
	e.logger.Debug("CVRF index updates structs populated", "count", cvrfUpdates.Len())
	var oldestUpdate, newestUpdate time.Time
	var oldestUpdateId, newestUpdateId string
	for i, item := range *cvrfUpdates {
		itemCurrentReleaseTime := item.CurrentReleaseDate
		if i == 0 || itemCurrentReleaseTime.Before(oldestUpdate) {
			oldestUpdate = itemCurrentReleaseTime
			oldestUpdateId = item.ID
		}
		if i == 0 || itemCurrentReleaseTime.After(newestUpdate) {
			newestUpdate = itemCurrentReleaseTime
			newestUpdateId = item.ID
		}
	}
	e.logger.Debug("CVRF index stats", "oldestUpdateId", oldestUpdateId, "oldestUpdate", oldestUpdate, "newestUpdateId", newestUpdateId, "newestUpdate", newestUpdate)
}
