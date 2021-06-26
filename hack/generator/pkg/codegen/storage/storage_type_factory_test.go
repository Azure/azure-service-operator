/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package storage

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/Azure/azure-service-operator/hack/generator/pkg/astmodel"
	"github.com/Azure/azure-service-operator/hack/generator/pkg/test"
)

func Test_StorageTypeFactory_GeneratesExpectedResults(t *testing.T) {
	g := NewGomegaWithT(t)

	group := "microsoft.person"
	idFactory := astmodel.NewIdentifierFactory()

	// Define Resources

	packageV1 := makeTestLocalPackageReference(group, "v20200101")
	packageV2 := makeTestLocalPackageReference(group, "v20211231")

	fullName := astmodel.NewPropertyDefinition("FullName", "fullName", astmodel.StringType).
		WithDescription("As would be used to address mail")
	familyName := astmodel.NewPropertyDefinition("FamilyName", "familyName", astmodel.StringType).
		WithDescription("Shared name of the family")
	knownAs := astmodel.NewPropertyDefinition("KnownAs", "knownAs", astmodel.StringType).
		WithDescription("How the person is generally known")
	vip := astmodel.NewPropertyDefinition("CustomerProgram", "customerProgram", astmodel.StringType).
		WithDescription("Level of customer programme (Silver/Gold/Platinum/Diamond)")

	v1definitions := createTestResource(packageV1, "Person", fullName, familyName, knownAs)
	v2definitions := createTestResource(packageV2, "Person", fullName, familyName, knownAs, vip)

	// Use Storage Factory

	factory := NewStorageTypeFactory(group, idFactory)
	factory.Add(v1definitions...)
	factory.Add(v2definitions...)

	outputTypes, err := factory.Types()
	g.Expect(err).To(BeNil())

	// Group type definitions by package

	groups := make(map[astmodel.PackageReference][]astmodel.TypeDefinition)
	for _, def := range outputTypes {
		ref := def.Name().PackageReference
		groups[ref] = append(groups[ref], def)
	}

	// Write files
	for _, defs := range groups {
		ref := defs[0].Name().PackageReference
		local, ok := ref.AsLocalPackage()
		g.Expect(ok).To(BeTrue())
		fileName := fmt.Sprintf("%s-%s", t.Name(), local.Version())
		file := test.CreateFileDefinition(defs...)
		test.AssertFileGeneratesExpectedCode(t, file, fileName)
	}
}

func createTestResource(
	pkg astmodel.PackageReference, name string, properties ...*astmodel.PropertyDefinition) []astmodel.TypeDefinition {

	statusProperty := astmodel.NewPropertyDefinition("Status", "status", astmodel.StringType)

	specName := astmodel.MakeTypeName(pkg, name+"_Spec")
	spec := astmodel.MakeTypeDefinition(
		specName,
		astmodel.NewObjectType().WithProperties(properties...))

	statusName := astmodel.MakeTypeName(pkg, name+"_Status")
	status := astmodel.MakeTypeDefinition(
		statusName,
		astmodel.NewObjectType().WithProperties(statusProperty))

	resource := astmodel.MakeTypeDefinition(
		astmodel.MakeTypeName(pkg, name),
		astmodel.NewResourceType(specName, statusName))

	return []astmodel.TypeDefinition{
		resource,
		spec,
		status,
	}
}

func makeTestLocalPackageReference(group string, version string) astmodel.LocalPackageReference {
	return astmodel.MakeLocalPackageReference("github.com/Azure/azure-service-operator/hack/generated", group, version)
}
