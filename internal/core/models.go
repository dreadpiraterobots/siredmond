// I thought about putting this in its own package, but it feels like overkill
// for now. To prove I'm a real software engineer, let me simply say:
// we Can AlWAys COme BaCK anD ImPRovE iT laTeR!
package core

import (
	"sort"
	"time"
)

// CVRFIndex matches the top-level OData structure from the MSRC API.
type CVRFIndex struct {
	Value []CVRFUpdate `json:"value"`
}

// CVRFUpdate represents an individual Microsoft security advisory entry.
// We are intentionally ignoring the 'Severity' field as it is consistently null...
// ...which sorta leaves us wondering why it's even there.
// I think it's probably because different products/remediations might have
// different severities for the same CVE, although I've spent years staring at this stuff
// and I can't remember a specific example. Maybe .NET vs. Windows?
// Maybe once SirEdmond is fully functional, we can use it to dig up a specific example.
// PS sometimes, I worry that my comments look like AI comments...
// ...but it's really me writing this.
type CVRFUpdate struct {
	ID                 string    `json:"ID"`
	Alias              string    `json:"Alias"`
	DocumentTitle      string    `json:"DocumentTitle"`
	InitialReleaseDate time.Time `json:"InitialReleaseDate"`
	CurrentReleaseDate time.Time `json:"CurrentReleaseDate"`
	CvrfUrl            string    `json:"CvrfUrl"`
}

// For downloading individual CVRF docs, it'll be nice to consider them in their original release order.
// Since this is Go, we might need to write some code to do this instead of leaning on a library.
// CVRFUpdates is a slice of CVRFUpdate that we can sort.
type CVRFUpdates []CVRFUpdate

// Len, Less, and Swap implement the sort.Interface for CVRFUpdates
func (s CVRFUpdates) Len() int {
	return len(s)
}

// Less reports whether the element with index i should sort before the element with index j.
func (s CVRFUpdates) Less(i, j int) bool {
	// At least there's a time library!
	return s[i].InitialReleaseDate.Before(s[j].InitialReleaseDate)
}

// Swap swaps the elements with indexes i and j.
func (s CVRFUpdates) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// SortByInitialReleaseDate sorts the CVRFUpdates in ascending order of InitialReleaseDate.
func (s CVRFUpdates) SortByInitialReleaseDate() {
	sort.Sort(s)
}
