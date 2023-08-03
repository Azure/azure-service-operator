/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package functions

import (
	"testing"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/test"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

func getExtendedResourcesTestData() (string, astmodel.LocalPackageReference) {
	testGroup := "microsoft.person"

	// We don't use test.MakeLocalPackageReference because that forces a 'v' prefix we don't need/want
	extensionPackage := astmodel.MakeLocalPackageReference(test.GoModulePrefix, testGroup, "", "customizations")
	return testGroup, extensionPackage
}

func TestGolden_GetExtendedResourceFunction_oneVersion_GeneratesExpectedCode(t *testing.T) {
	t.Parallel()
	idFactory := astmodel.NewIdentifierFactory()

	testGroup, extensionsPackageRef := getExtendedResourcesTestData()
	testPackage := test.MakeLocalPackageReference(testGroup, "v20200101")

	testResource := astmodel.MakeInternalTypeName(testPackage, "PersonA")
	testSlice := []astmodel.InternalTypeName{
		testResource,
	}

	ExtendedResourceFunction := NewGetExtendedResourcesFunction(idFactory, testSlice)

	resource := test.CreateObjectDefinitionWithFunction(extensionsPackageRef, "PersonAExtension", ExtendedResourceFunction)

	test.AssertSingleTypeDefinitionGeneratesExpectedCode(t, "GetExtendedResources", resource)
}

func TestGolden_GetExtendedResourceFunction_moreThanOneVersion_GeneratesExpectedCode(t *testing.T) {
	t.Parallel()
	idFactory := astmodel.NewIdentifierFactory()

	testGroup, extensionPackage := getExtendedResourcesTestData()
	resourceName := "PersonA"

	testPackageA := test.MakeLocalPackageReference(testGroup, "v20200101")
	testResourceA := astmodel.MakeInternalTypeName(testPackageA, resourceName)

	testPackageB := test.MakeLocalPackageReference(testGroup, "v20200801")
	testResourceB := astmodel.MakeInternalTypeName(testPackageB, resourceName)

	testSlice := []astmodel.InternalTypeName{
		testResourceA,
		testResourceB,
	}

	ExtendedResourceFunction := NewGetExtendedResourcesFunction(idFactory, testSlice)

	resource := test.CreateObjectDefinitionWithFunction(extensionPackage, "PersonAExtension", ExtendedResourceFunction)

	test.AssertSingleTypeDefinitionGeneratesExpectedCode(t, "GetExtendedResources", resource)
}
