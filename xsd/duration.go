package xsd

import (
	"errors"
	"regexp"
	"strconv"
	"time"
)

var (
	nanosecondsPerDay   = 24 * int64(time.Hour)
	nanosecondsPerMonth = 30 * nanosecondsPerDay
	nanosecondsPerYear  = 365 * nanosecondsPerDay
)

/*
	Duration is the [ISO 8601] extended format PnYn MnDTnH nMnS,
	where nY represents the number of years, nM the number of months,
	nD the number of days, 'T' is the date/time separator,
	nH the number of hours, nM the number of minutes and nS the number of seconds.
	The number of seconds can include decimal digits to arbitrary precision.
	PnYnMnDTnHnMnS

	Duration has the following ·constraining facets·:

	• pattern
	• enumeration
	• whiteSpace
	• maxInclusive
	• maxExclusive
	• minInclusive
	• minExclusive

	More info: https://www.w3.org/TR/xmlschema-2/#duration

	TODO: process restrictions
	TODO: Look at time.Duration go type
*/

type Duration time.Duration

func (d Duration) String() string {
	td := time.Duration(d).Nanoseconds()

	seconds := td % int64(time.Minute)
	if seconds > 0 {
		td -= seconds
		seconds = seconds / int64(time.Second)
	}
	minutes := td % int64(time.Hour)
	if minutes > 0 {
		td -= minutes
		minutes = minutes / int64(time.Minute)
	}
	hours := td % int64(24*time.Hour)
	if hours > 0 {
		td -= hours
		hours = hours / int64(time.Hour)
	}

	days := td / int64(24*time.Hour)

	result := "P" // time duration designator
	//years
	// if years > 0 {
	// 	result += strconv.FormatInt(years, 10) + "Y"
	// }
	// if months > 0 {
	// 	result += strconv.FormatInt(months, 10) + "M"
	// }
	if days > 0 {
		result += strconv.FormatInt(days, 10) + "D"
	}

	if hours > 0 || minutes > 0 || seconds > 0 {
		result += "T"
		if hours > 0 {
			result += strconv.FormatInt(hours, 10) + "H"
		}
		if minutes > 0 {
			result += strconv.FormatInt(minutes, 10) + "M"
		}
		if seconds > 0 {
			result += strconv.FormatInt(seconds, 10) + "S"
		}
	}

	if len(result) == 1 {
		result += "T0S"
	}

	return result
}

func (d *Duration) UnmarshalText(text []byte) error {
	durationRegex, _ := regexp.Compile(`^(-?)P(?:(\d+)Y)?(?:(\d+)M)?(?:(\d+)D)?(?:T(?:(\d+)H)?(?:(\d+)M)?(?:(\d*(?:\.\d+)?)S)?)?$`)

	if !durationRegex.MatchString(string(text)) {
		return errors.New("invalid format")
	}

	groups := durationRegex.FindStringSubmatch(string(text))

	var duration int64

	if groups[2] != "" {
		years, _ := strconv.Atoi(groups[2])
		duration += int64(years) * nanosecondsPerYear
	}
	if groups[3] != "" {
		months, _ := strconv.Atoi(groups[3])
		duration += int64(months) * nanosecondsPerMonth
	}
	if groups[4] != "" {
		days, _ := strconv.Atoi(groups[4])
		duration += int64(days) * nanosecondsPerDay
	}

	if groups[5] != "" {
		hours, _ := strconv.Atoi(groups[5])
		duration += int64(hours) * int64(time.Hour)
	}
	if groups[6] != "" {
		minutes, _ := strconv.Atoi(groups[6])
		duration += int64(minutes) * int64(time.Minute)
	}
	if groups[7] != "" {
		seconds, _ := strconv.Atoi(groups[7])
		duration += int64(seconds) * int64(time.Second)
	}

	*d = Duration(duration)

	return nil
}

func (d Duration) MarshalText() ([]byte, error) {
	return []byte(d.String()), nil
}
