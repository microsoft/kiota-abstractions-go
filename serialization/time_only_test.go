package serialization

import (
	"testing"
	"time"
)

func TestParseTimeOnly(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		wantTime    string
		wantError   bool
		wantNil     bool // add this field to explicitly check for nil returns
		precision   int
		errorString string
	}{
		{"basic time", "15:04:05", "15:04:05", false, false, 0, ""},
		{"one decimal", "15:04:05.1", "15:04:05.1", false, false, 1, ""},
		{"three decimals", "15:04:05.123", "15:04:05.123", false, false, 3, ""},
		{"max decimals", "15:04:05.123456789", "15:04:05.123456789", false, false, 9, ""},
		{"empty string", "", "", false, true, 0, ""}, // update to expect nil
		{"whitespace", "  ", "", false, true, 0, ""}, // update to expect nil
		{"too many decimals", "15:04:05.1234567890", "", true, false, 0, "time precision of 10 exceeds maximum allowed of 9"},
		{"invalid time", "25:04:05", "", true, false, 0, ""},
		{"leading zeros", "05:04:05", "05:04:05", false, false, 0, ""},
		{"leading zeros with precision", "05:04:05.123", "05:04:05.123", false, false, 3, ""},
		{"invalid format", "5:4:5", "", true, false, 0, ""},
		{"invalid minutes", "05:60:05", "", true, false, 0, ""},
		{"invalid seconds", "05:04:60", "", true, false, 0, ""},
		{"missing seconds", "05:04", "", true, false, 0, ""},
		{"extra components", "05:04:05:01", "", true, false, 0, ""},
		{"trailing dot", "05:04:05.", "", true, false, 0, ""},
		{"only dot decimal", "05:04:05.", "", true, false, 0, ""},
		{"non-numeric decimal", "05:04:05.abc", "", true, false, 0, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseTimeOnly(tt.input)

			if tt.wantNil {
				if got != nil {
					t.Errorf("ParseTimeOnly(%q) = %v, want nil", tt.input, got)
				}
				return
			}

			if tt.wantError {
				if err == nil {
					t.Errorf("ParseTimeOnly(%q) expected error, got nil", tt.input)
				}
				if tt.errorString != "" && err.Error() != tt.errorString {
					t.Errorf("ParseTimeOnly(%q) error = %v, want %v", tt.input, err, tt.errorString)
				}
				return
			}

			if err != nil {
				t.Errorf("ParseTimeOnly(%q) unexpected error: %v", tt.input, err)
				return
			}

			if got == nil {
				t.Errorf("ParseTimeOnly(%q) returned nil, want non-nil", tt.input)
				return
			}

			if got.String() != tt.wantTime {
				t.Errorf("ParseTimeOnly(%q) = %v, want %v", tt.input, got.String(), tt.wantTime)
			}

			if got.precision != tt.precision {
				t.Errorf("ParseTimeOnly(%q) precision = %v, want %v", tt.input, got.precision, tt.precision)
			}
		})
	}
}

func TestNewTimeOnly(t *testing.T) {
	tests := []struct {
		name      string
		input     time.Time
		wantTime  string
		precision int
	}{
		{"no decimals", time.Date(2023, 1, 1, 15, 4, 5, 0, time.UTC), "15:04:05", 0},
		{"three decimals", time.Date(2023, 1, 1, 15, 4, 5, 123000000, time.UTC), "15:04:05.123", 3},
		{"six decimals", time.Date(2023, 1, 1, 15, 4, 5, 123456000, time.UTC), "15:04:05.123456", 6},
		{"nine decimals", time.Date(2023, 1, 1, 15, 4, 5, 123456789, time.UTC), "15:04:05.123456789", 9},
		{"trailing zeros", time.Date(2023, 1, 1, 15, 4, 5, 120000000, time.UTC), "15:04:05.12", 2},
		{"all zeros", time.Date(2023, 1, 1, 15, 4, 5, 100000000, time.UTC), "15:04:05.1", 1},
		{"zero hour", time.Date(2023, 1, 1, 0, 4, 5, 123000000, time.UTC), "00:04:05.123", 3},
		{"midnight", time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC), "00:00:00", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewTimeOnly(tt.input)
			if got.String() != tt.wantTime {
				t.Errorf("NewTimeOnly() = %v, want %v", got.String(), tt.wantTime)
			}
			if got.precision != tt.precision {
				t.Errorf("NewTimeOnly() precision = %v, want %v", got.precision, tt.precision)
			}
		})
	}
}

func TestTimeOnly_String(t *testing.T) {
	tests := []struct {
		name      string
		time      time.Time
		precision int
		want      string
	}{
		{"zero precision", time.Date(2023, 1, 1, 15, 4, 5, 123456789, time.UTC), 0, "15:04:05"},
		{"precision 3", time.Date(2023, 1, 1, 15, 4, 5, 123456789, time.UTC), 3, "15:04:05.123"},
		{"precision 6", time.Date(2023, 1, 1, 15, 4, 5, 123456789, time.UTC), 6, "15:04:05.123456"},
		{"precision 9", time.Date(2023, 1, 1, 15, 4, 5, 123456789, time.UTC), 9, "15:04:05.123456789"},
		{"midnight zero precision", time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC), 0, "00:00:00"},
		{"midnight with precision", time.Date(2023, 1, 1, 0, 0, 0, 100000000, time.UTC), 1, "00:00:00.1"},
		{"leading zero minutes", time.Date(2023, 1, 1, 15, 4, 5, 0, time.UTC), 0, "15:04:05"},
		{"leading zero everything", time.Date(2023, 1, 1, 5, 4, 5, 0, time.UTC), 0, "05:04:05"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			timeOnly := TimeOnly{
				time:      tt.time,
				precision: tt.precision,
			}
			if got := timeOnly.String(); got != tt.want {
				t.Errorf("TimeOnly.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
