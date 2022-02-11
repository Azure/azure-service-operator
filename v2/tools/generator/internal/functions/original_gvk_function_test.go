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

func TestGolden_OriginalGVKFunction_ReadingOriginalVersionFromProperty_GeneratesExpectedCode(t *testing.T) {
	t.Parallel()
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

	test.AssertSingleTypeDefinitionGeneratesExpectedCode(t, "OriginalGVKFunction", resource)
}

func TestGolden_OriginalGVKFunction_ReadingOriginalVersionFromFunction_GeneratesExpectedCode(t *testing.T) {
	t.Parallel()
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

	test.AssertSingleTypeDefinitionGeneratesExpectedCode(t, "OriginalGVKFunction", resource)
}
