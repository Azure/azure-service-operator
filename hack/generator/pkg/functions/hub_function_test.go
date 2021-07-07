/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package functions

import (
	"testing"

	"github.com/Azure/azure-service-operator/hack/generator/pkg/astmodel"
	"github.com/Azure/azure-service-operator/hack/generator/pkg/test"
)

func Test_HubFunction_GeneratesExpectedCode(t *testing.T) {
	idFactory := astmodel.NewIdentifierFactory()

	testGroup := "microsoft.person"
	testPackage := test.MakeLocalPackageReference(testGroup, "v20200101")

	fullNameProperty := astmodel.NewPropertyDefinition("FullName", "fullName", astmodel.StringType).
		WithDescription("As would be used to address mail")

	hubFunction := NewHubFunction(idFactory)

	// Define a test resource
	spec := test.CreateSpec(testPackage, "Person", fullNameProperty)
	status := test.CreateStatus(testPackage, "Person")
	resource := test.CreateResource(testPackage, "Person", spec, status, hubFunction)

	fileDef := test.CreateFileDefinition(resource)
	test.AssertFileGeneratesExpectedCode(t, fileDef, t.Name())
}
