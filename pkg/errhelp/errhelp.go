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
	}
	reg := regexp.MustCompile(fmt.Sprintf(`(%s)\S+`, strings.Join(patterns, "|")))
	return reg.ReplaceAllString(err.Error(), "")

}
