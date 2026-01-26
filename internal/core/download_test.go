package core

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestDownloadAndCacheURLWritesFileAndReturnsHandle(t *testing.T) {
	// We use httptest to create a test server that returns a known response
	// and avoids network dependencies in the test.

	httpTestServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"ok"}`))
	}))
	defer httpTestServer.Close()

	engine := NewEngine()
	cacheDir := t.TempDir()

	// Write an actual file returned by the test server
	f, err := engine.downloadAndCacheURL(context.Background(), httpTestServer.URL, cacheDir, "test.json")
	if err != nil {
		t.Fatalf("downloadAndCacheURL failed: %v", err)
	}
	defer f.Close()

	// Ensure the returned file is readable from the start
	b, err := io.ReadAll(f)
	if err != nil {
		t.Fatalf("reading returned file failed: %v", err)
	}
	if string(b) != `{"message":"ok"}` {
		t.Fatalf("unexpected file contents: %s", string(b))
	}

	// Also check that the file exists on disk
	info, err := os.Stat(f.Name())
	if err != nil {
		t.Fatalf("stat cached file failed: %v", err)
	}
	if info.Size() != int64(len(b)) {
		t.Fatalf("cached file size mismatch: expected %d got %d", len(b), info.Size())
	}
}

func TestFetchURLReturnsErrorOnNon200(t *testing.T) {
	httpTestServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "BSOD!", http.StatusInternalServerError)
	}))
	defer httpTestServer.Close()

	engine := NewEngine()
	ctx := context.Background()
	_, err := engine.fetchURL(ctx, httpTestServer.URL)
	if err == nil {
		t.Fatalf("expected fetchURL to return error for non-200 response")
	}
}
