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

	testSpec := test.CreateSpec(testPackage, "PersonA", fullNameProperty)
	testStatus := test.CreateStatus(testPackage, "PersonA")
	testResource := test.CreateResource(testPackage, "PersonA", testSpec, testStatus)
	testSlice := []astmodel.TypeName{testResource.Name()}

	ExtendedResourceFunction := NewGetExtendedResourcesFunction(idFactory, testSlice, []astmodel.PackageReference{testPackage})

	resource := test.CreateObjectDefinitionWithFunction(testPackage, "PersonAExtension", ExtendedResourceFunction)

	test.AssertSingleTypeDefinitionGeneratesExpectedCode(t, "GetExtendedResources", resource)
}
