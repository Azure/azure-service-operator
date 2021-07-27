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

func TestInjectPropertyAssignmentFunctions(t *testing.T) {
	g := NewGomegaWithT(t)

	idFactory := astmodel.NewIdentifierFactory()
	// Test Resource V1

	specV1 := test.CreateSpec(
		test.Pkg2020,
		"Person",
		test.FullNameProperty,
		test.FamilyNameProperty,
		test.KnownAsProperty)
	statusV1 := test.CreateStatus(test.Pkg2020, "Person")
	resourceV1 := test.CreateResource(test.Pkg2020, "Person", specV1, statusV1)

	// Test Resource V2

	specV2 := test.CreateSpec(
		test.Pkg2021,
		"Person",
		test.FullNameProperty,
		test.FamilyNameProperty,
		test.KnownAsProperty,
		test.ResidentialAddress2021,
		test.PostalAddress2021)
	statusV2 := test.CreateStatus(test.Pkg2021, "Person")
	resourceV2 := test.CreateResource(test.Pkg2021, "Person", specV2, statusV2)

	types := make(astmodel.Types)
	types.AddAll(resourceV1, specV1, statusV1, resourceV2, specV2, statusV2, test.Address2021)

	state := NewState().WithTypes(types)

	// Use CreateConversionGraph to create the conversion graph needed prior to injecting property assignment funcs
	createConversionGraphStage := CreateConversionGraph()
	state, err := createConversionGraphStage.Run(context.TODO(), state)
	g.Expect(err).To(Succeed())

	// Next, create storage types so we have targets for the assignment functions
	createStorageTypesStage := CreateStorageTypes()
	state, err = createStorageTypesStage.Run(context.TODO(), state)
	g.Expect(err).To(Succeed())

	// Now run our stage and inject property assignment functions
	injectFunctions := InjectPropertyAssignmentFunctions(idFactory)
	finalState, err := injectFunctions.Run(context.TODO(), state)
	g.Expect(err).To(Succeed())

	test.AssertPackagesGenerateExpectedCode(t, finalState.Types())
}
