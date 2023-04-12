/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package testcommon

import (
	"fmt"
	"strings"

	"github.com/onsi/gomega/format"
	"github.com/onsi/gomega/types"
)

func EqualsIgnoreCase(expected string) types.GomegaMatcher {
	return &equalsIgnoreCaseMatcher{
		Expected: expected,
	}
}

type equalsIgnoreCaseMatcher struct {
	Expected string
}

func (matcher *equalsIgnoreCaseMatcher) Match(actual interface{}) (success bool, err error) {
	actualString, ok := actual.(string)
	if !ok {
		return false, fmt.Errorf("EqualsIgnoreCaseMatcher matcher requires a string .\nGot:%s", format.Object(actual, 1))
	}

	return strings.EqualFold(actualString, matcher.Expected), nil
}

func (matcher *equalsIgnoreCaseMatcher) FailureMessage(actual interface{}) (message string) {
	return format.Message(actual, "to match (case insensitive)", matcher.Expected)
}

func (matcher *equalsIgnoreCaseMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	return format.Message(actual, "not to match (case insensitive)", matcher.Expected)
}
