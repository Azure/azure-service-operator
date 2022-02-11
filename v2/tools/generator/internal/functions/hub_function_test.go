/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package functions

import (
	"testing"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/test"
)

func TestGolden_HubFunction_GeneratesExpectedCode(t *testing.T) {
	t.Parallel()
	idFactory := astmodel.NewIdentifierFactory()

	testGroup := "person"
	testPackage := test.MakeLocalPackageReference(testGroup, "v20200101")

	fullNameProperty := astmodel.NewPropertyDefinition("FullName", "fullName", astmodel.StringType).
		WithDescription("As would be used to address mail")

	hubFunction := NewHubFunction(idFactory)

	// Define a test resource
	spec := test.CreateSpec(testPackage, "Person", fullNameProperty)
	status := test.CreateStatus(testPackage, "Person")
	resource := test.CreateResource(testPackage, "Person", spec, status, hubFunction)

	test.AssertSingleTypeDefinitionGeneratesExpectedCode(t, "HubFunction", resource)
}
