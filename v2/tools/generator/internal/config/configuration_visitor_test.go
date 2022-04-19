/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package config

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/Azure/azure-service-operator/v2/internal/set"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/test"
)

func TestConfigurationVisitor_WhenVisitingEveryType_VisitsExpectedTypes(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	omc := createTestObjectModelConfigurationForVisitor()
	seen := set.Make[string]()
	visitor := NewEveryTypeConfigurationVisitor(
		func(configuration *TypeConfiguration) error {
			seen.Add(configuration.name)
			return nil
		})

	g.Expect(visitor.Visit(omc)).To(Succeed())
	g.Expect(seen).To(HaveLen(2))
	g.Expect(seen).To(HaveKey("SimplePerson"))
	g.Expect(seen).To(HaveKey("Person"))
}

func TestConfigurationVisitor_WhenVisitingASpecificType_VisitsExpectedType(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	omc := createTestObjectModelConfigurationForVisitor()
	seen := set.Make[string]()
	name := astmodel.MakeTypeName(test.Pkg2022, "Person")
	visitor := NewSingleTypeConfigurationVisitor(
		name,
		func(configuration *TypeConfiguration) error {
			seen.Add(configuration.name)
			return nil
		})

	g.Expect(visitor.Visit(omc)).To(Succeed())
	g.Expect(seen).To(HaveLen(1))
	g.Expect(seen).To(HaveKey("Person"))
}

func TestConfigurationVisitor_WhenVisitingEveryProperty_VisitsExpectedProperties(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	omc := createTestObjectModelConfigurationForVisitor()
	seen := set.Make[string]()
	visitor := NewEveryPropertyConfigurationVisitor(
		func(configuration *PropertyConfiguration) error {
			seen.Add(configuration.name)
			return nil
		})

	g.Expect(visitor.Visit(omc)).To(Succeed())
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
	seen := set.Make[string]()
	name := astmodel.MakeTypeName(test.Pkg2022, "Person")
	visitor := NewSinglePropertyConfigurationVisitor(
		name,
		"KnownAs",
		func(configuration *PropertyConfiguration) error {
			seen.Add(configuration.name)
			return nil
		})

	g.Expect(visitor.Visit(omc)).To(Succeed())
	g.Expect(seen).To(HaveLen(1))
	g.Expect(seen).To(HaveKey("KnownAs"))
}

func createTestObjectModelConfigurationForVisitor() *ObjectModelConfiguration {
	lastName := NewPropertyConfiguration("LastName")
	firstName := NewPropertyConfiguration("FirstName")

	person2020 := NewTypeConfiguration("SimplePerson")
	person2020.add(lastName)
	person2020.add(firstName)

	version2020 := NewVersionConfiguration(test.Pkg2020.Version())
	version2020.add(person2020)

	fullName := NewPropertyConfiguration("FullName")
	knownAs := NewPropertyConfiguration("KnownAs")
	familyName := NewPropertyConfiguration("FamilyName")

	person2022 := NewTypeConfiguration("Person")
	person2022.add(fullName)
	person2022.add(knownAs)
	person2022.add(familyName)

	version2022 := NewVersionConfiguration(test.Pkg2022.Version())
	version2022.add(person2022)

	group := NewGroupConfiguration(test.Group)
	group.add(version2020)
	group.add(version2022)

	modelConfig := NewObjectModelConfiguration()
	modelConfig.add(group)

	return modelConfig
}
