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

func TestGolden_GetSpecFunction_GeneratesExpectedCode(t *testing.T) {
	idFactory := astmodel.NewIdentifierFactory()

	getSpecFunction := NewGetSpecFunction(idFactory)

	// Define a test resource
	spec := test.CreateSpec(test.Pkg2020, "Person", test.FullNameProperty)
	status := test.CreateStatus(test.Pkg2020, "Person")
	resource := test.CreateResource(test.Pkg2020, "Person", spec, status, getSpecFunction)

	test.AssertSingleTypeDefinitionGeneratesExpectedCode(t, "GetSpecFunction", resource)
}
