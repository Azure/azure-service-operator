/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/Azure/azure-service-operator/hack/generator/pkg/astmodel"
	"github.com/Azure/azure-service-operator/hack/generator/pkg/test"
)

// TestInjectConvertibleInterface checks that the pipeline stage does what we expect when run in relative isolation,
// with only a few expected (and closely reated) stages in operation
func TestInjectConvertibleInterface(t *testing.T) {
	g := NewGomegaWithT(t)

	idFactory := astmodel.NewIdentifierFactory()
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

	state := NewState().WithTypes(types)

	// Run CreateConversionGraph first to populate the conversion graph
	createConversionGraph := CreateConversionGraph()
	state, err := createConversionGraph.Run(context.TODO(), state)
	g.Expect(err).To(Succeed())

	// Run CreateStorageTypes to create additional resources into which we will inject
	createStorageTypes := CreateStorageTypes()
	state, err = createStorageTypes.Run(context.TODO(), state)
	g.Expect(err).To(Succeed())

	// Run InjectPropertyAssignmentFunctions to create those functions
	injectPropertyFns := InjectPropertyAssignmentFunctions(idFactory)
	state, err = injectPropertyFns.Run(context.TODO(), state)

	// Now run our stage
	injectFunctions := ImplementConvertibleInterface(idFactory)
	state, err = injectFunctions.Run(context.TODO(), state)
	g.Expect(err).To(Succeed())

	// When verifying the golden file, check that the implementations of ConvertTo() and ConvertFrom() are what you
	// expect - if you don't have expectations, check that they do the right thing.
	test.AssertPackagesGenerateExpectedCode(t, state.Types(), t.Name())
}
