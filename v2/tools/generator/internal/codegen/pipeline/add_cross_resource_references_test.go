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

var IDsArrayProperty = astmodel.NewPropertyDefinition("Ids", "ids", astmodel.NewArrayType(astmodel.ARMIDType)).
	WithDescription("This is the important line here. The final type should be []genruntime.ResourceReference")
var IDsMapProperty = astmodel.NewPropertyDefinition("Ids", "ids", astmodel.NewMapType(astmodel.StringType, astmodel.ARMIDType)).
	WithDescription("This is the important line here. The final type should be map[string]genruntime.ResourceReference")

func TestAddCrossResourceReferences_HandlesArray(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// Define a test resource
	spec := test.CreateSpec(test.Pkg2020, "Person", test.FullNameProperty, IDsArrayProperty)
	status := test.CreateStatus(test.Pkg2020, "Person", test.FullNameProperty)
	resource := test.CreateResource(test.Pkg2020, "Person", spec, status)

	defs := make(astmodel.TypeDefinitionSet)
	defs.AddAll(resource, status, spec)

	idFactory := astmodel.NewIdentifierFactory()
	configuration := config.NewConfiguration()

	state, err := RunTestPipeline(
		NewState().WithDefinitions(defs),
		TransformCrossResourceReferences(configuration, idFactory))
	g.Expect(err).ToNot(HaveOccurred())

	test.AssertPackagesGenerateExpectedCode(t, state.Definitions())
}

func TestAddCrossResourceReferences_HandlesMap(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// Define a test resource
	spec := test.CreateSpec(test.Pkg2020, "Person", test.FullNameProperty, IDsMapProperty)
	status := test.CreateStatus(test.Pkg2020, "Person", test.FullNameProperty)
	resource := test.CreateResource(test.Pkg2020, "Person", spec, status)

	defs := make(astmodel.TypeDefinitionSet)
	defs.AddAll(resource, status, spec)

	idFactory := astmodel.NewIdentifierFactory()
	configuration := config.NewConfiguration()

	state, err := RunTestPipeline(
		NewState().WithDefinitions(defs),
		TransformCrossResourceReferences(configuration, idFactory))
	g.Expect(err).ToNot(HaveOccurred())

	test.AssertPackagesGenerateExpectedCode(t, state.Definitions())
}
