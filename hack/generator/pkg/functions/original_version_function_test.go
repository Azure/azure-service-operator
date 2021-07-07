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

func Test_OriginalVersionFunction_GeneratesExpectedCode(t *testing.T) {
	idFactory := astmodel.NewIdentifierFactory()

	originalVersionFunction := NewOriginalVersionFunction(idFactory)
	demoType := astmodel.NewObjectType().WithFunction(originalVersionFunction)

	demoPkg := astmodel.MakeLocalPackageReference("github.com/Azure/azure-service-operator/hack/generated/_apis", "Person", "vDemo")
	demoName := astmodel.MakeTypeName(demoPkg, "Demo")

	demoDef := astmodel.MakeTypeDefinition(demoName, demoType)

	fileDef := test.CreateFileDefinition(demoDef)
	test.AssertFileGeneratesExpectedCode(t, fileDef, t.Name())
}
