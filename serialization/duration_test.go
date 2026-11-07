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
	// Arrange: 8 days is not an exact multiple of 7, so weeks should not be used
	duration := &duration{
		Days: 8,
	}

	// Act
	result := duration.string()

	// Assert
	assert.Equal(t, 0, duration.Weeks)
	assert.Equal(t, 8, duration.Days)
	assert.Equal(t, "P8D", result)
}

func TestItNormalizesExactWeeksFromDays(t *testing.T) {
	// Arrange: 14 days is exactly 2 weeks, so it should be normalized to P2W
	duration := &duration{
		Days: 14,
	}

	// Act
	result := duration.string()

	// Assert
	assert.Equal(t, 2, duration.Weeks)
	assert.Equal(t, 0, duration.Days)
	assert.Equal(t, "P2W", result)
}

func TestItNormalizes7DaysToWeek(t *testing.T) {
	// Arrange: 7 days is exactly 1 week
	duration := &duration{
		Days: 7,
	}

	// Act
	result := duration.string()

	// Assert
	assert.Equal(t, 1, duration.Weeks)
	assert.Equal(t, 0, duration.Days)
	assert.Equal(t, "P1W", result)
}

func TestItRefusesWeeksWithDays(t *testing.T) {
	// Arrange: weeks cannot coexist with days per ISO 8601
	duration := &duration{
		Weeks: 1,
		Days:  1,
	}

	// Act
	result := duration.normalize()

	// Assert
	assert.Equal(t, errWeeksNotAlone, result)
}

func TestItRefusesWeeksWithHours(t *testing.T) {
	// Arrange: weeks cannot coexist with time components per ISO 8601
	duration := &duration{
		Weeks: 1,
		Hours: 2,
	}

	// Act
	result := duration.normalize()

	// Assert
	assert.Equal(t, errWeeksNotAlone, result)
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

func TestItNormalizes12MonthsTo1Year(t *testing.T) {
	duration := &duration{
		Months: 12,
	}
	result := duration.string()
	assert.Equal(t, 0, duration.Months)
	assert.Equal(t, 1, duration.Years)
	assert.Equal(t, "P1Y", result)
}

func TestItNormalizes24MonthsTo2Years(t *testing.T) {
	duration := &duration{
		Months: 24,
	}
	result := duration.string()
	assert.Equal(t, 0, duration.Months)
	assert.Equal(t, 2, duration.Years)
	assert.Equal(t, "P2Y", result)
}

func TestToDurationPreservesMilliseconds(t *testing.T) {
	d := NewDuration(0, 0, 0, 0, 0, 1, 500)
	dur, err := d.ToDuration()
	assert.NoError(t, err)
	assert.Equal(t, 1500*time.Millisecond, dur)

	subSecond := NewDuration(0, 0, 0, 0, 0, 0, 450)
	subDur, err := subSecond.ToDuration()
	assert.NoError(t, err)
	assert.Equal(t, 450*time.Millisecond, subDur)
}

func TestFromDurationToDurationRoundTripMilliseconds(t *testing.T) {
	expected := 2500 * time.Millisecond
	isoDur := FromDuration(expected)
	actual, err := isoDur.ToDuration()
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func TestCompositeDaysWithTimeDoesNotNormalizeToWeeksOrPanic(t *testing.T) {
	testCases := []struct {
		duration *ISODuration
		expected string
	}{
		{
			duration: NewDuration(0, 0, 7, 1, 0, 0, 0),
			expected: "P7DT1H",
		},
		{
			duration: NewDuration(0, 0, 14, 0, 30, 0, 0),
			expected: "P14DT30M",
		},
		{
			duration: NewDuration(0, 0, 21, 0, 0, 45, 0),
			expected: "P21DT45S",
		},
		{
			duration: NewDuration(0, 0, 28, 2, 15, 30, 0),
			expected: "P28DT2H15M30S",
		},
	}

	for _, tc := range testCases {
		assert.NotPanics(t, func() {
			actual := tc.duration.String()
			assert.Equal(t, tc.expected, actual)
		})
	}

	parsed, err := ParseISODuration("P7DT1H")
	assert.NoError(t, err)
	assert.NotPanics(t, func() {
		assert.Equal(t, "P7DT1H", parsed.String())
	})
}

func TestCompositeDurationConversionsAndPureWeekPreservation(t *testing.T) {
	fromDur := FromDuration(7*24*time.Hour + time.Hour)
	assert.NotPanics(t, func() {
		assert.Equal(t, "P7DT1H", fromDur.String())
	})

	composite := NewDuration(0, 0, 7, 2, 30, 0, 0)
	dur, err := composite.ToDuration()
	assert.NoError(t, err)
	expectedDur := 7*24*time.Hour + 2*time.Hour + 30*time.Minute
	assert.Equal(t, expectedDur, dur)

	pure7 := NewDuration(0, 0, 7, 0, 0, 0, 0)
	assert.Equal(t, "P1W", pure7.String())
	pure14 := NewDuration(0, 0, 14, 0, 0, 0, 0)
	assert.Equal(t, "P2W", pure14.String())

	invalidWeekWithHour := NewDuration(0, 1, 0, 1, 0, 0, 0)
	_, err = invalidWeekWithHour.ToDuration()
	assert.ErrorIs(t, err, errWeeksNotAlone)
}
