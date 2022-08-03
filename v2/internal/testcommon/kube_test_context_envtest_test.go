/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package testcommon

import (
	"reflect"
	"testing"

	. "github.com/onsi/gomega"
)

func Test_CfgToKey_HasAllConfigDotValuesKeys(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// This makes sure that if config.Values or testConfig changes
	// then cfgToKey is updated to match.
	key := cfgToKey(testConfig{})

	testConfigType := reflect.TypeOf(testConfig{})

	for i, field := range reflect.VisibleFields(testConfigType) {
		// Skip the embedded struct and TerminateWhenDone field
		if (i == 0 && field.Name == "Values") || field.Name == "CountsTowardsLimit" {
			continue
		}
		g.Expect(key).To(ContainSubstring(field.Name + ":"))
	}
}
