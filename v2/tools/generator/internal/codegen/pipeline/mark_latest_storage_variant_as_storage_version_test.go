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

func TestMarkLatestStorageVariantAsStorageVersion(t *testing.T) {
	g := NewGomegaWithT(t)

	// Test Resource V1
	specV1 := test.CreateSpec(test.Pkg2020, "Person", test.FullNameProperty)
	statusV1 := test.CreateStatus(test.Pkg2020, "Person")
	resourceV1 := test.CreateResource(test.Pkg2020, "Person", specV1, statusV1)

	// Test Resource V2
	specV2 := test.CreateSpec(test.Pkg2021, "Person", test.FullNameProperty)
	statusV2 := test.CreateStatus(test.Pkg2021, "Person")
	resourceV2 := test.CreateResource(test.Pkg2021, "Person", specV2, statusV2)

	types := make(astmodel.Types)
	types.AddAll(resourceV1, specV1, statusV1, resourceV2, specV2, statusV2)

	cfg := config.NewConfiguration()
	initialState, err := RunTestPipeline(
		NewState(types),
		CreateConversionGraph(cfg), // First create the conversion graph showing relationships
		CreateStorageTypes(),       // Then create the storage types
	)
	g.Expect(err).To(Succeed())

	finalState, err := RunTestPipeline(
		initialState,
		MarkLatestStorageVariantAsHubVersion(), // Mark hub/storage variants
	)

	g.Expect(err).To(Succeed())

	// Check that the expected types are flagged as storage types
	// Look for the version of Person flagged with the comment "// +kubebuilder:storageversion"
	test.AssertPackagesGenerateExpectedCode(t, finalState.types, test.DiffWithTypes(initialState.Types()))
}
