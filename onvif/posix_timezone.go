package onvif

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type PosixTimezone struct {
	Name    string
	Offset  posixTimezoneOffset
	Dailing *posixTimezoneDailing
}

type posixTimezoneOffset struct {
	Hours   int
	Minutes int
}

type posixTimezoneDailing struct {
	Name   string
	Offset posixTimezoneOffset
	From   posixTimezoneTime
	To     posixTimezoneTime
}

type posixTimezoneTime struct {
	Month   int
	Week    int
	Weekday int
	Hours   int
	Minutes int
	Seconds int
}

func newPosixTimezoneDailing(st *posixTimezoneOffset, name string) *posixTimezoneDailing {
	return &posixTimezoneDailing{
		Name: name,
		Offset: posixTimezoneOffset{
			Hours: st.Hours + 1, //TODO: + или -
		},
		From: posixTimezoneTime{
			Month:   3,
			Week:    2,
			Weekday: 0,
			Hours:   2,
		},
		To: posixTimezoneTime{
			Month:   11,
			Week:    1,
			Weekday: 0,
			Hours:   2,
		},
	}
}

func ParsePosixTimezone(timezone string) (*PosixTimezone, error) {
	type PosixTimezonePart int
	const (
		ST PosixTimezonePart = iota + 1
		STOFFSET
		STOFFSETMINUTES
		DT
		DTOFFSET
		DTOFFSETMINUTES
	)

	var err error
	tz := PosixTimezone{}

	parts := strings.Split(timezone, `,`)
	if len(parts) != 1 && len(parts) != 3 {
		return nil, errors.New("invalid format")
	}

	local := parts[0]
	state, start := ST, 0

	for i, ch := range local {
		nstate, nstart := state, start

		switch state {
		case ST:
			if ch < 'A' || ch > 'Z' {
				tz.Name = local[start:i]
				nstate, nstart = STOFFSET, i
			}
		case STOFFSET:
			if ch >= 'A' && ch <= 'Z' {
				tz.Offset.Hours, err = strconv.Atoi(local[start:i])
				nstate, nstart = DT, i
			}
			if ch == ':' || ch == ' ' {
				tz.Offset.Hours, err = strconv.Atoi(local[start:i])
				nstate, nstart = STOFFSETMINUTES, i+1
			}
		case STOFFSETMINUTES:
			if ch >= 'A' && ch <= 'Z' {
				tz.Offset.Minutes, err = strconv.Atoi(local[start:i])
				nstate, nstart = DT, i
			}
		case DT:
			if ch < 'A' && ch > 'Z' {
				tz.Dailing = newPosixTimezoneDailing(&tz.Offset, local[start:i])
				nstate, nstart = DTOFFSET, i
			}
		case DTOFFSET:
			if ch == ':' || ch == ' ' {
				tz.Dailing.Offset.Hours, err = strconv.Atoi(local[start:i])
				nstate, nstart = DTOFFSETMINUTES, i+1
			}
		case DTOFFSETMINUTES:
		}

		if err != nil {
			//TODO: add state, start to err desc
			return nil, err
		}
		state, start = nstate, nstart
	}

	switch state {
	case ST:
		tz.Name = local[start:]
	case STOFFSET:
		tz.Offset.Hours, err = strconv.Atoi(local[start:])
	case STOFFSETMINUTES:
		tz.Offset.Minutes, err = strconv.Atoi(local[start:])
	case DT:
		tz.Dailing = newPosixTimezoneDailing(&tz.Offset, local[start:])
	case DTOFFSET:
		tz.Dailing.Offset.Hours, err = strconv.Atoi(local[start:])
	case DTOFFSETMINUTES:
		tz.Dailing.Offset.Minutes, err = strconv.Atoi(local[start:])
	}
	if err != nil {
		//TODO: add state, start to err desc
		return nil, err
	}

	if len(parts) > 1 {
		err = parsePosixTimezoneTime(&tz.Dailing.From, parts[1])
		if err != nil {
			return &tz, err
		}
		err = parsePosixTimezoneTime(&tz.Dailing.To, parts[2])
		if err != nil {
			return &tz, err
		}
	}

	return &tz, nil
}

