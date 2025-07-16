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
		wantNil     bool
		precision   int
		errorString string
	}{
		{"basic time", "15:04:05", "15:04:05", false, false, 0, ""},
		{"one decimal", "15:04:05.1", "15:04:05.1", false, false, 1, ""},
		{"three decimals", "15:04:05.123", "15:04:05.123", false, false, 3, ""},
		{"max decimals", "15:04:05.123456789", "15:04:05.123456789", false, false, 9, ""},
		{"empty string", "", "", false, true, 0, ""},
		{"whitespace", "  ", "", false, true, 0, ""},
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

			if tt.precision == 0 {
				if got.String() != tt.wantTime {
					t.Errorf("ParseTimeOnly(%q) = %v, want %v", tt.input, got.String(), tt.wantTime)
				}
			} else {
				if got.StringWithPrecision(tt.precision) != tt.wantTime {
					t.Errorf("ParseTimeOnly(%q) StringWithPrecision(%d) = %v, want %v", tt.input, tt.precision, got.StringWithPrecision(tt.precision), tt.wantTime)
				}
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
			if tt.precision == 0 {
				if got.String() != tt.wantTime {
					t.Errorf("NewTimeOnly() String() = %v, want %v", got.String(), tt.wantTime)
				}
			} else {
				if got.StringWithPrecision(tt.precision) != tt.wantTime {
					t.Errorf("NewTimeOnly() StringWithPrecision(%d) = %v, want %v", tt.precision, got.StringWithPrecision(tt.precision), tt.wantTime)
				}
			}
			detectedPrecision := DetectPrecision(tt.input)
			if detectedPrecision != tt.precision {
				t.Errorf("DetectPrecision() = %v, want %v", detectedPrecision, tt.precision)
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
			timeOnly := TimeOnly{time: tt.time}
			if got := timeOnly.StringWithPrecision(tt.precision); got != tt.want {
				t.Errorf("TimeOnly.StringWithPrecision(%d) = %v, want %v", tt.precision, got, tt.want)
			}
		})
	}
}

func TestParseTimeOnlyWithPrecision(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		wantTime      string
		wantPrecision int
		wantError     bool
		wantNil       bool
		errorString   string
	}{
		{"basic time", "15:04:05", "15:04:05", 0, false, false, ""},
		{"one decimal", "15:04:05.1", "15:04:05.1", 1, false, false, ""},
		{"three decimals", "15:04:05.123", "15:04:05.123", 3, false, false, ""},
		{"max decimals", "15:04:05.123456789", "15:04:05.123456789", 9, false, false, ""},
		{"empty string", "", "", 0, false, true, ""},
		{"whitespace", "  ", "", 0, false, true, ""},
		{"too many decimals", "15:04:05.1234567890", "", 0, true, false, "time precision of 10 exceeds maximum allowed of 9"},
		{"invalid time", "25:04:05", "", 0, true, false, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, precision, err := ParseTimeOnlyWithPrecision(tt.input)

			if tt.wantNil {
				if got != nil {
					t.Errorf("ParseTimeOnlyWithPrecision(%q) = %v, want nil", tt.input, got)
				}
				return
			}

			if tt.wantError {
				if err == nil {
					t.Errorf("ParseTimeOnlyWithPrecision(%q) expected error, got nil", tt.input)
				}
				if tt.errorString != "" && err.Error() != tt.errorString {
					t.Errorf("ParseTimeOnlyWithPrecision(%q) error = %v, want %v", tt.input, err, tt.errorString)
				}
				return
			}

			if err != nil {
				t.Errorf("ParseTimeOnlyWithPrecision(%q) unexpected error: %v", tt.input, err)
				return
			}

			if got == nil {
				t.Errorf("ParseTimeOnlyWithPrecision(%q) returned nil, want non-nil", tt.input)
				return
			}

			if got.StringWithPrecision(precision) != tt.wantTime {
				t.Errorf("ParseTimeOnlyWithPrecision(%q) StringWithPrecision(%d) = %v, want %v", tt.input, precision, got.StringWithPrecision(precision), tt.wantTime)
			}

			if precision != tt.wantPrecision {
				t.Errorf("ParseTimeOnlyWithPrecision(%q) precision = %v, want %v", tt.input, precision, tt.wantPrecision)
			}
		})
	}
}

func TestDetectPrecision(t *testing.T) {
	tests := []struct {
		name  string
		input time.Time
		want  int
	}{
		{"no decimals", time.Date(2023, 1, 1, 15, 4, 5, 0, time.UTC), 0},
		{"one decimal", time.Date(2023, 1, 1, 15, 4, 5, 100000000, time.UTC), 1},
		{"three decimals", time.Date(2023, 1, 1, 15, 4, 5, 123000000, time.UTC), 3},
		{"six decimals", time.Date(2023, 1, 1, 15, 4, 5, 123456000, time.UTC), 6},
		{"nine decimals", time.Date(2023, 1, 1, 15, 4, 5, 123456789, time.UTC), 9},
		{"trailing zeros", time.Date(2023, 1, 1, 15, 4, 5, 120000000, time.UTC), 2},
		{"all zeros", time.Date(2023, 1, 1, 15, 4, 5, 100000000, time.UTC), 1},
		{"midnight", time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC), 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DetectPrecision(tt.input)
			if got != tt.want {
				t.Errorf("DetectPrecision() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTimeOnly_String_DefaultPrecision(t *testing.T) {
	tests := []struct {
		name string
		time time.Time
		want string
	}{
		{"with nanoseconds", time.Date(2023, 1, 1, 15, 4, 5, 123456789, time.UTC), "15:04:05"},
		{"without nanoseconds", time.Date(2023, 1, 1, 15, 4, 5, 0, time.UTC), "15:04:05"},
		{"midnight", time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC), "00:00:00"},
		{"midnight with nanoseconds", time.Date(2023, 1, 1, 0, 0, 0, 123456789, time.UTC), "00:00:00"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			timeOnly := TimeOnly{time: tt.time}
			if got := timeOnly.String(); got != tt.want {
				t.Errorf("TimeOnly.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTimeOnly_StringRoundTrip(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"basic time", "15:04:05", "15:04:05"},
		{"time with decimals", "15:04:05.123", "15:04:05"},
		{"leading zeros", "05:04:05", "05:04:05"},
		{"leading zeros with decimals", "05:04:05.123456", "05:04:05"},
		{"midnight", "00:00:00", "00:00:00"},
		{"midnight with decimals", "00:00:00.999", "00:00:00"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parsed, err := ParseTimeOnly(tt.input)
			if err != nil {
				t.Errorf("ParseTimeOnly(%q) unexpected error: %v", tt.input, err)
				return
			}
			if got := parsed.String(); got != tt.want {
				t.Errorf("ParseTimeOnly(%q).String() = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestTimeOnly_StringWithPrecision_EdgeCases(t *testing.T) {
	timeOnly := TimeOnly{time: time.Date(2023, 1, 1, 15, 4, 5, 123456789, time.UTC)}

	tests := []struct {
		name      string
		precision int
		want      string
	}{
		{"negative precision", -1, "15:04:05"},
		{"zero precision", 0, "15:04:05"},
		{"valid precision", 3, "15:04:05.123"},
		{"max precision", 9, "15:04:05.123456789"},
		{"out of range precision", 10, "15:04:05"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := timeOnly.StringWithPrecision(tt.precision); got != tt.want {
				t.Errorf("StringWithPrecision(%d) = %v, want %v", tt.precision, got, tt.want)
			}
		})
	}
}
