/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package functions

import (
	"fmt"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astbuilder"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

var GetAPIVersionFunctionName = "Get" + astmodel.APIVersionProperty

// NewGetAPIVersionFunction returns a function that returns a hard coded API Version enum value, e.g. "2024-09-01"
func NewGetAPIVersionFunction(
	apiVersionEnumValue astmodel.EnumValue,
	idFactory astmodel.IdentifierFactory,
) astmodel.Function {
	comment := fmt.Sprintf("returns the ARM API version of the resource. This is always %s", apiVersionEnumValue.Value)
	result := NewObjectFunction(GetAPIVersionFunctionName, idFactory,
		createBodyReturningValue(
			astbuilder.TextLiteral(apiVersionEnumValue.Value),
			astmodel.StringType,
			comment,
			ReceiverTypeStruct))

	return result
}
