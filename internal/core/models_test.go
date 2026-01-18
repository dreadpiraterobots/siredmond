package core

import (
	"encoding/json"
	"testing"
	"time"
)

// TestCVRFIndexMapping ensures that the MSRC OData JSON response correctly
// populates our CVRFIndexUpdates and CVRFIndexUpdateObject structs.
// This guards against breaking changes in the MSRC API or our own model definitions.
func TestCVRFIndexMapping(t *testing.T) {
	data := `{
		"value": [
			{
				"ID": "2025-Dec",
				"Alias": "2025-Dec",
				"DocumentTitle": "Test not-actually-real CVRF",
				"CurrentReleaseDate": "2025-12-28T23:00:00Z",
				"CvrfUrl": "https://api.msrc.microsoft.com/cvrf/v3.0/cvrf/2025-Dec"
			}
		]
	}`

	var updates CVRFIndexUpdates
	err := json.Unmarshal([]byte(data), &updates)

	// Assertions to ensure the data integrity of our models.
	if err != nil {
		t.Fatalf("Failed to map MSRC JSON to CVRFIndexUpdates: %v", err)
	}

	// Verify the slice contains exactly one object cos that's what the test data has.
	if len(updates.Value) != 1 {
		t.Errorf("Index slice count mismatch: expected 1, got %d", len(updates.Value))
	}

	// Verify the ID, which is used everywhere and thus really important.
	gotID := updates.Value[0].ID
	if gotID != "2025-Dec" {
		t.Errorf("Index object ID mismatch: expected '2025-Dec', got '%s'", gotID)
	}

	// Verify the URL for this month's CVRF, which is the source for the subsequent XML downloads.
	gotURL := updates.Value[0].CvrfUrl
	expectedURL := "https://api.msrc.microsoft.com/cvrf/v3.0/cvrf/2025-Dec"

	// See, this is what I love about unit tests: definitely not self-referential!
	if gotURL != expectedURL {
		t.Errorf("CvrfUrl mapping failed: expected %s, got %s", expectedURL, gotURL)
	}

	// Verify the timestamp was converted to a Go time.Time object correctly.
	expectedTime, _ := time.Parse(time.RFC3339, "2025-12-28T23:00:00Z")
	if !updates.Value[0].CurrentReleaseDate.Equal(expectedTime) {
		t.Errorf("CurrentReleaseDate mismatch: expected %v, got %v", expectedTime, updates.Value[0].CurrentReleaseDate)
	}
}