// date/time, time is optional.
// Here, date is in the Mm.n.d format, where:
// 		* Mm (1-12) for 12 months
//		* n (1-5) 1 for the first week and 5 for the last week in the month
//		* d (0-6) 0 for Sunday and 6 for Saturday
// Example: M3.2.0/2:00:00
func parsePosixTimezoneTime(t *posixTimezoneTime, timePart string) error {
	type PosixTimezoneTimePart int
	const (
		M PosixTimezoneTimePart = iota + 1
		month
		week
		weekend
		hours
		minutes
		seconds
	)

	state, start := M, 0

	for i, ch := range timePart {
		nstate, nstart := state, start
		var err error

		switch state {
		case M:
			if ch != 'M' {
				return errors.New("invalid format")
			}
			nstate, nstart = month, i+1
		case month:
			if ch == '.' {
				t.Month, err = strconv.Atoi(timePart[start:i])
				nstate, nstart = week, i+1
			}
		case week:
			if ch == '.' {
				t.Week, err = strconv.Atoi(timePart[start:i])
				nstate, nstart = weekend, i+1
			}
		case weekend:
			if ch == '/' {
				t.Weekday, err = strconv.Atoi(timePart[start:i])
				nstate, nstart = hours, i+1
			}
		case hours:
			if ch == ':' {
				t.Hours, err = strconv.Atoi(timePart[start:i])
				nstate, nstart = minutes, i+1
			}
		case minutes:
			if ch == ':' {
				t.Minutes, err = strconv.Atoi(timePart[start:i])
				nstate, nstart = seconds, i+1
			}
		case seconds:
		}

		if err != nil {
			//TODO: add state, start to err desc
			return err
		}
		state, start = nstate, nstart
	}

	var err error
	switch state {
	case weekend:
		t.Weekday, err = strconv.Atoi(timePart[start:])
	case seconds:
		t.Seconds, err = strconv.Atoi(timePart[start:])
	}
	if err != nil {
		//TODO: add state, start to err desc
		return err
	}

	return nil
}

func getTimeOfYear(year int, t posixTimezoneTime) time.Time {
	res := time.Date(
		year, time.Month(t.Month), 1,
		t.Hours, t.Minutes, t.Seconds, 0,
		time.UTC)

	weekday := int(res.Weekday())
	res = res.AddDate(0, 0, (7-weekday+t.Weekday)%7)

	if t.Week < 5 {
		res = res.AddDate(0, 0, 7*(t.Week-1))
	} else {
		res = res.AddDate(0, 0, 7*(t.Week-1))
		tt := res.AddDate(0, 0, 7)
		if int(tt.Month()) != t.Month {
			return res
		} else {
			return tt
		}
	}

	return res
}

func (tz *PosixTimezone) LocalToUTC(t time.Time) time.Time {
	offset := tz.Offset

	if tz.Dailing != nil {
		from := getTimeOfYear(t.Year(), tz.Dailing.From)
		to := getTimeOfYear(t.Year(), tz.Dailing.To)
		diffFrom, diffTo := t.Sub(from), t.Sub(to)

		fmt.Println(t, from, to)

		if tz.Dailing.From.Month < tz.Dailing.To.Month {
			if int64(diffFrom) > 0 && int64(diffTo) < 0 {
				offset = tz.Dailing.Offset
			}
		} else {
			if int64(diffFrom) > 0 || int64(diffTo) < 0 {
				offset = tz.Dailing.Offset
			}
		}
	}

	duration := time.Duration(offset.Hours)*time.Hour + time.Duration(offset.Minutes)*time.Minute

	return t.Add(duration)
}
