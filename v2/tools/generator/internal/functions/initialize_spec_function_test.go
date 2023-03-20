/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package functions

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/test"
)

func TestGolden_InitializeSpecFunction_GeneratesExpectedCode(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	idFactory := astmodel.NewIdentifierFactory()

	// Define a test resource
	spec := test.CreateSpec(test.Pkg2020, "Person", test.FullNameProperty)
	status := test.CreateStatus(test.Pkg2020, "Person")
	resource := test.CreateResource(test.Pkg2020, "Person", spec, status)

	// Create a spec initialization function for that resource
	initializeSpecFunction, err := NewInitializeSpecFunction(resource, "Initialize_From_Status", idFactory)
	g.Expect(err).ToNot(HaveOccurred())

	// Inject the function into the resource
	injector := astmodel.NewFunctionInjector()
	resource, err = injector.Inject(resource, initializeSpecFunction)
	g.Expect(err).ToNot(HaveOccurred())

	test.AssertSingleTypeDefinitionGeneratesExpectedCode(t, "InitializeSpecFunction", resource)
}
