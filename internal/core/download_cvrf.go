package core

import (
	"context"
	"encoding/json"
	"fmt"
)

// DownloadCVRF downloads the CVRF updates index, and then fetches individual CVRF documents
// if the local cached copy is older/missing
func (e *Engine) DownloadCVRF() error {
	cacheDir, err := e.getCacheDir("cvrf", cvrfVersion)
	if err != nil {
		return err
	}
	e.logger.Debug("Updating CVRF index", "cvrfUpdatesURL", cvrfUpdatesURL, "cacheDir", cacheDir)

	cvrfIndexUpdatesBody, err := e.downloadAndCacheURL(context.Background(), cvrfUpdatesURL, cacheDir, "updates.json")
	if err != nil {
		return err
	}
	defer cvrfIndexUpdatesBody.Close()

	// Parse the JSON response of the updates index into the structs
	var cvrfIndexUpdates CVRFIndexUpdates
	if err := json.NewDecoder(cvrfIndexUpdatesBody).Decode(&cvrfIndexUpdates); err != nil {
		return fmt.Errorf("failed to decode CVRF index updates: %w", err)
	}
	e.logger.Debug("CVRF index updates structs populated", "count", len(cvrfIndexUpdates.Value))

	return nil
}
