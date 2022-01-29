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
		if i == 0 && field.Name == "Values" {
			// Skip the embedded struct - we'll check for the fields
			// inside it.
			continue
		}
		g.Expect(key).To(ContainSubstring(field.Name + ":"))
	}
}
