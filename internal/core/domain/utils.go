package domain

import "time"

// NewNowTime returns a now in UTC
func NewNowTime() time.Time {
	return time.Now().UTC()
}

// NewNowTimeString returns a Now in UTC using rfc3339 (iso8601) 2026-02-04THH:MM:ssZ (feb 4th 2026)
func NewNowTimeString() string {
	return time.Now().UTC().Format(time.RFC3339)
}