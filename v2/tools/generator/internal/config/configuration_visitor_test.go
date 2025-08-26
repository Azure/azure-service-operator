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

func TestConfigurationVisitor_WhenVisitingASpecificVersion_VisitsExpectedVersion(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	omc := createTestObjectModelConfigurationForVisitor(t)
	seen := set.Make[string]()
	visitor := newSingleVersionConfigurationVisitor(
		test.Pkg2022,
		func(configuration *VersionConfiguration) error {
			seen.Add(configuration.name)
			return nil
		})

	g.Expect(visitor.visit(omc)).To(Succeed())
	g.Expect(seen).To(HaveLen(1))
	g.Expect(seen).To(HaveKey(test.Pkg2022.Version()))
}

func TestConfigurationVisitor_WhenVisitingEveryType_VisitsExpectedTypes(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	omc := createTestObjectModelConfigurationForVisitor(t)
	seen := set.Make[string]()
	visitor := newEveryTypeConfigurationVisitor(
		func(configuration *TypeConfiguration) error {
			seen.Add(configuration.name)
			return nil
		})

	g.Expect(visitor.visit(omc)).To(Succeed())
	g.Expect(seen).To(HaveLen(2))
	g.Expect(seen).To(HaveKey("SimplePerson"))
	g.Expect(seen).To(HaveKey("Person"))
}

func TestConfigurationVisitor_WhenVisitingASpecificType_VisitsExpectedType(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	omc := createTestObjectModelConfigurationForVisitor(t)
	seen := set.Make[string]()
	name := astmodel.MakeInternalTypeName(test.Pkg2022, "Person")
	visitor := newSingleTypeConfigurationVisitor(
		name,
		func(configuration *TypeConfiguration) error {
			seen.Add(configuration.name)
			return nil
		})

	g.Expect(visitor.visit(omc)).To(Succeed())
	g.Expect(seen).To(HaveLen(1))
	g.Expect(seen).To(HaveKey("Person"))
}

func TestConfigurationVisitor_WhenVisitingEveryProperty_VisitsExpectedProperties(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	omc := createTestObjectModelConfigurationForVisitor(t)
	seen := set.Make[string]()
	visitor := newEveryPropertyConfigurationVisitor(
		func(configuration *PropertyConfiguration) error {
			seen.Add(configuration.name)
			return nil
		})

	g.Expect(visitor.visit(omc)).To(Succeed())
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

	omc := createTestObjectModelConfigurationForVisitor(t)
	seen := set.Make[string]()
	name := astmodel.MakeInternalTypeName(test.Pkg2022, "Person")
	visitor := newSinglePropertyConfigurationVisitor(
		name,
		"KnownAs",
		func(configuration *PropertyConfiguration) error {
			seen.Add(configuration.name)
			return nil
		})

	g.Expect(visitor.visit(omc)).To(Succeed())
	g.Expect(seen).To(HaveLen(1))
	g.Expect(seen).To(HaveKey("KnownAs"))
}

func TestConfigurationVisitor_WhenVisitingAllGroups_VisitsExpectedGroups(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	omc := createTestObjectModelConfigurationForVisitor(t)
	seen := set.Make[string]()
	visitor := newEveryGroupConfigurationVisitor(
		func(configuration *GroupConfiguration) error {
			seen.Add(configuration.name)
			return nil
		})

	g.Expect(visitor.visit(omc)).To(Succeed())
	g.Expect(seen).To(HaveLen(2))
	g.Expect(seen).To(HaveKey(test.Group))
	g.Expect(seen).To(HaveKey("OtherGroup"))
}

func TestConfigurationVisitor_WhenVisitingAllVersions_VisitsExpectedVersions(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	omc := createTestObjectModelConfigurationForVisitor(t)
	seen := set.Make[string]()
	visitor := newEveryVersionConfigurationVisitor(
		func(configuration *VersionConfiguration) error {
			seen.Add(configuration.name)
			return nil
		})

	g.Expect(visitor.visit(omc)).To(Succeed())
	g.Expect(seen).To(HaveLen(4))
	g.Expect(seen).To(HaveKey(test.Pkg2020.Version()))
	g.Expect(seen).To(HaveKey(test.Pkg2022.Version()))
	g.Expect(seen).To(HaveKey("v1"))
	g.Expect(seen).To(HaveKey("v2"))
}

func createTestObjectModelConfigurationForVisitor(t *testing.T) *ObjectModelConfiguration {
	g := NewGomegaWithT(t)

	lastName := NewPropertyConfiguration("LastName")
	firstName := NewPropertyConfiguration("FirstName")

	person2020 := NewTypeConfiguration("SimplePerson")
	g.Expect(person2020.addProperty(lastName.name, lastName)).To(Succeed())
	g.Expect(person2020.addProperty(firstName.name, firstName)).To(Succeed())

	version2020 := NewVersionConfiguration(test.Pkg2020.Version())
	g.Expect(version2020.addType(person2020.name, person2020)).To(Succeed())

	fullName := NewPropertyConfiguration("FullName")
	knownAs := NewPropertyConfiguration("KnownAs")
	familyName := NewPropertyConfiguration("FamilyName")

	person2022 := NewTypeConfiguration("Person")
	g.Expect(person2022.addProperty(fullName.name, fullName)).To(Succeed())
	g.Expect(person2022.addProperty(knownAs.name, knownAs)).To(Succeed())
	g.Expect(person2022.addProperty(familyName.name, familyName)).To(Succeed())

	version2022 := NewVersionConfiguration(test.Pkg2022.Version())
	g.Expect(version2022.addType(person2022.name, person2022)).To(Succeed())

	group := NewGroupConfiguration(test.Group)
	g.Expect(group.addVersion(version2020.name, version2020)).To(Succeed())
	g.Expect(group.addVersion(version2022.name, version2022)).To(Succeed())

	group2 := NewGroupConfiguration("OtherGroup")
	g.Expect(group2.addVersion(
		"v1",
		NewVersionConfiguration("v1"))).To(Succeed())
	g.Expect(group2.addVersion(
		"v2",
		NewVersionConfiguration("v2"))).To(Succeed())

	modelConfig := NewObjectModelConfiguration()
	g.Expect(modelConfig.addGroup(group.name, group)).To(Succeed())
	g.Expect(modelConfig.addGroup(group2.name, group2)).To(Succeed())

	return modelConfig
}
