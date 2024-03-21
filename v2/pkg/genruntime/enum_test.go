/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package genruntime

import "testing"

func Test_ToEnum_WhenCalled_ReturnsExpectedResult(t *testing.T) {
	t.Parallel()

	type testResult string

	const (
		success = testResult("Success")
		failure = testResult("Failure")
		unknown = testResult("Unknown")
		pending = testResult("Pending")
	)

	values := map[string]testResult{
		"success": success,
		"failure": failure,
		"unknown": unknown,
		"pending": pending,
	}

	cases := []struct {
		str      string
		expected testResult
	}{
		{"success", success},
		{"Failure", failure},
		{"UNKNOWN", unknown},
		{"fatal", testResult("fatal")},
	}

	for _, c := range cases {
		c := c
		t.Run(
			c.str,
			func(t *testing.T) {
				t.Parallel()
				actual := ToEnum(c.str, values)
				if actual != c.expected {
					t.Errorf("Expected %s, but got %s", c.expected, actual)
				}
			})
	}
}
