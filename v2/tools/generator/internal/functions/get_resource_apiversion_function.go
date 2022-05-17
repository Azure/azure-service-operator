/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package functions

import (
	"fmt"

	"github.com/dave/dst"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astbuilder"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

var GetAPIVersionFunctionName = "Get" + astmodel.APIVersionProperty

// NewGetAPIVersionFunction returns a function that returns a static API Version enum value ('APIVersionValue')
func NewGetAPIVersionFunction(
	apiVersionTypeName astmodel.TypeName,
	apiVersionEnumValue astmodel.EnumValue,
	idFactory astmodel.IdentifierFactory,
) astmodel.Function {
	comment := fmt.Sprintf("returns the ARM API version of the resource. This is always %s", apiVersionEnumValue.Value)
	value := dst.NewIdent(apiVersionTypeName.Name() + apiVersionEnumValue.Identifier)

	result := NewObjectFunction(GetAPIVersionFunctionName, idFactory,
		createBodyReturningValue(
			astbuilder.CallFunc("string", value),
			astmodel.StringType,
			comment,
			ReceiverTypeStruct))

	result.AddPackageReference(astmodel.GenRuntimeReference)
	result.AddReferencedTypes(apiVersionTypeName)

	return result
}
