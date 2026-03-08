package serialization

import (
	assert "github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestItParsesADuration(t *testing.T) {
	duration, err := ParseISODuration("PT1H")
	assert.Nil(t, err)
	assert.Equal(t, "PT1H", duration.String())
}

func TestItMakesAnISODurationFromATimeDurationFor1h(t *testing.T) {
	duration := time.Duration(1) * time.Hour
	isoDuration := FromDuration(duration)
	assert.Equal(t, "PT1H", isoDuration.String())
}

func TestItMakesAnISODurationFromATimeDurationFor1d(t *testing.T) {
	duration := time.Duration(24) * time.Hour
	isoDuration := FromDuration(duration)
	assert.Equal(t, "P1D", isoDuration.String())
}

func TestItMakesAnNewISODurationFor1h(t *testing.T) {
	isoDuration := NewDuration(0, 0, 0, 1, 0, 0, 0)
	assert.Equal(t, "PT1H", isoDuration.String())
}

func TestItMakesAnNewISODurationFor1dAnd1h(t *testing.T) {
	isoDuration := NewDuration(0, 0, 1, 1, 0, 0, 0)
	assert.Equal(t, "P1DT1H", isoDuration.String())
}

func TestItMakesAnNewISODurationFor1wAnd1dAnd1h(t *testing.T) {
	isoDuration := NewDuration(0, 1, 1, 1, 0, 0, 0)
	assert.Equal(t, "P1W1DT1H", isoDuration.String())
}

func TestItMakesAnNewISODurationFor1yAnd1dAnd1h(t *testing.T) {
	isoDuration := NewDuration(1, 0, 1, 1, 0, 0, 0)
	assert.Equal(t, "P1Y1DT1H", isoDuration.String())
}

func TestNewISODurationFromStringPreservesDayFormat(t *testing.T) {
	cases := []string{"P90D", "P365D", "P180D", "P30D"}
	for _, c := range cases {
		d, err := NewISODurationFromString(c)
		assert.Nil(t, err)
		assert.Equal(t, c, d.String())
	}
}

func TestNewISODurationFromStringPreservesWeekFormat(t *testing.T) {
	d, err := NewISODurationFromString("P2W")
	assert.Nil(t, err)
	assert.Equal(t, "P2W", d.String())
}

func TestNewISODurationFromStringPreservesTimeFormat(t *testing.T) {
	d, err := NewISODurationFromString("PT1H")
	assert.Nil(t, err)
	assert.Equal(t, "PT1H", d.String())
}

func TestNewISODurationFromStringParsesFieldsCorrectly(t *testing.T) {
	d, err := NewISODurationFromString("P90D")
	assert.Nil(t, err)
	assert.Equal(t, 90, d.GetDays())
	assert.Equal(t, 0, d.GetWeeks())
}

func TestNewISODurationFromStringReturnsErrorOnInvalidInput(t *testing.T) {
	_, err := NewISODurationFromString("invalid")
	assert.NotNil(t, err)
}
