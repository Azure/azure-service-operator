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

func Test_TypeAccess_Lookup_ReturnsConfiguredValue_WhenPresent(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// Arrange
	ref := astmodel.MakeInternalTypeName(test.Pkg2020, "Person")
	value := ExplicitCollections

	model := NewObjectModelConfiguration()
	g.Expect(
		model.ModifyType(
			ref,
			func(tc *TypeConfiguration) error {
				tc.PayloadType.Set(value)
				return nil
			}),
	).To(Succeed())

	access := makeTypeAccess[PayloadType](
		model,
		func(t *TypeConfiguration) *configurable[PayloadType] {
			return &t.PayloadType
		})

	// Act
	actual, ok := access.Lookup(ref)

	// Assert
	g.Expect(ok).To(BeTrue())
	g.Expect(actual).To(Equal(value))
}
