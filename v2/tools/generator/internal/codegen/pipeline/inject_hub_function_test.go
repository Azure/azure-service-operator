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

func TestInjectHubFunction_WhenResourceIsStorageVersion_GeneratesExpectedFile(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	idFactory := astmodel.NewIdentifierFactory()
	cfg := config.NewConfiguration()

	// Define a test resource
	spec := test.CreateSpec(test.Pkg2020, "Person", test.FullNameProperty, test.FamilyNameProperty, test.KnownAsProperty)
	status := test.CreateStatus(test.Pkg2020, "Person")
	resource := test.CreateResource(test.Pkg2020, "Person", spec, status)

	resource = resource.WithType(
		resource.Type().(*astmodel.ResourceType).MarkAsStorageVersion())

	defs := make(astmodel.TypeDefinitionSet)
	defs.AddAll(resource, status, spec)

	initialState, err := RunTestPipeline(
		NewState().WithDefinitions(defs),
		CreateStorageTypes(),
		CreateConversionGraph(cfg),
		MarkLatestStorageVariantAsHubVersion(),
		InjectHubFunction(idFactory),
	)
	g.Expect(err).To(Succeed())

	finalState, err := RunTestPipeline(
		initialState,
		InjectHubFunction(idFactory))

	g.Expect(err).To(Succeed())

	test.AssertPackagesGenerateExpectedCode(t, finalState.Definitions())
}

func TestInjectHubFunction_WhenResourceIsNotStorageVersion_GeneratesExpectedFile(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	idFactory := astmodel.NewIdentifierFactory()
	cfg := config.NewConfiguration()

	// Define a test resource
	spec := test.CreateSpec(test.Pkg2020, "Person", test.FullNameProperty, test.FamilyNameProperty, test.KnownAsProperty)
	status := test.CreateStatus(test.Pkg2020, "Person")
	resource := test.CreateResource(test.Pkg2020, "Person", spec, status)

	defs := make(astmodel.TypeDefinitionSet)
	defs.AddAll(resource, status, spec)

	initialState, err := RunTestPipeline(
		NewState().WithDefinitions(defs),
		CreateStorageTypes(),
		CreateConversionGraph(cfg),
		MarkLatestStorageVariantAsHubVersion(),
		InjectHubFunction(idFactory),
	)
	g.Expect(err).To(Succeed())

	finalState, err := RunTestPipeline(
		initialState,
		InjectHubFunction(idFactory))
	g.Expect(err).To(Succeed())
	
	test.AssertPackagesGenerateExpectedCode(t, finalState.Definitions(), test.DiffWithTypes(initialState.Definitions()))
}
