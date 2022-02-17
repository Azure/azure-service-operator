/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/test"
)

func TestGolden_ResoureExtension_OneVersion(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// Test Resource V1
	resourceV1, specV1, statusV1 := getResourceExtensionTestData(test.Pkg2020, "Person")
	defs := make(astmodel.TypeDefinitionSet)
	defs.AddAll(resourceV1, specV1, statusV1)

	initialState, err := RunTestPipeline(
		NewState().WithDefinitions(defs))
	g.Expect(err).To(Succeed())

	finalState, err := RunTestPipeline(initialState, CreateResourceExtensions("testPath", astmodel.NewIdentifierFactory()))
	g.Expect(err).To(Succeed())

	test.AssertPackagesGenerateExpectedCode(t, finalState.Definitions())
}

func TestGolden_ResoureExtension_MoreThanOneVersion(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// Test Resource V1
	resourceV1, specV1, statusV1 := getResourceExtensionTestData(test.Pkg2020, "Person")

	// Test Resource V2
	resourceV2, specV2, statusV2 := getResourceExtensionTestData(test.Pkg2021, "Person")

	defs := make(astmodel.TypeDefinitionSet)
	defs.AddAll(resourceV1, specV1, statusV1, resourceV2, specV2, statusV2)

	initialState, err := RunTestPipeline(
		NewState().WithDefinitions(defs))
	g.Expect(err).To(Succeed())

	finalState, err := RunTestPipeline(initialState, CreateResourceExtensions("testPath", astmodel.NewIdentifierFactory()))
	g.Expect(err).To(Succeed())

	test.AssertPackagesGenerateExpectedCode(t, finalState.Definitions())
}

func getResourceExtensionTestData(pkg astmodel.LocalPackageReference, resourceName string) (astmodel.TypeDefinition, astmodel.TypeDefinition, astmodel.TypeDefinition) {

	spec := test.CreateSpec(pkg, resourceName, test.FullNameProperty, test.FamilyNameProperty, test.KnownAsProperty)
	status := test.CreateStatus(pkg, resourceName)
	resource := test.CreateResource(pkg, resourceName, spec, status)

	return resource, spec, status

}
