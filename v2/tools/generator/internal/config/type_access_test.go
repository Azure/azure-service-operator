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

func Test_TypeAccess_LookupWithPropertyOverride_ReturnsOverriddenValue_WhenPresent(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// Arrange
	ref := astmodel.MakeInternalTypeName(test.Pkg2020, "Person")
	typeValue := ExplicitCollections
	propertyValue := OmitEmptyProperties

	model := NewObjectModelConfiguration()
	g.Expect(
		model.ModifyType(
			ref,
			func(tc *TypeConfiguration) error {
				tc.PayloadType.Set(typeValue)
				return nil
			}),
	).To(Succeed())
	g.Expect(
		model.ModifyProperty(
			ref,
			"Name",
			func(tc *PropertyConfiguration) error {
				tc.PayloadType.Set(propertyValue)
				return nil
			}),
	).To(Succeed())

	access := makeTypeAccess[PayloadType](
		model,
		func(t *TypeConfiguration) *configurable[PayloadType] {
			return &t.PayloadType
		}).
		withPropertyOverride(
			func(p *PropertyConfiguration) *configurable[PayloadType] {
				return &p.PayloadType
			})

	// Act
	actual, ok := access.Lookup(ref, "Name")

	// Assert
	g.Expect(ok).To(BeTrue())
	g.Expect(actual).To(Equal(propertyValue))
}

func Test_TypeAccess_LookupWithPropertyOverride_ReturnsTypeValue_WhenNoOverridePresent(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// Arrange
	ref := astmodel.MakeInternalTypeName(test.Pkg2020, "Person")
	typeValue := ExplicitCollections
	propertyValue := OmitEmptyProperties

	model := NewObjectModelConfiguration()
	g.Expect(
		model.ModifyType(
			ref,
			func(tc *TypeConfiguration) error {
				tc.PayloadType.Set(typeValue)
				return nil
			}),
	).To(Succeed())
	g.Expect(
		model.ModifyProperty(
			ref,
			"Name",
			func(tc *PropertyConfiguration) error {
				tc.PayloadType.Set(propertyValue)
				return nil
			}),
	).To(Succeed())

	access := makeTypeAccess[PayloadType](
		model,
		func(t *TypeConfiguration) *configurable[PayloadType] {
			return &t.PayloadType
		}).
		withPropertyOverride(
			func(p *PropertyConfiguration) *configurable[PayloadType] {
				return &p.PayloadType
			})

	// Act
	actual, ok := access.Lookup(ref, "Status")

	// Assert
	g.Expect(ok).To(BeTrue())
	g.Expect(actual).To(Equal(typeValue))
}
