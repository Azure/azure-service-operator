/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package functions

import (
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/test"
	"testing"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

func TestGolden_GetExtendedResourceFunction_GeneratesExpectedCode(t *testing.T) {
	idFactory := astmodel.NewIdentifierFactory()

	testGroup := "microsoft.person"
	testPackage := test.MakeLocalPackageReference(testGroup, "v20200101")

	fullNameProperty := astmodel.NewPropertyDefinition("FullName", "fullName", astmodel.StringType).
		WithDescription("As would be used to address mail")

	ExtendedResourceFunction := NewGetExtendedResourcesFunction(idFactory)

	// Define a test resource
	spec := test.CreateSpec(testPackage, "Person", fullNameProperty)
	status := test.CreateStatus(testPackage, "Person")
	resource := test.CreateResource(testPackage, "Person", spec, status, ExtendedResourceFunction)

	test.AssertSingleTypeDefinitionGeneratesExpectedCode(t, "GetExtendedResources", resource)
}
