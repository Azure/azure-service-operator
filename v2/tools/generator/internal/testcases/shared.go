/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package testcases

import (
	"fmt"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

// idOfGeneratorMethod generates the name of a Generator method, used to create example instances for property testing
func idOfGeneratorMethod(typeName astmodel.TypeName, idFactory astmodel.IdentifierFactory) string {
	name := idFactory.CreateIdentifier(
		fmt.Sprintf("%sGenerator", typeName.Name()),
		astmodel.Exported)
	return name
}
