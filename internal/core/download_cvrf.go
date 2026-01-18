package core

import (
	"context"
	"encoding/json"
	"fmt"
)

// DownloadCVRF downloads the CVRF updates index, and then fetches individual CVRF documents
// if the local cached copy is older/missing
func (e *Engine) DownloadCVRF() error {
	cvrfCacheDir, err := e.getCacheDir("cvrf", cvrfVersion)
	if err != nil {
		return err
	}
	cvrfIndex, err := e.DownloadCVRFIndex(cvrfCacheDir)
	if err != nil {
		return err
	}
	e.logger.Debug("CVRF index updates structs populated", "count", len(cvrfIndex.Value))

	return nil
}

func (e *Engine) DownloadCVRFIndex(cvrfCacheDir string) (*CVRFIndexUpdates, error) {
	e.logger.Debug("Updating CVRF index", "cvrfUpdatesURL", cvrfUpdatesURL, "cacheDir", cvrfCacheDir)
	cvrfIndexUpdatesBody, err := e.downloadAndCacheURL(context.Background(), cvrfUpdatesURL, cvrfCacheDir, "updates.json")
	if err != nil {
		return nil, err
	}
	defer cvrfIndexUpdatesBody.Close()

	// Parse the JSON response of the updates index into the structs
	var cvrfIndex CVRFIndexUpdates
	if err := json.NewDecoder(cvrfIndexUpdatesBody).Decode(&cvrfIndex); err != nil {
		return nil, fmt.Errorf("failed to decode CVRF index updates: %w", err)
	}
	return &cvrfIndex, nil
}
