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

func TestGolden_InjectSpecInitializationFunctions(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	idFactory := astmodel.NewIdentifierFactory()

	//TODO: Add new properties to the spec and status with identical names but typenames with different names
	// Need to extend proeprty conversions to handle these cases

	addressSpec := test.CreateObjectDefinition(
		test.Pkg2020,
		"Address"+astmodel.SpecSuffix,
		test.FullAddressProperty,
		test.CityProperty)

	addressSpecProperty := astmodel.NewPropertyDefinition("Address", "address", addressSpec.Name())

	// We want a status type that has a wholly different name from the spec type
	locationStatus := test.CreateObjectDefinition(
		test.Pkg2020,
		"Location"+astmodel.StatusSuffix,
		test.FullAddressProperty,
		test.CityProperty,
		test.StatusProperty)

	addressStatusProperty := astmodel.NewPropertyDefinition("Address", "address", locationStatus.Name())

	spec := test.CreateSpec(
		test.Pkg2020,
		"Person",
		test.FullNameProperty,
		test.FamilyNameProperty,
		test.KnownAsProperty,
		addressSpecProperty)

	status := test.CreateStatus(
		test.Pkg2020,
		"Person",
		test.FullNameProperty,
		test.FamilyNameProperty,
		test.KnownAsProperty,
		addressStatusProperty)

	resource := test.CreateResource(test.Pkg2020, "Person", spec, status)

	defs := make(astmodel.TypeDefinitionSet)
	defs.AddAll(resource, spec, status, addressSpec, locationStatus)

	state := NewState().WithDefinitions(defs)

	cfg := config.NewConfiguration()
	finalState, err := RunTestPipeline(
		state,
		CreateStorageTypes(), // First create the storage types
		InjectSpecInitializationFunctions(cfg, idFactory))
	g.Expect(err).To(Succeed())

	test.AssertPackagesGenerateExpectedCode(t, finalState.Definitions())
}
