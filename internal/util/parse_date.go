package util

import (
	"fmt"
	"time"
)

// ParseDate parses a string to time.Time. Returns nil if the input is empty.
func ParseDate(dateStr string, format string) (*time.Time, error) {
	// Jika string kosong, kembalikan nil tanpa melakukan parsing
	if dateStr == "" {
		return nil, nil
	}

	// Parsing string ke time.Time
	parsedDate, err := time.Parse(format, dateStr)
	if err != nil {
		return nil, fmt.Errorf("invalid date format: %v", err)
	}

	// Mengembalikan pointer ke parsedDate
	return &parsedDate, nil
}
