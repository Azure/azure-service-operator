/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package functions

import (
	"testing"

	"github.com/Azure/azure-service-operator/v2/internal/generator/astmodel"
	"github.com/Azure/azure-service-operator/v2/internal/generator/test"
)

func TestGolden_EmptyStatusFunction_GeneratesExpectedCode(t *testing.T) {
	idFactory := astmodel.NewIdentifierFactory()

	// Define a test resource
	spec := test.CreateSpec(test.Pkg2020, "Person", test.FullNameProperty)
	status := test.CreateStatus(test.Pkg2020, "Person")
	emptyStatusFunction := NewEmptyStatusFunction(status.Name(), idFactory)
	resource := test.CreateResource(test.Pkg2020, "Person", spec, status, emptyStatusFunction)

	test.AssertSingleTypeDefinitionGeneratesExpectedCode(t, "EmptyStatusFunction", resource)
}
