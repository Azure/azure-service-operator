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

func TestConfigurationVisitor_WhenVisitingEveryType_VisitsExpectedTypes(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	omc := createTestObjectModelConfigurationForVisitor()
	seen := astmodel.MakeStringSet()
	visitor := NewEveryTypeConfigurationVisitor(
		func(configuration *TypeConfiguration) error {
			seen.Add(configuration.name)
			return nil
		})

	visitor.Visit(omc)

	g.Expect(seen).To(HaveLen(2))
	g.Expect(seen).To(HaveKey("SimplePerson"))
	g.Expect(seen).To(HaveKey("Person"))
}

func TestConfigurationVisitor_WhenVisitingASpecificType_VisitsExpectedType(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	omc := createTestObjectModelConfigurationForVisitor()
	seen := astmodel.MakeStringSet()
	name := astmodel.MakeTypeName(test.Pkg2022, "Person")
	visitor := NewSingleTypeConfigurationVisitor(
		name,
		func(configuration *TypeConfiguration) error {
			seen.Add(configuration.name)
			return nil
		})

	visitor.Visit(omc)

	g.Expect(seen).To(HaveLen(1))
	g.Expect(seen).To(HaveKey("Person"))
}

func TestConfigurationVisitor_WhenVisitingEveryProperty_VisitsExpectedProperties(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	omc := createTestObjectModelConfigurationForVisitor()
	seen := astmodel.MakeStringSet()
	visitor := NewEveryPropertyConfigurationVisitor(
		func(configuration *PropertyConfiguration) error {
			seen.Add(configuration.name)
			return nil
		})

	visitor.Visit(omc)

	g.Expect(seen).To(HaveLen(5))
	g.Expect(seen).To(HaveKey("FamilyName"))
	g.Expect(seen).To(HaveKey("FirstName"))
	g.Expect(seen).To(HaveKey("FullName"))
	g.Expect(seen).To(HaveKey("KnownAs"))
	g.Expect(seen).To(HaveKey("LastName"))
}

func TestConfigurationVisitor_WhenVisitingASpecificProperty_VisitsExpectedProperty(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	omc := createTestObjectModelConfigurationForVisitor()
	seen := astmodel.MakeStringSet()
	name := astmodel.MakeTypeName(test.Pkg2022, "Person")
	visitor := NewSinglePropertyConfigurationVisitor(
		name,
		"KnownAs",
		func(configuration *PropertyConfiguration) error {
			seen.Add(configuration.name)
			return nil
		})

	visitor.Visit(omc)

	g.Expect(seen).To(HaveLen(1))
	g.Expect(seen).To(HaveKey("KnownAs"))
}

func createTestObjectModelConfigurationForVisitor() *ObjectModelConfiguration {

	lastName := NewPropertyConfiguration("LastName")
	firstName := NewPropertyConfiguration("FirstName")

	person2020 := NewTypeConfiguration("SimplePerson").Add(lastName).Add(firstName)
	version2020 := NewVersionConfiguration(test.Pkg2020.Version()).Add(person2020)

	fullName := NewPropertyConfiguration("FullName")
	knownAs := NewPropertyConfiguration("KnownAs")
	familyName := NewPropertyConfiguration("FamilyName")

	person2022 := NewTypeConfiguration("Person").Add(fullName).Add(knownAs).Add(familyName)
	version2022 := NewVersionConfiguration(test.Pkg2022.Version()).Add(person2022)

	group := NewGroupConfiguration(test.Group).Add(version2020).Add(version2022)
	modelConfig := NewObjectModelConfiguration().Add(group)

	return modelConfig
}
