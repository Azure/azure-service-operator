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

	specV1 := test.CreateSpec(pkg2020, "Person", fullNameProperty, familyNameProperty, knownAsProperty)
	statusV1 := test.CreateStatus(pkg2020, "Person")
	resourceV1 := test.CreateResource(pkg2020, "Person", specV1, statusV1)

	// Test Resource V2

	specV2 := test.CreateSpec(
		pkg2021,
		"Person",
		fullNameProperty,
		familyNameProperty,
		knownAsProperty,
		residentialAddress2021,
		postalAddress2021)
	statusV2 := test.CreateStatus(pkg2021, "Person")
	resourceV2 := test.CreateResource(pkg2021, "Person", specV2, statusV2)

	types := make(astmodel.Types)
	types.AddAll(resourceV1, specV1, statusV1, resourceV2, specV2, statusV2, address2021)

	// Run CreateStorageTypes first to populate the conversion graph
	createStorageTypes := CreateStorageTypes()
	state := NewState().WithTypes(types)
	initialState, err := createStorageTypes.Run(context.TODO(), state)
	g.Expect(err).To(Succeed())

	// Now run our stage
	injectFunctions := InjectPropertyAssignmentFunctions(idFactory)
	finalState, err := injectFunctions.Run(context.TODO(), initialState)
	g.Expect(err).To(Succeed())

	test.AssertPackagesGenerateExpectedCode(t, finalState.Types(), t.Name())
}
