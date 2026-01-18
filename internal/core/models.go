// I thought about putting this in its own package, but it feels like overkill
// for now. To prove I'm a real software engineer, let me simply say:
// we Can AlWAys COme BaCK anD ImPRovE iT laTeR!
package core

import "time"

// CVRFIndexUpdates matches the top-level OData structure from the MSRC API.
type CVRFIndexUpdates struct {
	Value []CVRFIndexUpdateObject `json:"value"`
}

// CVRFIndexUpdateObject represents an individual Microsoft security advisory entry.
// We are intentionally ignoring the 'Severity' field as it is consistently null...
// ...which sorta leaves us wondering why it's even there.
// I think it's probably because different products/remediations might have
// different severities for the same CVE, although I've spent years staring at this stuff
// and I can't remember a specific example. Maybe .NET vs. Windows?
// Maybe once SirEdmond is fully functional, we can use it to dig up a specific example.
// PS sometimes, I worry that my comments look like AI comments...
// ...but it's really me writing this.
type CVRFIndexUpdateObject struct {
	ID                 string    `json:"ID"`
	Alias              string    `json:"Alias"`
	DocumentTitle      string    `json:"DocumentTitle"`
	InitialReleaseDate time.Time `json:"InitialReleaseDate"`
	CurrentReleaseDate time.Time `json:"CurrentReleaseDate"`
	CvrfUrl            string    `json:"CvrfUrl"`
}
