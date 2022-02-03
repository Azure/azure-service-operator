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

func getExtendedResourcesTestData() (string, astmodel.LocalPackageReference) {
	testGroup := "microsoft.person"
	extensionPackage := test.MakeLocalPackageReference(testGroup, "extensions")
	return testGroup, extensionPackage

}

func TestGolden_GetExtendedResourceFunction_oneVersion_GeneratesExpectedCode(t *testing.T) {
	idFactory := astmodel.NewIdentifierFactory()

	testGroup, extensionsPackageRef := getExtendedResourcesTestData()
	testPackage := test.MakeLocalPackageReference(testGroup, "v20200101")

	testResource := test.CreateTypeName(testPackage, "PersonA")
	testSlice := []astmodel.TypeName{testResource}

	ExtendedResourceFunction := NewGetExtendedResourcesFunction(idFactory, testSlice)

	resource := test.CreateObjectDefinitionWithFunction(extensionsPackageRef, "PersonAExtension", ExtendedResourceFunction)

	test.AssertSingleTypeDefinitionGeneratesExpectedCode(t, "GetExtendedResources", resource)
}

func TestGolden_GetExtendedResourceFunction_moreThanOneVersion_GeneratesExpectedCode(t *testing.T) {
	idFactory := astmodel.NewIdentifierFactory()

	testGroup, extensionPackage := getExtendedResourcesTestData()
	resourceName := "PersonA"

	testPackageA := test.MakeLocalPackageReference(testGroup, "v20200101")
	testResourceA := test.CreateTypeName(testPackageA, resourceName)

	testPackageB := test.MakeLocalPackageReference(testGroup, "v20200801")
	testResourceB := test.CreateTypeName(testPackageB, resourceName)

	testSlice := []astmodel.TypeName{testResourceA, testResourceB}

	ExtendedResourceFunction := NewGetExtendedResourcesFunction(idFactory, testSlice)

	resource := test.CreateObjectDefinitionWithFunction(extensionPackage, "PersonAExtension", ExtendedResourceFunction)

	test.AssertSingleTypeDefinitionGeneratesExpectedCode(t, "GetExtendedResources", resource)
}
