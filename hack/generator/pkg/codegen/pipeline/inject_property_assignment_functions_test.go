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
	"github.com/Azure/azure-service-operator/hack/generator/pkg/codegen/storage"
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

	graph := storage.NewConversionGraph()

	// Run CreateStorageTypes first to populate the conversion graph
	createStorageTypes := CreateStorageTypes(graph)
	types, err := createStorageTypes.Run(context.TODO(), types)
	g.Expect(err).To(Succeed())

	// Now run our stage
	injectFunctions := InjectPropertyAssignmentFunctions(graph, idFactory)
	types, err = injectFunctions.Run(context.TODO(), types)
	g.Expect(err).To(Succeed())

	test.AssertPackagesGenerateExpectedCode(t, types, t.Name())
}
