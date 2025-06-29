package internal

import (
	"fmt"
	"strconv"
	"time"
)

func parseDate(dateStr string) (*time.Time, error) {
	if dateStr == "" {
		return nil, nil
	}

	layout := "02/01/2006"
	t, err := time.Parse(layout, dateStr)
	if err != nil {
		return nil, fmt.Errorf("failed to parse date %q: %w", dateStr, err)
	}
	return &t, nil
}

func parseInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return n
}
