/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package config

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/test"
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

func Test_GroupAccess_LookupWithTypeOverride_ReturnsOverriddenValue_WhenPresent(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// Arrange
	ref := astmodel.MakeInternalTypeName(test.Pkg2020, "Person")
	groupValue := ExplicitCollections
	typeValue := OmitEmptyProperties

	model := NewObjectModelConfiguration()
	g.Expect(
		model.ModifyGroup(
			ref.InternalPackageReference(),
			func(tc *GroupConfiguration) error {
				tc.PayloadType.Set(groupValue)
				return nil
			}),
	).To(Succeed())
	g.Expect(
		model.ModifyType(
			ref,
			func(tc *TypeConfiguration) error {
				tc.PayloadType.Set(typeValue)
				return nil
			}),
	).To(Succeed())

	access := makeGroupAccess[PayloadType](
		model, func(g *GroupConfiguration) *configurable[PayloadType] {
			return &g.PayloadType
		}).
		withTypeOverride(func(t *TypeConfiguration) *configurable[PayloadType] {
			return &t.PayloadType
		})

	// Act
	actual, ok := access.Lookup(ref)

	// Assert
	g.Expect(ok).To(BeTrue())
	g.Expect(actual).To(Equal(typeValue))
}

func Test_GroupAccess_LookupWithTypeOverride_ReturnsGroupValue_WhenNoOverridePresent(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// Arrange
	ref := astmodel.MakeInternalTypeName(test.Pkg2020, "Person")
	groupValue := ExplicitCollections
	typeValue := OmitEmptyProperties

	model := NewObjectModelConfiguration()
	g.Expect(
		model.ModifyGroup(
			ref.InternalPackageReference(),
			func(tc *GroupConfiguration) error {
				tc.PayloadType.Set(groupValue)
				return nil
			}),
	).To(Succeed())
	g.Expect(
		model.ModifyType(
			astmodel.MakeInternalTypeName(test.Pkg2020, "Company"), // not ref
			func(tc *TypeConfiguration) error {
				tc.PayloadType.Set(typeValue)
				return nil
			}),
	).To(Succeed())

	access := makeGroupAccess[PayloadType](
		model, func(g *GroupConfiguration) *configurable[PayloadType] {
			return &g.PayloadType
		}).
		withTypeOverride(func(t *TypeConfiguration) *configurable[PayloadType] {
			return &t.PayloadType
		})

	// Act
	actual, ok := access.Lookup(ref)

	// Assert
	g.Expect(ok).To(BeTrue())
	g.Expect(actual).To(Equal(groupValue))
}
