package core

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

const (
	// HTTP client config
	httpClientTimeout = 30 * time.Second
	// Base URL for the CVRF API
	cvrfBaseURL    = "https://api.msrc.microsoft.com/cvrf/v3.0/"
	cvrfUpdatesURL = cvrfBaseURL + "updates"
)

func (e *Engine) getCacheDir(cacheType ...string) (string, error) {
	// Try the standard OS cache directory, fall back to temp if unavailable
	baseDir, err := os.UserCacheDir()
	if err != nil {
		e.logger.Warn("Could not find system cache dir, falling back to temp", "err", err)
		baseDir = os.TempDir()
	}

	// Build the base path: .../siredmond/cache
	cachePath := filepath.Join(baseDir, "siredmond", "cache")

	// If a cache type was provided, append it
	// e.g. .../siredmond/cache/cvrf/3.0
	if len(cacheType) > 0 {
		// Go by Example: https://gobyexample.com/range-over-built-in-types
		for _, t := range cacheType {
			cachePath = filepath.Join(cachePath, t)
		}
	}

	err = os.MkdirAll(cachePath, 0755)
	return cachePath, err
}

// fetchURL is a generic helper to perform HTTP GET requests with proper headers
// We pass in context for timeout/cancellation support
func (e *Engine) fetchURL(ctx context.Context, url string) (io.ReadCloser, error) {

	// Set up the HTTP request
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	client := &http.Client{
		Timeout: httpClientTimeout,
	}

	// Perform the HTTP request
	e.logger.Debug("Requesting URL", slog.String("url", url))
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}

	// Check for non-200 response codes
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("server returned error: %s", resp.Status)
	}

	return resp.Body, nil
}

// downloadAndCacheURL downloads data from the given URL and caches it locally
// and returns the file handle to the cached file
func (e *Engine) downloadAndCacheURL(ctx context.Context, url, cacheDir, filename string) (*os.File, error) {

	// Cache file path; fetch resource from URL
	filePath := filepath.Join(cacheDir, filename)
	respBody, err := e.fetchURL(ctx, url)
	if err != nil {
		return nil, err
	}
	defer respBody.Close()

	// Create cache file and write downloaded data
	file, err := os.Create(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to create cache file: %w", err)
	}
	_, err = io.Copy(file, respBody)
	if err != nil {
		file.Close() // Clean up on failure
		return nil, fmt.Errorf("streaming download failed: %w", err)
	}
	// Ensure data is flushed to disk
	file.Sync()

	// Reset file pointer to beginning for caller
	// Could we do something clever with io.TeeReader here instead?
	// Probably, but this is straightforward and works well enough
	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		file.Close() // Clean up on failure
		return nil, fmt.Errorf("failed to reset file pointer: %w", err)
	}

	return file, nil
}

// DownloadCVRF downloads the CVRF updates index, and then fetches individual CVRF documents
// if the local cached copy is older/missing
func (e *Engine) DownloadCVRF() error {
	cacheDir, err := e.getCacheDir("cvrf")
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
