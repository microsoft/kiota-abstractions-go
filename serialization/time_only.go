package serialization

import (
	"fmt"
	"strings"
	"time"
)

// TimeOnly is represents the time part of a date time (time) value.
type TimeOnly struct {
	time      time.Time
	precision int // number of decimal places for nanoseconds
}

const timeOnlyFormat = "15:04:05.000000000"

var timeOnlyParsingFormats = map[int]string{
	0: "15:04:05", //Go doesn't seem to support optional parameters in time.Parse, which is sad
	1: "15:04:05.0",
	2: "15:04:05.00",
	3: "15:04:05.000",
	4: "15:04:05.0000",
	5: "15:04:05.00000",
	6: "15:04:05.000000",
	7: "15:04:05.0000000",
	8: "15:04:05.00000000",
	9: timeOnlyFormat,
}

// String returns the time only as a string following the RFC3339 standard.
func (t TimeOnly) String() string {
	if t.precision == 0 {
		return t.time.Format("15:04:05")
	}
	return t.time.Format(timeOnlyParsingFormats[t.precision])
}

// ParseTimeOnly parses a string into a TimeOnly following the RFC3339 standard.
func ParseTimeOnly(s string) (*TimeOnly, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return nil, nil
	}

	precision := 0
	if parts := strings.Split(s, "."); len(parts) > 1 {
		precision = len(parts[1])
		if precision >= len(timeOnlyParsingFormats) {
			return nil, fmt.Errorf("time precision of %d exceeds maximum allowed of %d", precision, len(timeOnlyParsingFormats)-1)
		}
	}

	timeValue, err := time.Parse(timeOnlyParsingFormats[precision], s)
	if err != nil {
		return nil, err
	}

	return &TimeOnly{time: timeValue, precision: precision}, nil
}

// NewTimeOnly creates a new TimeOnly from a time.Time.
func NewTimeOnly(t time.Time) *TimeOnly {
	precision := 0
	nanos := t.Nanosecond()
	if nanos > 0 {
		precision = len(strings.TrimRight(fmt.Sprintf("%09d", nanos), "0"))
	}
	return &TimeOnly{
		time:      t,
		precision: precision,
	}
}
