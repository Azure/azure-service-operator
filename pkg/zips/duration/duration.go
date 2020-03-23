/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package duration

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type (
	ISO8601 struct {
		time.Duration
	}
)

var (
	durationRegex = regexp.MustCompile(`P([\d\.]+Y)?([\d\.]+M)?([\d\.]+D)?T?([\d\.]+H)?([\d\.]+M)?([\d\.]+?S)?`)
)

/*
Duration marshalling

TODO: remove after go v2: https://github.com/golang/go/issues/10275
*/

func (d ISO8601) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, d.String())), nil
}

func (d *ISO8601) UnmarshalJSON(b []byte) error {
	var v interface{}
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	switch value := v.(type) {
	case float64:
		d.Duration = time.Duration(value)
		return nil
	case string:
		pd, err := parseDuration(value)
		if err != nil {
			return err
		}

		d.Duration = pd
		return nil
	default:
		return errors.New("invalid duration")
	}
}

func (d ISO8601) String() string {
	if d.Duration == time.Duration(0) {
		return "PT0S"
	}

	builder := new(strings.Builder)
	builder.WriteString("P")
	part, duration := stringPart(d.Duration, time.Hour*24*365, "Y")
	if part != "" {
		builder.WriteString(part)
	}

	part, duration = stringPart(duration, time.Hour*24*30, "M")
	if part != "" {
		builder.WriteString(part)
	}

	part, duration = stringPart(duration, time.Hour*24*7, "W")
	if part != "" {
		builder.WriteString(part)
	}

	part, duration = stringPart(duration, time.Hour*24, "D")
	if part != "" {
		builder.WriteString(part)
	}

	if duration <= 0 {
		return builder.String()
	}

	builder.WriteString("T")
	part, duration = stringPart(duration, time.Hour, "H")
	if part != "" {
		builder.WriteString(part)
	}

	part, duration = stringPart(duration, time.Minute, "M")
	if part != "" {
		builder.WriteString(part)
	}

	part, _ = stringPart(duration, time.Second, "S")
	if part != "" {
		builder.WriteString(strings.TrimRight(part, "0"))
	}

	return builder.String()
}

func stringPart(duration time.Duration, unit time.Duration, unitSymbol string) (string, time.Duration) {
	if duration >= unit {
		if unitSymbol == "S" {
			remainder := duration % unit
			if remainder == 0 {
				units := int64(duration / unit)
				return fmt.Sprintf("%dS", units), 0
			}

			units := float64(duration) / float64(unit)
			return fmt.Sprintf("%sS", strconv.FormatFloat(units, 'f', -1, 64)), 0
		}

		units := int64(duration / unit)
		return fmt.Sprintf("%d%s", units, unitSymbol), duration - time.Duration(units*int64(unit))
	}
	return "", duration
}

// parseDuration converts a ISO8601 duration into a time.Duration
func parseDuration(str string) (time.Duration, error) {
	matches := durationRegex.FindStringSubmatch(str)

	years, err := parseDurationPart(matches[1], time.Hour*24*365)
	if err != nil {
		return 0, err
	}

	months, err := parseDurationPart(matches[2], time.Hour*24*30)
	if err != nil {
		return 0, err
	}

	days, err := parseDurationPart(matches[3], time.Hour*24)
	if err != nil {
		return 0, err
	}

	hours, err := parseDurationPart(matches[4], time.Hour)
	if err != nil {
		return 0, err
	}

	minutes, err := parseDurationPart(matches[5], time.Second*60)
	if err != nil {
		return 0, err
	}

	seconds, err := parseDurationPart(matches[6], time.Second)
	if err != nil {
		return 0, err
	}

	return years + months + days + hours + minutes + seconds, nil
}

func parseDurationPart(value string, unit time.Duration) (time.Duration, error) {
	if len(value) != 0 {
		parsed, err := strconv.ParseFloat(value[:len(value)-1], 64)
		if err != nil {
			return 0, err
		}
		return time.Duration(float64(unit) * parsed), nil
	}
	return 0, nil
}
