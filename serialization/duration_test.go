package serialization

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestItNormalizesMS(t *testing.T) {
	// Arrange
	duration := &duration{
		MilliSeconds: 1001,
	}

	// Act
	result := duration.String()

	// Assert
	assert.Equal(t, 1, duration.Seconds)
	assert.Equal(t, 1, duration.MilliSeconds)
	assert.Equal(t, "PT1S", result)
}

func TestItNormalizesS(t *testing.T) {
	// Arrange
	duration := &duration{
		Seconds: 61,
	}

	// Act
	result := duration.String()

	// Assert
	assert.Equal(t, 1, duration.Seconds)
	assert.Equal(t, 1, duration.Minutes)
	assert.Equal(t, "PT1M1S", result)
}

func TestToDurationWithMonths(t *testing.T) {
	// Arrange
	duration := &duration{
		Years:   1,
		Months:  2,
		Weeks:   1,
		Days:    3,
		Hours:   4,
		Minutes: 5,
		Seconds: 6,
	}

	// Act
	result, err := duration.ToDurationWithMonths(30)

	// Assert
	assert.Nil(t, err)
	expected := time.duration(
		(1 * 365 * 24 * time.Hour) + // 1 year
			(2 * 30 * 24 * time.Hour) + // 2 months (30 days each)
			(7 * 24 * time.Hour) + // 1 week
			(3 * 24 * time.Hour) + // 3 days
			(4 * time.Hour) + // 4 hours
			(5 * time.Minute) + // 5 minutes
			(6 * time.Second), // 6 seconds
	)
	assert.Equal(t, expected, result)
}

func TestFromString(t *testing.T) {
}

func TestItNormalizesMi(t *testing.T) {
	// Arrange
	duration := &duration{
		Minutes: 61,
	}

	// Act
	result := duration.String()

	// Assert
	assert.Equal(t, 1, duration.Hours)
	assert.Equal(t, 1, duration.Minutes)
	assert.Equal(t, "PT1H1M", result)
}

func TestItNormalizesH(t *testing.T) {
	// Arrange
	duration := &duration{
		Hours: 25,
	}

	// Act
	result := duration.String()

	// Assert
	assert.Equal(t, 1, duration.Hours)
	assert.Equal(t, 1, duration.Days)
	assert.Equal(t, "P1DT1H", result)
}

func TestItNormalizesD(t *testing.T) {
	// Arrange
	duration := &duration{
		Days: 8,
	}

	// Act
	result := duration.String()

	// Assert
	assert.Equal(t, 1, duration.Weeks)
	assert.Equal(t, 1, duration.Days)
	assert.Equal(t, "P1W1D", result)
}

func TestItDoesntNormalizesW(t *testing.T) {
	// Arrange
	duration := &duration{
		Weeks: 56,
	}

	// Act
	result := duration.String()

	// Assert
	assert.Equal(t, 56, duration.Weeks)
	assert.Equal(t, 0, duration.Years)
	assert.Equal(t, "P56W", result)
}

func TestItDoesntNormalizesDWithMo(t *testing.T) {
	// Arrange
	duration := &duration{
		Days:   15,
		Months: 2,
	}

	// Act
	result := duration.String()

	// Assert
	assert.Equal(t, 15, duration.Days)
	assert.Equal(t, 0, duration.Weeks)
	assert.Equal(t, 2, duration.Months)
	assert.Equal(t, "P2M15D", result)
}

func TestItNormalizesMo(t *testing.T) {
	// Arrange
	duration := &duration{
		Months: 13,
	}

	// Act
	result := duration.String()

	// Assert
	assert.Equal(t, 1, duration.Months)
	assert.Equal(t, 1, duration.Years)
	assert.Equal(t, "P1Y1M", result)
}

func TestItRefusesMoAndW(t *testing.T) {
	// Arrange
	duration := &duration{
		Months: 13,
		Weeks:  10,
	}

	// Act
	result := duration.Normalize()

	// Assert
	assert.Equal(t, errWeeksNotWithYearsOrMonth, result)
}

func TestItRefusesYAndW(t *testing.T) {
	// Arrange
	duration := &duration{
		Years: 13,
		Weeks: 10,
	}

	// Act
	result := duration.Normalize()

	// Assert
	assert.Equal(t, errWeeksNotWithYearsOrMonth, result)
}

func TestItFailsMoToDuration(t *testing.T) {
	// Arrange
	duration := &duration{
		Months: 13,
		Weeks:  10,
	}

	// Act
	result, err := duration.ToDuration()

	// Assert
	assert.Equal(t, time.duration(0), result)
	assert.Equal(t, errMonthsInDurationUseOverload, err)
}

func TestItParsesMonth(t *testing.T) {
	// Act
	duration, err := FromString("P1M")

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, 1, duration.Months)
}

func TestItParsesMonthAndMinutes(t *testing.T) {
	// Act
	duration, err := FromString("P1MT1M")

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, 1, duration.Months)
	assert.Equal(t, 1, duration.Minutes)
}
