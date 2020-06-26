// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package errhelp

import (
	"fmt"
	"regexp"
	"strings"
)

// StripErrorIDs takes an error and returns its string representation after filtering some common ID patterns
func StripErrorIDs(err error) string {
	patterns := []string{
		"RequestID=",
		"CorrelationId:\\s",
		"Tracking ID: ",
		"requestId",
	}
	reg := regexp.MustCompile(fmt.Sprintf(`(%s)\S+`, strings.Join(patterns, "|")))
	return reg.ReplaceAllString(err.Error(), "")

}

// StripErrorTimes removes the hours:minutes:seconds from a date to prevent updates to Status.Message from changing unnecessarily
func StripErrorTimes(err string) string {
	reg := regexp.MustCompile(`(T\d\d:\d\d:\d\d)\"`)
	return reg.ReplaceAllString(err, "")

}
