// duration provides a partial implementation of ISO8601 durations. (no months)

package serialization

import (
	"bytes"
	"errors"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"time"
)

var (
	// errBadFormat is returned when parsing fails
	errBadFormat = errors.New("bad format string")

	errWeeksNotWithYearsOrMonth = errors.New("weeks are not allowed with years or months")

	errMonthsInDurationUseOverload = errors.New("months are not allowed with the ToDuration method, use the overload instead")

	full = regexp.MustCompile(`P((?P<year>\d+)Y)?((?P<month>\d+)M)?((?P<day>\d+)D)?(T((?P<hour>\d+)H)?((?P<minute>\d+)M)?((?P<second>\d+(?:\.\d+))S)?)?`)
	week = regexp.MustCompile(`P((?P<week>\d+)W)`)
)

type duration struct {
	Years        int
	Months       int
	Weeks        int
	Days         int
	Hours        int
	Minutes      int
	Seconds      int
	MilliSeconds int
}

func durationFromString(dur string) (*duration, error) {
	var (
		match []string
		re    *regexp.Regexp
	)

	if week.MatchString(dur) {
		match = week.FindStringSubmatch(dur)
		re = week
	} else if full.MatchString(dur) {
		match = full.FindStringSubmatch(dur)
		re = full
	} else {
		return nil, errBadFormat
	}

	d := &duration{}

	for i, name := range re.SubexpNames() {
		part := match[i]
		if i == 0 || name == "" || part == "" {
			continue
		}

		val, err := strconv.ParseFloat(part, 64)
		if err != nil {
			return nil, err
		}
		switch name {
		case "year":
			d.Years = int(val)
		case "month":
			d.Months = int(val)
		case "week":
			d.Weeks = int(val)
		case "day":
			d.Days = int(val)
		case "hour":
			c := time.Duration(val) * time.Hour
			d.Hours = int(c.Hours())
		case "minute":
			c := time.Duration(val) * time.Minute
			d.Minutes = int(c.Minutes())
		case "second":
			s, milli := math.Modf(val)
			d.Seconds = int(s)
			d.MilliSeconds = int(milli * 1000)
		default:
			return nil, fmt.Errorf("unknown field %s", name)
		}
	}

	return d, nil
}

// String prints out the value passed in.
func (d *duration) string() string {
	var s bytes.Buffer

	err := d.normalize()

	if err != nil {
		panic(err)
	}

	s.WriteString("P")
	if d.Years > 0 {
		s.WriteString(fmt.Sprintf("%dY", d.Years))
	}
	if d.Months > 0 {
		s.WriteString(fmt.Sprintf("%dM", d.Months))
	}
	if d.Weeks > 0 {
		s.WriteString(fmt.Sprintf("%dW", d.Weeks))
	}
	if d.Days > 0 {
		s.WriteString(fmt.Sprintf("%dD", d.Days))
	}
	if d.hasTimePart() {
		s.WriteString("T")
		if d.Hours > 0 {
			s.WriteString(fmt.Sprintf("%dH", d.Hours))
		}
		if d.Minutes > 0 {
			s.WriteString(fmt.Sprintf("%dM", d.Minutes))
		}
		if d.Seconds > 0 {
			s.WriteString(fmt.Sprintf("%dS", d.Seconds))
		}
	}

	return s.String()
}

// Normalize makes sure that all fields are represented as the smallest meaningful value possible by dividing them out by the conversion factor to the larger unit.
// e.g. if you have a duration of 10 day, 25 hour, and 61 minute, it will be normalized to 1 week 5 days, 2 hours, and 1 minute.
// this function does not normalize days to months, weeks to months or weeks to years as they do not always convert with the same value.
// it also won't normalize days to weeks if months or years are present, and will return an error if the value is invalid
func (d *duration) normalize() error {
	msToS := 1000
	StoM := 60
	MtoH := 60
	HtoD := 24
	DtoW := 7
	MtoY := 12
	if d.MilliSeconds >= msToS {
		d.Seconds += d.MilliSeconds / msToS
		d.MilliSeconds %= msToS
	}
	if d.Seconds >= StoM {
		d.Minutes += d.Seconds / StoM
		d.Seconds %= StoM
	}
	if d.Minutes >= MtoH {
		d.Hours += d.Minutes / MtoH
		d.Minutes %= MtoH
	}
	if d.Hours >= HtoD {
		d.Days += d.Hours / HtoD
		d.Hours %= HtoD
	}
	if d.Days >= DtoW && d.Months == 0 && d.Years == 0 {
		d.Weeks += d.Days / DtoW
		d.Days %= DtoW
	}
	if d.Months > MtoY {
		d.Years += d.Months / MtoY
		d.Months %= MtoY
	}

	if d.Weeks != 0 && (d.Years != 0 || d.Months != 0) {
		return errWeeksNotWithYearsOrMonth
	}

	return nil
	// a month is not always 30 days, so we don't normalize that
	// a month is not always 4 weeks, so we don't normalize that
	// a year is not always 52 weeks, so we don't normalize that
}

func (d *duration) hasTimePart() bool {
	return d.Hours != 0 || d.Minutes != 0 || d.Seconds != 0
}

func (d *duration) toDuration() (time.Duration, error) {
	if d.Months != 0 {
		return 0, errMonthsInDurationUseOverload
	}
	return d.toDurationWithMonths(31)
}

func (d *duration) toDurationWithMonths(daysInAMonth int) (time.Duration, error) {
	day := time.Hour * 24
	year := day * 365
	month := day * time.Duration(daysInAMonth)

	tot := time.Duration(0)

	err := d.normalize()
	if err != nil {
		return tot, err
	}

	tot += year * time.Duration(d.Years)
	tot += month * time.Duration(d.Months)
	tot += day * 7 * time.Duration(d.Weeks)
	tot += day * time.Duration(d.Days)
	tot += time.Hour * time.Duration(d.Hours)
	tot += time.Minute * time.Duration(d.Minutes)
	tot += time.Second * time.Duration(d.Seconds)

	return tot, nil
}
