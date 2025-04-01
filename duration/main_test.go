package duration

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestItNormalizesMS(t *testing.T) {
	// Arrange
	duration := &Duration{
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
	duration := &Duration{
		Seconds: 61,
	}

	// Act
	result := duration.String()

	// Assert
	assert.Equal(t, 1, duration.Seconds)
	assert.Equal(t, 1, duration.Minutes)
	assert.Equal(t, "PT1M1S", result)
}

func TestItNormalizesMi(t *testing.T) {
	// Arrange
	duration := &Duration{
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
	duration := &Duration{
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
	duration := &Duration{
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
	duration := &Duration{
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
	duration := &Duration{
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
	duration := &Duration{
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
	duration := &Duration{
		Months: 13,
		Weeks:  10,
	}

	// Act
	result := duration.Normalize()

	// Assert
	assert.Equal(t, ErrWeeksNotWithYearsOrMonth, result)
}

func TestItRefusesYAndW(t *testing.T) {
	// Arrange
	duration := &Duration{
		Years: 13,
		Weeks: 10,
	}

	// Act
	result := duration.Normalize()

	// Assert
	assert.Equal(t, ErrWeeksNotWithYearsOrMonth, result)
}

func TestItFailsMoToDuration(t *testing.T) {
	// Arrange
	duration := &Duration{
		Months: 13,
		Weeks:  10,
	}

	// Act
	result, err := duration.ToDuration()

	// Assert
	assert.Equal(t, time.Duration(0), result)
	assert.Equal(t, ErrMonthsInDurationUseOverload, err)
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
