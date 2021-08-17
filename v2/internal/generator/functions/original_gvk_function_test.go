/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package functions

import (
	"testing"

	"github.com/Azure/azure-service-operator/v2/internal/generator/astmodel"
	test "github.com/Azure/azure-service-operator/v2/test/generator"
)

func TestGolden_OriginalGVKFunction_ReadingOriginalVersionFromProperty_GeneratesExpectedCode(t *testing.T) {
	idFactory := astmodel.NewIdentifierFactory()

	testGroup := "microsoft.person"
	testPackage := test.MakeLocalPackageReference(testGroup, "v20200101")

	fullNameProperty := astmodel.NewPropertyDefinition("FullName", "fullName", astmodel.StringType).
		WithDescription("As would be used to address mail")

	originalGVKFunction := NewOriginalGVKFunction(ReadProperty, idFactory)

	// Define a test resource
	spec := test.CreateSpec(testPackage, "Person", fullNameProperty)
	status := test.CreateStatus(testPackage, "Person")
	resource := test.CreateResource(testPackage, "Person", spec, status, originalGVKFunction)

	fileDef := test.CreateFileDefinition(resource)
	test.AssertFileGeneratesExpectedCode(t, fileDef, "OriginalGVKFunction")
}

func TestGolden_OriginalGVKFunction_ReadingOriginalVersionFromFunction_GeneratesExpectedCode(t *testing.T) {
	idFactory := astmodel.NewIdentifierFactory()

	testGroup := "microsoft.person"
	testPackage := test.MakeLocalPackageReference(testGroup, "v20200101")

	fullNameProperty := astmodel.NewPropertyDefinition("FullName", "fullName", astmodel.StringType).
		WithDescription("As would be used to address mail")

	originalGVKFunction := NewOriginalGVKFunction(ReadFunction, idFactory)

	// Define a test resource
	spec := test.CreateSpec(testPackage, "Person", fullNameProperty)
	status := test.CreateStatus(testPackage, "Person")
	resource := test.CreateResource(testPackage, "Person", spec, status, originalGVKFunction)

	fileDef := test.CreateFileDefinition(resource)
	test.AssertFileGeneratesExpectedCode(t, fileDef, "OriginalGVKFunction")
}
