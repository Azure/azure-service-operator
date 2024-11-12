/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/go-logr/logr"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/config"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/test"
)

func Test_RenameProperties_RenamesExpectedProperty(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	personSpec := test.CreateSpec(test.Pkg2020, "Person", test.FullNameProperty, test.KnownAsProperty, test.FamilyNameProperty)
	personStatus := test.CreateStatus(test.Pkg2020, "Person", test.FullNameProperty, test.KnownAsProperty, test.FamilyNameProperty)
	person := test.CreateResource(test.Pkg2020, "Person", personSpec, personStatus)

	defs := astmodel.MakeTypeDefinitionSetFromDefinitions(person, personSpec, personStatus)

	cfg := config.NewConfiguration()
	omc := cfg.ObjectModelConfiguration
	g.Expect(omc.ModifyProperty(
		personSpec.Name(),
		"KnownAs",
		func(p *config.PropertyConfiguration) error {
			p.RenameTo.Set("Alias")
			return nil
		})).To(Succeed())

	initialState, err := RunTestPipeline(
		NewState(defs))
	g.Expect(err).To(Succeed())

	finalState, err := RunTestPipeline(
		initialState,
		RenameProperties(omc))
	g.Expect(err).To(Succeed())

	// When verifying the golden file, ensure the property has been renamed as expected
	test.AssertPackagesGenerateExpectedCode(t, finalState.definitions, test.DiffWithTypes(initialState.definitions))
}

func Test_RenameProperties_PopulatesExpectedARMProperty(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	idFactory := astmodel.NewIdentifierFactory()

	// Arrange API Version

	apiVersionValue := astmodel.MakeEnumValue("apiVersion", `"2020-06-01"`)
	apiVersion := astmodel.MakeTypeDefinition(
		astmodel.MakeInternalTypeName(test.Pkg2020, "APIVersion"),
		astmodel.NewEnumType(astmodel.StringType, apiVersionValue))

	apiVersionProperty := astmodel.NewPropertyDefinition("APIVersion", "apiVersion", apiVersion.Name()).
		MakeTypeOptional().
		MakeRequired()

	// Arrange Sample Resource

	personSpec := test.CreateSpec(
		test.Pkg2020,
		"Person",
		test.NameProperty,
		apiVersionProperty,
		test.FullNameProperty,
		test.KnownAsProperty,
		test.FamilyNameProperty)

	personStatus := test.CreateStatus(
		test.Pkg2020,
		"Person",
		test.NameProperty,
		test.FullNameProperty,
		test.KnownAsProperty,
		test.FamilyNameProperty)

	personResourceType := astmodel.NewResourceType(personSpec.Name(), personStatus.Name()).
		WithAPIVersion(apiVersion.Name(), apiVersionValue)

	person := astmodel.MakeTypeDefinition(astmodel.MakeInternalTypeName(test.Pkg2020, "Person"), personResourceType)

	// Arrange Initial State

	defs := astmodel.MakeTypeDefinitionSetFromDefinitions(
		person,
		personSpec,
		personStatus,
		apiVersion)

	cfg := config.NewConfiguration()
	omc := cfg.ObjectModelConfiguration
	g.Expect(omc.ModifyProperty(
		personSpec.Name(),
		"KnownAs",
		func(p *config.PropertyConfiguration) error {
			p.RenameTo.Set("Alias")
			return nil
		})).To(Succeed())

	initialState, err := RunTestPipeline(
		NewState(defs),
		CreateARMTypes(omc, idFactory, logr.Discard()),
		ApplyARMConversionInterface(idFactory, omc),
	)
	g.Expect(err).To(Succeed())

	// Act to generate our final state

	finalState, err := RunTestPipeline(
		initialState,
		RenameProperties(omc),
	)
	g.Expect(err).To(Succeed())

	// Assert - When verifying the golden file, ensure the property has been renamed as expected

	test.AssertPackagesGenerateExpectedCode(t, finalState.definitions, test.DiffWithTypes(initialState.definitions))
}

func Test_RenameProperties_WhenFlattening_PopulatesExpectedARMProperty(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	idFactory := astmodel.NewIdentifierFactory()

	// Arrange API Version

	apiVersionValue := astmodel.MakeEnumValue("apiVersion", `"2020-06-01"`)
	apiVersion := astmodel.MakeTypeDefinition(
		astmodel.MakeInternalTypeName(test.Pkg2020, "APIVersion"),
		astmodel.NewEnumType(astmodel.StringType, apiVersionValue))

	apiVersionProperty := astmodel.NewPropertyDefinition("APIVersion", "apiVersion", apiVersion.Name()).
		MakeTypeOptional().
		MakeRequired()

	// Arrange Sample Resource

	address := test.CreateObjectDefinition(test.Pkg2020, "Address", test.FullAddressProperty, test.CityProperty)
	addressProperty := astmodel.NewPropertyDefinition("Address", "address", address.Name()).
		SetFlatten(true).
		MakeTypeOptional()

	personSpec := test.CreateSpec(
		test.Pkg2020,
		"Person",
		test.NameProperty,
		apiVersionProperty,
		test.FullNameProperty,
		test.KnownAsProperty,
		test.FamilyNameProperty,
		addressProperty,
	)

	personStatus := test.CreateStatus(
		test.Pkg2020,
		"Person",
		test.NameProperty,
		test.FullNameProperty,
		test.KnownAsProperty,
		test.FamilyNameProperty,
		// addressProperty,
	)

	personResourceType := astmodel.NewResourceType(personSpec.Name(), personStatus.Name()).
		WithAPIVersion(apiVersion.Name(), apiVersionValue)

	person := astmodel.MakeTypeDefinition(astmodel.MakeInternalTypeName(test.Pkg2020, "Person"), personResourceType)

	// Arrange Initial State

	defs := astmodel.MakeTypeDefinitionSetFromDefinitions(
		person,
		personSpec,
		personStatus,
		apiVersion,
		address)

	cfg := config.NewConfiguration()
	omc := cfg.ObjectModelConfiguration
	g.Expect(omc.ModifyProperty(
		personSpec.Name(),
		"FullAddress",
		func(p *config.PropertyConfiguration) error {
			p.RenameTo.Set("PostalAddress")
			return nil
		})).To(Succeed())

	initialState, err := RunTestPipeline(
		NewState(defs),
		CreateARMTypes(omc, idFactory, logr.Discard()),
		ApplyARMConversionInterface(idFactory, omc),
		FlattenProperties(logr.Discard()),
		SimplifyDefinitions(),
	)
	g.Expect(err).To(Succeed())

	// Act to generate our final state

	finalState, err := RunTestPipeline(
		initialState,
		RenameProperties(omc),
	)
	g.Expect(err).To(Succeed())

	// Assert - When verifying the golden file, ensure the property has been renamed as expected

	test.AssertPackagesGenerateExpectedCode(t, finalState.definitions, test.DiffWithTypes(initialState.definitions))
}
