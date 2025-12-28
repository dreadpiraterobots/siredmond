package core

import (
	"fmt"
	"os"
	"path/filepath"
)

func (e *Engine) getCacheDir(cacheType ...string) (string, error) {
	// Try the standard OS cache directory, fall back to temp if unavailable
	baseDir, err := os.UserCacheDir()
	if err != nil {
		e.logger.Warn("could not find system cache dir, falling back to temp", "err", err)
		baseDir = os.TempDir()
	}

	// Build the base path: .../siredmond/cache
	cachePath := filepath.Join(baseDir, "siredmond", "cache")

	// If a type was provided, append it: .../siredmond/cache/cvrf
	if len(cacheType) > 0 && cacheType[0] != "" {
		cachePath = filepath.Join(cachePath, cacheType[0])
	}

	err = os.MkdirAll(cachePath, 0755)
	return cachePath, err
}

func (e *Engine) DownloadCVRF() error {
	e.logger.Info("Updating CVRF index")
	cacheDir, err := e.getCacheDir("cvrf")
	if err != nil {
		return err
	}
	e.logger.Debug(fmt.Sprintf("using cache directory: %s", cacheDir))
	return nil
}
