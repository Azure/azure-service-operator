/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package config

import (
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"testing"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/test"
	. "github.com/onsi/gomega"
)

func Test_GroupAccess_Lookup_ReturnsConfiguredValue_WhenPresent(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// Arrange
	ref := test.Pkg2020
	value := ExplicitCollections

	model := NewObjectModelConfiguration()
	g.Expect(
		model.ModifyGroup(
			ref,
			func(gc *GroupConfiguration) error {
				gc.PayloadType.Set(value)
				return nil
			}),
	).To(Succeed())

	access := makeGroupAccess[PayloadType](
		model,
		func(g *GroupConfiguration) *configurable[PayloadType] {
			return &g.PayloadType
		})

	// Act
	actual, ok := access.Lookup(ref)

	// Assert
	g.Expect(ok).To(BeTrue())
	g.Expect(actual).To(Equal(value))
}

