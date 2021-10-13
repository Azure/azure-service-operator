/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package testcommon

import (
	"reflect"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/Azure/azure-service-operator/v2/internal/controller/config"
)

func Test_CfgToKey_HasAllConfigDotValuesKeys(t *testing.T) {
	g := NewGomegaWithT(t)

	// this makes sure that if config.Values changes then cfgToKey is updated to match
	key := cfgToKey(config.Values{})

	valuesType := reflect.TypeOf(config.Values{})

	for _, field := range reflect.VisibleFields(valuesType) {
		g.Expect(key).To(ContainSubstring(field.Name + ":"))
	}
}
