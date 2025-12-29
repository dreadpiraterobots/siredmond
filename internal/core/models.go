package core

import "time"

// CVRFIndexUpdates matches the top-level OData structure from the MSRC API.
type CVRFIndexUpdates struct {
	Value []CVRFIndexUpdateObject `json:"value"`
}

// CVRFIndexUpdateObject represents an individual Microsoft security advisory entry.
// We are intentionally ignoring the 'Severity' field as it is consistently null.
type CVRFIndexUpdateObject struct {
	ID                 string    `json:"ID"`
	Alias              string    `json:"Alias"`
	DocumentTitle      string    `json:"DocumentTitle"`
	InitialReleaseDate time.Time `json:"InitialReleaseDate"`
	CurrentReleaseDate time.Time `json:"CurrentReleaseDate"`
	CvrfUrl            string    `json:"CvrfUrl"`
}
