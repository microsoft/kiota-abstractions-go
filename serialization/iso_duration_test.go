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

func TestItMakesAnNewISODurationFor1wOnly(t *testing.T) {
	isoDuration := NewDuration(0, 1, 0, 0, 0, 0, 0)
	assert.Equal(t, "P1W", isoDuration.String())
}

func TestItMakesAnNewISODurationFor1yAnd1dAnd1h(t *testing.T) {
	isoDuration := NewDuration(1, 0, 1, 1, 0, 0, 0)
	assert.Equal(t, "P1Y1DT1H", isoDuration.String())
}

func TestISODurationSetters(t *testing.T) {
	d := NewDuration(0, 0, 0, 0, 0, 0, 0)

	d.SetYears(2)
	assert.Equal(t, 2, d.GetYears())

	d.SetWeeks(3)
	assert.Equal(t, 3, d.GetWeeks())

	d.SetDays(4)
	assert.Equal(t, 4, d.GetDays())

	d.SetHours(5)
	assert.Equal(t, 5, d.GetHours())

	d.SetMinutes(6)
	assert.Equal(t, 6, d.GetMinutes())

	d.SetSeconds(7)
	assert.Equal(t, 7, d.GetSeconds())

	d.SetMilliSeconds(8)
	assert.Equal(t, 8, d.GetMilliSeconds())
}

func TestISODurationSettersNilReceiver(t *testing.T) {
	var d *ISODuration
	assert.NotPanics(t, func() {
		d.SetYears(1)
		d.SetWeeks(1)
		d.SetDays(1)
		d.SetHours(1)
		d.SetMinutes(1)
		d.SetSeconds(1)
		d.SetMilliSeconds(1)
	})
}
