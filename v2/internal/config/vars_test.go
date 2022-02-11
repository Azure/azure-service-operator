// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package config_test

import (
	"reflect"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/Azure/azure-service-operator/v2/internal/config"
)

func Test_String_HasAllKeys(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// This makes sure that if config.Values or testConfig changes
	// then cfgToKey is updated to match.
	s := config.Values{}.String()

	testConfigType := reflect.TypeOf(config.Values{})

	for _, field := range reflect.VisibleFields(testConfigType) {
		g.Expect(s).To(ContainSubstring(field.Name + ":"))
	}
}
