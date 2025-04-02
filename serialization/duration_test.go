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
	result := duration.string()

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
	result := duration.string()

	// Assert
	assert.Equal(t, 1, duration.Seconds)
	assert.Equal(t, 1, duration.Minutes)
	assert.Equal(t, "PT1M1S", result)
}

func TestItNormalizesMi(t *testing.T) {
	// Arrange
	duration := &duration{
		Minutes: 61,
	}

	// Act
	result := duration.string()

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
	result := duration.string()

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
	result := duration.string()

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
	result := duration.string()

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
	result := duration.string()

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
	result := duration.string()

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
	result := duration.normalize()

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
	result := duration.normalize()

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
	result, err := duration.toDuration()

	// Assert
	assert.Equal(t, time.Duration(0), result)
	assert.Equal(t, errMonthsInDurationUseOverload, err)
}

func TestItParsesMonth(t *testing.T) {
	// Act
	duration, err := durationFromString("P1M")

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, 1, duration.Months)
}

func TestItParsesMonthAndMinutes(t *testing.T) {
	// Act
	duration, err := durationFromString("P1MT1M")

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, 1, duration.Months)
	assert.Equal(t, 1, duration.Minutes)
}

func TestToDurationWithMonths(t *testing.T) {
	// Arrange
	duration := &duration{
		Years:   1,
		Months:  2,
		Days:    4,
		Hours:   5,
		Minutes: 6,
		Seconds: 7,
	}

	// Act
	result, err := duration.toDurationWithMonths(30)

	// Assert
	assert.Nil(t, err)
	expectedResult := time.Hour*24*365 + time.Hour*24*30*2 + time.Hour*24*4 + time.Hour*5 + time.Minute*6 + time.Second*7
	assert.Equal(t, expectedResult, result)
}
