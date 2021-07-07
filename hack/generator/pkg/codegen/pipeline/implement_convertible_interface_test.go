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

	graph := storage.NewConversionGraph()

	// Run CreateConversionGraph first to populate the conversion graph
	createConversionGraph := CreateConversionGraph(graph)
	types, err := createConversionGraph.Run(context.TODO(), types)
	g.Expect(err).To(Succeed())

	// Run CreateStorageTypes to create additional resources
	createStorageTypes := CreateStorageTypes(graph)
	types, err = createStorageTypes.Run(context.TODO(), types)
	g.Expect(err).To(Succeed())

	// Run InjectPropertyAssignmentFunctions to create those functions
	injectPropertyFns := InjectPropertyAssignmentFunctions(graph, idFactory)
	types, err = injectPropertyFns.Run(context.TODO(), types)

	// Now run our stage
	injectFunctions := ImplementConvertibleInterface(graph, idFactory)
	types, err = injectFunctions.Run(context.TODO(), types)
	g.Expect(err).To(Succeed())

	test.AssertPackagesGenerateExpectedCode(t, types, t.Name())
}
