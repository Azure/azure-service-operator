/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/config"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/test"
)

// TestGolden_InjectSpecInitializationFunctions creates parallel spec and status versions of a Person and checks
// to see the right code is generated, using golden files as reference.
// The spec Address is deliberately mapped to two different spec types (Address and Location) to ensure we generate
// both conversion functions correctly.
func TestGolden_InjectSpecInitializationFunctions(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	idFactory := astmodel.NewIdentifierFactory()

	// Define a spec type for addresses
	addressSpec := test.CreateObjectDefinition(
		test.Pkg2020,
		"Address"+astmodel.SpecSuffix,
		test.FullAddressProperty,
		test.CityProperty)

	// Define a status type for addresses
	addressStatus := test.CreateObjectDefinition(
		test.Pkg2020,
		"Address"+astmodel.StatusSuffix,
		test.FullAddressProperty,
		test.CityProperty)

	// Define a second status type for addresses, with a wholly different name
	locationStatus := test.CreateObjectDefinition(
		test.Pkg2020,
		"Location"+astmodel.StatusSuffix,
		test.FullAddressProperty,
		test.CityProperty,
		test.StatusProperty)

	// Define a spec property for addresses
	addressesSpecProperty := astmodel.NewPropertyDefinition(
		"Addresses",
		"addresses",
		astmodel.NewArrayType(addressSpec.Name())).
		WithDescription("All the places this person has lived.")

	// Define a status property for locations, using the same type as the spec property for addresses
	locationsSpecProperty := astmodel.NewPropertyDefinition(
		"Locations",
		"locations",
		astmodel.NewArrayType(addressSpec.Name())).
		WithDescription("Other places this person frequents.")

	// Define a status property for addresses
	addressesStatusProperty := astmodel.NewPropertyDefinition(
		"Addresses",
		"addresses",
		astmodel.NewArrayType(addressStatus.Name())).
		WithDescription("All the places this person has lived.")

	// Define a status proeprty for locations
	locationsStatusProperty := astmodel.NewPropertyDefinition(
		"Locations",
		"locations",
		astmodel.NewArrayType(locationStatus.Name())).
		WithDescription("Other places this person frequents.")

	personReferenceProperty := astmodel.NewPropertyDefinition(
		"PersonReference",
		"personReference",
		astmodel.ResourceReferenceType)

	personIDProperty := astmodel.NewPropertyDefinition(
		"PersonId",
		"personId",
		astmodel.StringType)

	spec := test.CreateSpec(
		test.Pkg2020,
		"Person",
		test.FullNameProperty,
		test.FamilyNameProperty,
		test.KnownAsProperty,
		addressesSpecProperty,
		locationsSpecProperty,
		personReferenceProperty,
	)

	status := test.CreateStatus(
		test.Pkg2020,
		"Person",
		test.FullNameProperty,
		test.FamilyNameProperty,
		test.KnownAsProperty,
		addressesStatusProperty,
		locationsStatusProperty,
		personIDProperty,
	)

	resource := test.CreateResource(test.Pkg2020, "Person", spec, status)

	defs := make(astmodel.TypeDefinitionSet)
	defs.AddAll(resource, spec, status, addressSpec, addressStatus, locationStatus)

	state := NewState(defs)

	cfg := config.NewConfiguration()
	finalState, err := RunTestPipeline(
		state,
		CreateStorageTypes(),            // First create the storage types
		CreateConversionGraph(cfg, "v"), // Then create the conversion graph
		InjectSpecInitializationFunctions(cfg, idFactory)) // Then create the spec initialization functions
	g.Expect(err).To(Succeed())

	test.AssertPackagesGenerateExpectedCode(t, finalState.Definitions())
}
